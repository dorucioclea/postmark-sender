package internal

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/paysuper/paysuper-recurring-repository/tools"
	"github.com/paysuper/postmark-sender/internal/config"
	"github.com/paysuper/postmark-sender/pkg"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	rabbitmq "gopkg.in/ProtocolONE/rabbitmq.v1/pkg"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	loggerName = "PAYSUPER_POSTMARK_SENDER"

	logParameterRequest  = "request"
	logParameterResponse = "response"

	errorBadResponse = "postmark api return not success http status"
)

type Application struct {
	cfg        *config.Config
	log        *zap.Logger
	broker     rabbitmq.BrokerInterface
	httpClient *http.Client

	fatalFn func(msg string, fields ...zap.Field)
}

type PostmarkResponse struct {
	ErrorCode int32  `json:"ErrorCode"`
	Message   string `json:"Message"`
}

func NewApplication() *Application {
	app := &Application{}

	app.initLogger()
	app.initConfig()
	app.initBroker()

	app.httpClient = tools.NewLoggedHttpClient(zap.S())

	return app
}

func (app *Application) initLogger() {
	logger, err := zap.NewProduction()

	if err != nil {
		log.Fatalf("Application logger initialization failed with error: %s\n", err)
	}

	app.log = logger.Named(loggerName)
	zap.ReplaceGlobals(app.log)

	app.fatalFn = zap.L().Fatal

	zap.L().Info("Logger init...")
}

func (app *Application) initConfig() {
	cfg, err := config.NewConfig()

	if err != nil {
		app.fatalFn("Config init failed", zap.Error(err))
	}

	app.cfg = cfg

	zap.L().Info("Configuration parsed successfully...")
}

func (app *Application) initBroker() {
	var err error

	if app.broker == nil {
		app.broker, err = rabbitmq.NewBroker(app.cfg.BrokerAddress)

		if err != nil {
			app.fatalFn(
				"Creating RabbitMq broker failed",
				zap.Error(err),
				zap.String("url", app.cfg.BrokerAddress),
			)
		}
	}

	app.broker.SetExchangeName(pkg.PostmarkSenderTopicName)
	err = app.broker.RegisterSubscriber(pkg.PostmarkSenderTopicName, app.emailProcess)

	if err != nil {
		app.fatalFn("Registration RabbitMQ broker handler failed", zap.Error(err))
	}

	zap.L().Info("Broker created...")
}

func (app *Application) Run() {
	zap.L().Info("Application started...")

	if err := app.broker.Subscribe(nil); err != nil {
		app.fatalFn("Application subscriber start failed...", zap.Error(err))
	}
}

func (app *Application) Stop() {
	if err := app.log.Sync(); err != nil {
		app.fatalFn("Logger sync failed", zap.Error(err))
	} else {
		zap.L().Info("Logger synced")
	}
}

func (app *Application) emailProcess(payload *pkg.Payload, d amqp.Delivery) error {
	if payload.From == "" {
		payload.From = app.cfg.PostmarkEmailFrom
	}

	if payload.Cc == "" && app.cfg.PostmarkEmailCc != "" {
		payload.Cc = app.cfg.PostmarkEmailCc
	}

	if payload.Bcc == "" && app.cfg.PostmarkEmailBcc != "" {
		payload.Bcc = app.cfg.PostmarkEmailBcc
	}

	if payload.TrackOpens != app.cfg.PostmarkEmailTrackOpens {
		payload.TrackOpens = app.cfg.PostmarkEmailTrackOpens
	}

	if payload.TrackLinks != app.cfg.PostmarkEmailTrackTrackLinks {
		payload.TrackLinks = app.cfg.PostmarkEmailTrackTrackLinks
	}

	return app.sendEmail(payload)
}

func (app *Application) sendEmail(payload *pkg.Payload) error {
	b, err := json.Marshal(payload)

	if err != nil {
		zap.L().Error(
			"Email payload marshaling failed",
			zap.Error(err),
			zap.Any(logParameterRequest, payload),
		)

		return err
	}

	req, err := http.NewRequest(http.MethodPost, app.cfg.PostmarkApiUrl, bytes.NewBuffer(b))

	if err != nil {
		zap.L().Error(
			"Creating http request failed",
			zap.Error(err),
			zap.Any(logParameterRequest, payload),
		)

		return err
	}

	req.Header.Add(pkg.HeaderContentType, pkg.MIMEApplicationJSON)
	req.Header.Add(pkg.HeaderAccept, pkg.MIMEApplicationJSON)
	req.Header.Add(pkg.HeaderXPostmarkServerToken, app.cfg.PostmarkApiToken)

	rsp, err := app.httpClient.Do(req)

	if err != nil {
		zap.L().Error(
			"Send email failed",
			zap.Error(err),
			zap.Any(logParameterRequest, payload),
		)

		return err
	}

	b, err = ioutil.ReadAll(rsp.Body)
	_ = rsp.Body.Close()

	if err != nil {
		zap.L().Error(
			"Reading response body failed",
			zap.Error(err),
			zap.Any(logParameterRequest, payload),
		)

		return err
	}

	msg := &PostmarkResponse{}
	err = json.Unmarshal(b, msg)

	if err != nil {
		zap.L().Error(
			"Incorrect json response",
			zap.Any(logParameterRequest, payload),
			zap.ByteString(logParameterResponse, b),
		)

		return err
	}

	if rsp.StatusCode != http.StatusOK || msg.IsSuccess() == false {
		zap.L().Error(
			errorBadResponse,
			zap.Any(logParameterRequest, payload),
			zap.Any(logParameterResponse, msg),
		)

		return errors.New(errorBadResponse)
	}

	return nil
}

func (m *PostmarkResponse) IsSuccess() bool {
	return m.ErrorCode == pkg.PostMarkErrorCodeSuccess
}