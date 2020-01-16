package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/golang/protobuf/jsonpb"
	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/paysuper/paysuper-proto/go/postmarkpb"
	"github.com/paysuper/postmark-sender/internal/config"
	"github.com/paysuper/postmark-sender/pkg"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	rabbitmq "gopkg.in/ProtocolONE/rabbitmq.v1/pkg"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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

type postmarkHttpTransport struct {
	Transport http.RoundTripper
}

type postmarkContextKey struct {
	name string
}

func NewApplication() *Application {
	app := &Application{}

	app.initLogger()
	app.initConfig()

	if err := app.initBroker(); err != nil {
		app.fatalFn(
			"RabbitMq broker failed",
			zap.Error(err),
			zap.String("url", app.cfg.BrokerAddress),
		)
	}

	app.httpClient = NewHttpClient()

	return app
}

func NewHttpClient() *http.Client {
	return &http.Client{Transport: &postmarkHttpTransport{}}
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

func (app *Application) initBroker() error {
	var err error

	if app.broker == nil {
		app.broker, err = rabbitmq.NewBroker(app.cfg.BrokerAddress)

		if err != nil {
			return err
		}
	}

	app.broker.SetExchangeName(pkg.PostmarkSenderTopicName)
	err = app.broker.RegisterSubscriber(pkg.PostmarkSenderTopicName, app.emailProcess)

	if err != nil {
		return err
	}

	zap.L().Info("Broker created...")

	return nil
}

func (app *Application) Run() error {
	zap.L().Info("Application started...")

	if err := app.broker.Subscribe(nil); err != nil {
		zap.L().Error("Application subscriber start failed...", zap.Error(err))
		return err
	}

	return nil
}

func (app *Application) Stop() {
	if err := app.log.Sync(); err != nil {
		app.fatalFn("Logger sync failed", zap.Error(err))
	} else {
		zap.L().Info("Logger synced")
	}
}

func (app *Application) emailProcess(payload *postmarkpb.Payload, d amqp.Delivery) error {
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

func (app *Application) sendEmail(payload *postmarkpb.Payload) error {
	if len(payload.TemplateModel) > 0 {
		if payload.TemplateObjectModel == nil {
			payload.TemplateObjectModel = &structpb.Struct{
				Fields: map[string]*structpb.Value{},
			}
		}
		for key, item := range payload.TemplateModel {
			payload.TemplateObjectModel.Fields[key] = &structpb.Value{
				Kind: &structpb.Value_StringValue{StringValue: item},
			}
		}
	}

	march := &jsonpb.Marshaler{}
	var buf bytes.Buffer
	err := march.Marshal(&buf, payload)

	if err != nil {
		zap.L().Error(
			"Email payload marshaling failed",
			zap.Error(err),
			zap.Any(logParameterRequest, payload),
		)

		return err
	}

	req, err := http.NewRequest(http.MethodPost, app.cfg.PostmarkApiUrl, &buf)

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

	b, err := ioutil.ReadAll(rsp.Body)
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

func (t *postmarkHttpTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx := context.WithValue(req.Context(), &postmarkContextKey{name: "PostmarkRequestStart"}, time.Now())
	req = req.WithContext(ctx)

	var reqBody []byte

	if req.Body != nil {
		reqBody, _ = ioutil.ReadAll(req.Body)
	}

	req.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
	rsp, err := http.DefaultTransport.RoundTrip(req)

	if err != nil {
		return rsp, err
	}

	var rspBody []byte

	if rsp.Body != nil {
		rspBody, err = ioutil.ReadAll(rsp.Body)

		if err != nil {
			return rsp, err
		}
	}

	rsp.Body = ioutil.NopCloser(bytes.NewBuffer(rspBody))

	req.Header.Set(pkg.HeaderXPostmarkServerToken, "*****")

	zap.L().Info(
		req.URL.Path,
		zap.Any("request_headers", req.Header),
		zap.ByteString("request_body", reqBody),
		zap.Int("response_status", rsp.StatusCode),
		zap.Any("response_headers", rsp.Header),
		zap.ByteString("response_body", rspBody),
	)

	return rsp, err
}
