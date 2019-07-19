package internal

import (
	"github.com/paysuper/paysuper-billing-server/pkg/proto/grpc"
	"github.com/paysuper/postmark-sender/internal/config"
	"github.com/paysuper/postmark-sender/internal/mock"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
	rabbitmq "gopkg.in/ProtocolONE/rabbitmq.v1/pkg"
	"net/http"
	"os"
	"testing"
)

type ApplicationTestSuite struct {
	suite.Suite
	app         *Application
	zapRecorder *observer.ObservedLogs
}

func Test_Application(t *testing.T) {
	suite.Run(t, new(ApplicationTestSuite))
}

func (suite *ApplicationTestSuite) SetupTest() {
	suite.app = NewApplication()
	assert.NotNil(suite.T(), suite.app)
	assert.Nil(suite.T(), suite.app.cfg)
	assert.Nil(suite.T(), suite.app.log)
	assert.Nil(suite.T(), suite.app.broker)
	assert.Nil(suite.T(), suite.app.httpClient)

	cfg, err := config.NewConfig()

	if err != nil {
		suite.FailNow("Config load failed", "%v", err)
	}

	suite.app.cfg = cfg
	suite.app.broker = mock.NewBrokerMockOk()
	suite.app.httpClient = mock.NewClientStatusOk()

	var core zapcore.Core

	lvl := zap.NewAtomicLevel()
	core, suite.zapRecorder = observer.New(lvl)
	suite.app.log = zap.New(core)
	zap.ReplaceGlobals(suite.app.log)
}

func (suite *ApplicationTestSuite) TearDownTest() {}

func (suite *ApplicationTestSuite) TestApplication_InitBroker_Ok() {
	suite.app.initBroker()
	messages := suite.zapRecorder.All()
	assert.Equal(suite.T(), "Broker created...", messages[0].Message)
}

func (suite *ApplicationTestSuite) TestApplication_InitBroker_NewBroker_Error() {
	suite.app.fatalFn = zap.L().Info
	suite.app.cfg.BrokerAddress = ""
	suite.app.broker = nil

	suite.app.initBroker()
	messages := suite.zapRecorder.All()
	assert.Equal(suite.T(), "Creating RabbitMq broker failed", messages[0].Message)
}

func (suite *ApplicationTestSuite) TestApplication_InitBroker_RegisterSubscriber_Error() {
	suite.app.fatalFn = zap.L().Info
	suite.app.cfg.BrokerAddress = ""
	suite.app.broker = mock.NewBrokerMockError()

	suite.app.initBroker()
	messages := suite.zapRecorder.All()
	assert.Equal(suite.T(), "Registration RabbitMQ broker handler failed", messages[0].Message)
}

func (suite *ApplicationTestSuite) TestApplication_InitConfig_Ok() {
	suite.app.initConfig()
	messages := suite.zapRecorder.All()
	assert.Equal(suite.T(), "Configuration parsed successfully...", messages[0].Message)
}

func (suite *ApplicationTestSuite) TestApplication_InitConfig_Error() {
	v := os.Getenv("POSTMARK_EMAIL_SENDER")
	err := os.Unsetenv("POSTMARK_EMAIL_SENDER")
	assert.NoError(suite.T(), err)

	suite.app.fatalFn = zap.L().Info

	suite.app.initConfig()
	messages := suite.zapRecorder.All()
	assert.Equal(suite.T(), "Config init failed", messages[0].Message)

	err = os.Setenv("POSTMARK_EMAIL_SENDER", v)
	assert.NoError(suite.T(), err)
}

func (suite *ApplicationTestSuite) TestApplication_InitLogger_Ok() {
	app := NewApplication()
	assert.Nil(suite.T(), app.log)

	app.initLogger()
	assert.NotNil(suite.T(), app.log)
	assert.IsType(suite.T(), &zap.Logger{}, app.log)
}

func (suite *ApplicationTestSuite) TestApplication_Init_Ok() {
	app := NewApplication()
	assert.Nil(suite.T(), app.log)
	assert.Nil(suite.T(), app.cfg)
	assert.Nil(suite.T(), app.broker)
	assert.Nil(suite.T(), app.httpClient)
	assert.Nil(suite.T(), app.fatalFn)

	app.Init()
	assert.NotNil(suite.T(), app.log)
	assert.IsType(suite.T(), &zap.Logger{}, app.log)
	assert.NotNil(suite.T(), app.cfg)
	assert.IsType(suite.T(), &config.Config{}, app.cfg)
	assert.NotNil(suite.T(), app.broker)
	assert.Implements(suite.T(), (*rabbitmq.BrokerInterface)(nil), app.broker)
	assert.NotNil(suite.T(), app.httpClient)
	assert.IsType(suite.T(), &http.Client{}, app.httpClient)
	assert.NotNil(suite.T(), app.fatalFn)
}

func (suite *ApplicationTestSuite) TestApplication_Run_Ok() {
	suite.app.Run()
	messages := suite.zapRecorder.All()
	assert.Equal(suite.T(), "Application started...", messages[0].Message)
}

func (suite *ApplicationTestSuite) TestApplication_Run_Subscribe_Error() {
	suite.app.broker = mock.NewBrokerMockError()
	suite.app.fatalFn = zap.L().Info

	suite.app.Run()
	messages := suite.zapRecorder.All()
	assert.Equal(suite.T(), "Application subscriber start failed...", messages[1].Message)
}

func (suite *ApplicationTestSuite) TestApplication_EmailConfirmProcess_Ok() {
	profile := &grpc.UserProfile{
		Email: &grpc.UserProfileEmail{
			Email:           "dmitriy.sinichkin@protocol.one",
			ConfirmationUrl: "http://localhost?token=123456",
		},
		Personal: &grpc.UserProfilePersonal{
			FirstName: "Unit test",
			LastName:  "Unit Test",
			Position:  "test",
		},
	}
	err := suite.app.emailConfirmProcess(profile, amqp.Delivery{})
	assert.NoError(suite.T(), err)
}

func (suite *ApplicationTestSuite) TestApplication_EmailConfirmProcess_CreatingHttpRequestError() {
	profile := &grpc.UserProfile{
		Email: &grpc.UserProfileEmail{
			Email:           "dmitriy.sinichkin@protocol.one",
			ConfirmationUrl: "http://localhost?token=123456",
		},
		Personal: &grpc.UserProfilePersonal{
			FirstName: "Unit test",
			LastName:  "Unit Test",
			Position:  "test",
		},
	}

	suite.app.cfg.PostmarkApiUrl = "\n"

	err := suite.app.emailConfirmProcess(profile, amqp.Delivery{})
	assert.Error(suite.T(), err)
	messages := suite.zapRecorder.All()
	assert.Equal(suite.T(), "Creating http request failed", messages[0].Message)
}

func (suite *ApplicationTestSuite) TestApplication_EmailConfirmProcess_SendEmailFailedError() {
	profile := &grpc.UserProfile{
		Email: &grpc.UserProfileEmail{
			Email:           "dmitriy.sinichkin@protocol.one",
			ConfirmationUrl: "http://localhost?token=123456",
		},
		Personal: &grpc.UserProfilePersonal{
			FirstName: "Unit test",
			LastName:  "Unit Test",
			Position:  "test",
		},
	}

	suite.app.httpClient = mock.NewTransportHttpError()

	err := suite.app.emailConfirmProcess(profile, amqp.Delivery{})
	assert.Error(suite.T(), err)
	messages := suite.zapRecorder.All()
	assert.Equal(suite.T(), "Send email failed", messages[0].Message)
}

func (suite *ApplicationTestSuite) TestApplication_EmailConfirmProcess_ReadingResponseBodyError() {
	profile := &grpc.UserProfile{
		Email: &grpc.UserProfileEmail{
			Email:           "dmitriy.sinichkin@protocol.one",
			ConfirmationUrl: "http://localhost?token=123456",
		},
		Personal: &grpc.UserProfilePersonal{
			FirstName: "Unit test",
			LastName:  "Unit Test",
			Position:  "test",
		},
	}

	suite.app.httpClient = mock.NewClientStatusErrorIoReader()

	err := suite.app.emailConfirmProcess(profile, amqp.Delivery{})
	assert.Error(suite.T(), err)
	messages := suite.zapRecorder.All()
	assert.Equal(suite.T(), "Reading response body failed", messages[0].Message)
}

func (suite *ApplicationTestSuite) TestApplication_EmailConfirmProcess_IncorrectJsonResponseError() {
	profile := &grpc.UserProfile{
		Email: &grpc.UserProfileEmail{
			Email:           "dmitriy.sinichkin@protocol.one",
			ConfirmationUrl: "http://localhost?token=123456",
		},
		Personal: &grpc.UserProfilePersonal{
			FirstName: "Unit test",
			LastName:  "Unit Test",
			Position:  "test",
		},
	}

	suite.app.httpClient = mock.NewClientStatusIncorrectResponse()

	err := suite.app.emailConfirmProcess(profile, amqp.Delivery{})
	assert.Error(suite.T(), err)
	messages := suite.zapRecorder.All()
	assert.Equal(suite.T(), "Incorrect json response", messages[0].Message)
}

func (suite *ApplicationTestSuite) TestApplication_EmailConfirmProcess_BadResponseStatusError() {
	profile := &grpc.UserProfile{
		Email: &grpc.UserProfileEmail{
			Email:           "dmitriy.sinichkin@protocol.one",
			ConfirmationUrl: "http://localhost?token=123456",
		},
		Personal: &grpc.UserProfilePersonal{
			FirstName: "Unit test",
			LastName:  "Unit Test",
			Position:  "test",
		},
	}

	suite.app.httpClient = mock.NewClientStatusBadStatus()

	err := suite.app.emailConfirmProcess(profile, amqp.Delivery{})
	assert.Error(suite.T(), err)
	messages := suite.zapRecorder.All()
	assert.Equal(suite.T(), errorBadResponse, messages[0].Message)
}
