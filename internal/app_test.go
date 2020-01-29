package internal

import (
	"errors"
	"github.com/paysuper/paysuper-proto/go/postmarkpb"
	"github.com/paysuper/paysuper-tools/http/mocks"
	"github.com/paysuper/postmark-sender/pkg"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap/zaptest/observer"
	rabbitMocks "gopkg.in/ProtocolONE/rabbitmq.v1/pkg/mocks"
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
	suite.app = &Application{}
	suite.app.initConfig()
	suite.app.initLogger()

	suite.app.broker = &rabbitMocks.BrokerInterface{}
	suite.app.httpClient = mocks.NewClientStatusOk()
}

func (suite *ApplicationTestSuite) TearDownTest() {}

func (suite *ApplicationTestSuite) TestApplication_InitBroker_Ok() {
	broker := &rabbitMocks.BrokerInterface{}
	broker.On("SetExchangeName", pkg.PostmarkSenderTopicName).Return(nil)
	broker.On("RegisterSubscriber", pkg.PostmarkSenderTopicName, mock.Anything).Return(nil)
	suite.app.broker = broker

	assert.NoError(suite.T(), suite.app.initBroker())
}

func (suite *ApplicationTestSuite) TestApplication_InitBroker_NewBroker_Error() {
	suite.app.cfg.BrokerAddress = ""
	suite.app.broker = nil

	assert.Error(suite.T(), suite.app.initBroker())
}

func (suite *ApplicationTestSuite) TestApplication_InitBroker_RegisterSubscriber_Error() {
	broker := &rabbitMocks.BrokerInterface{}
	broker.On("SetExchangeName", pkg.PostmarkSenderTopicName).Return(nil)
	broker.On("RegisterSubscriber", pkg.PostmarkSenderTopicName, mock.Anything).Return(errors.New("error"))
	suite.app.broker = broker

	assert.Error(suite.T(), suite.app.initBroker())
}

func (suite *ApplicationTestSuite) TestApplication_Run_Subscribe_Error() {
	broker := &rabbitMocks.BrokerInterface{}
	var x chan bool
	broker.On("Subscribe", x).Return(errors.New("error"))
	suite.app.broker = broker

	assert.Error(suite.T(), suite.app.Run())
}

func (suite *ApplicationTestSuite) TestApplication_EmailConfirmProcess_Ok() {
	payload := &postmarkpb.Payload{
		TemplateAlias: "template1",
		TemplateModel: map[string]string{"param1": "value1"},
	}

	err := suite.app.emailProcess(payload, amqp.Delivery{})
	assert.NoError(suite.T(), err)
}

func (suite *ApplicationTestSuite) TestApplication_EmailConfirmProcess_SendEmailFailedError() {
	payload := &postmarkpb.Payload{
		TemplateAlias: "template1",
		TemplateModel: map[string]string{"param1": "value1"},
	}

	suite.app.httpClient = mocks.NewTransportHttpError()

	err := suite.app.emailProcess(payload, amqp.Delivery{})
	assert.Errorf(suite.T(), err, "Send email failed")
}

func (suite *ApplicationTestSuite) TestApplication_EmailConfirmProcess_ReadingResponseBodyError() {
	payload := &postmarkpb.Payload{
		TemplateAlias: "template1",
		TemplateModel: map[string]string{"param1": "value1"},
	}

	suite.app.httpClient = mocks.NewClientStatusErrorIoReader()

	err := suite.app.emailProcess(payload, amqp.Delivery{})
	assert.Errorf(suite.T(), err, "Reading response body failed")
}

func (suite *ApplicationTestSuite) TestApplication_EmailConfirmProcess_IncorrectJsonResponseError() {
	payload := &postmarkpb.Payload{
		TemplateAlias: "template1",
		TemplateModel: map[string]string{"param1": "value1"},
	}
	suite.app.httpClient = mocks.NewClientStatusIncorrectResponse()

	err := suite.app.emailProcess(payload, amqp.Delivery{})
	assert.Errorf(suite.T(), err, "Incorrect json response")
}

func (suite *ApplicationTestSuite) TestApplication_EmailConfirmProcess_BadResponseStatusError() {
	payload := &postmarkpb.Payload{
		TemplateAlias: "template1",
		TemplateModel: map[string]string{"param1": "value1"},
	}

	suite.app.httpClient = mocks.NewClientStatusBadStatus()

	err := suite.app.emailProcess(payload, amqp.Delivery{})
	assert.Errorf(suite.T(), err, errorBadResponse)
}
