package lambdabase

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	mockassert "github.com/efritz/go-mockgen/assert"
	"github.com/go-nacelle/nacelle"
	"github.com/stretchr/testify/assert"
)

var testSQSPayload = `{
	"Records": [
		{"messageId": "m1", "body": "foo"},
		{"messageId": "m2", "body": "bar"},
		{"messageId": "m3", "body": "baz"}
	]
}`

var testSQSMessages = []events.SQSMessage{
	{MessageId: "m1", Body: "foo"},
	{MessageId: "m2", Body: "bar"},
	{MessageId: "m3", Body: "baz"},
}

func TestSQSEventInit(t *testing.T) {
	handler := NewMockSqsEventHandlerInitializer()
	outer := &sqsEventHandler{
		handler:  handler,
		Logger:   nacelle.NewNilLogger(),
		Services: nacelle.NewServiceContainer(),
	}

	config := nacelle.NewConfig(nacelle.NewTestEnvSourcer(nil))
	err := outer.Init(config)
	assert.Nil(t, err)
	mockassert.CalledOnceMatching(t, handler.InitFunc, func(t assert.TestingT, call interface{}) bool {
		return assert.ObjectsAreEqual(call.(SqsEventHandlerInitializerInitFuncCall).Arg0, config) // TODO - ergonomics
	})
}

func TestSQSEventBadInjection(t *testing.T) {
	handler := &badInjectionSQSEventHandler{}
	outer := &sqsEventHandler{
		handler:  handler,
		Logger:   nacelle.NewNilLogger(),
		Services: makeBadContainer(),
	}

	config := nacelle.NewConfig(nacelle.NewTestEnvSourcer(nil))
	err := outer.Init(config)
	assert.Contains(t, err.Error(), "ServiceA")
}

func TestSQSEventInitError(t *testing.T) {
	handler := NewMockSqsEventHandlerInitializer()
	handler.InitFunc.SetDefaultReturn(fmt.Errorf("oops"))
	outer := &sqsEventHandler{
		handler:  handler,
		Logger:   nacelle.NewNilLogger(),
		Services: makeBadContainer(),
	}

	config := nacelle.NewConfig(nacelle.NewTestEnvSourcer(nil))
	err := outer.Init(config)
	assert.EqualError(t, err, "oops")
}

func TestSQSMessageInit(t *testing.T) {
	handler := NewMockSqsMessageHandlerInitializer()
	outer := &sqsMessageHandler{
		handler:  handler,
		Logger:   nacelle.NewNilLogger(),
		Services: nacelle.NewServiceContainer(),
	}

	config := nacelle.NewConfig(nacelle.NewTestEnvSourcer(nil))
	err := outer.Init(config)
	assert.Nil(t, err)
	mockassert.CalledOnceMatching(t, handler.InitFunc, func(t assert.TestingT, call interface{}) bool {
		return assert.ObjectsAreEqual(call.(SqsMessageHandlerInitializerInitFuncCall).Arg0, config) // TODO - ergonomics
	})
}

func TestSQSMessageBadInjection(t *testing.T) {
	handler := &badInjectionSQSMessageHandler{}
	outer := &sqsMessageHandler{
		handler:  handler,
		Logger:   nacelle.NewNilLogger(),
		Services: makeBadContainer(),
	}

	config := nacelle.NewConfig(nacelle.NewTestEnvSourcer(nil))
	err := outer.Init(config)
	assert.Contains(t, err.Error(), "ServiceA")
}

func TestSQSMessageInitError(t *testing.T) {
	handler := NewMockSqsMessageHandlerInitializer()
	handler.InitFunc.SetDefaultReturn(fmt.Errorf("oops"))
	outer := &sqsMessageHandler{
		handler:  handler,
		Logger:   nacelle.NewNilLogger(),
		Services: makeBadContainer(),
	}

	config := nacelle.NewConfig(nacelle.NewTestEnvSourcer(nil))
	err := outer.Init(config)
	assert.EqualError(t, err, "oops")
}

func TestSQSEventInvoke(t *testing.T) {
	handler := NewMockSqsEventHandlerInitializer()
	outer := &sqsEventHandler{
		handler: handler,
		Logger:  nacelle.NewNilLogger(),
	}

	response, err := outer.Invoke(context.Background(), []byte(testSQSPayload))
	assert.Nil(t, err)
	assert.Nil(t, response)
	mockassert.CalledOnceMatching(t, handler.HandleFunc, func(t assert.TestingT, call interface{}) bool {
		return assert.ObjectsAreEqual(call.(SqsEventHandlerInitializerHandleFuncCall).Arg1, testSQSMessages) // TODO - ergonomics
	})
}

func TestSQSEventInvokeError(t *testing.T) {
	handler := NewMockSqsEventHandlerInitializer()
	outer := &sqsEventHandler{
		handler: handler,
		Logger:  nacelle.NewNilLogger(),
	}

	handler.HandleFunc.SetDefaultReturn(fmt.Errorf("oops"))
	_, err := outer.Invoke(context.Background(), []byte(testSQSPayload))
	assert.EqualError(t, err, "failed to process SQS event (oops)")
}

func TestSQSMessageHandle(t *testing.T) {
	handler := NewMockSqsMessageHandlerInitializer()
	outer := &sqsMessageHandler{handler: handler}

	err := outer.Handle(context.Background(), testSQSMessages, nacelle.NewNilLogger())
	assert.Nil(t, err)

	for _, message := range testSQSMessages {
		mockassert.CalledOnceMatching(t, handler.HandleFunc, func(t assert.TestingT, call interface{}) bool {
			return assert.ObjectsAreEqual(call.(SqsMessageHandlerInitializerHandleFuncCall).Arg1, message) // TODO - ergonomics
		})
	}
}

func TestSQSMessageHandleError(t *testing.T) {
	handler := NewMockSqsMessageHandlerInitializer()
	handler.HandleFunc.PushReturn(nil)
	handler.HandleFunc.PushReturn(fmt.Errorf("oops"))
	outer := &sqsMessageHandler{handler: handler}

	err := outer.Handle(context.Background(), testSQSMessages, nacelle.NewNilLogger())
	assert.EqualError(t, err, "failed to process SQS message m2 (oops)")
	mockassert.CalledN(t, handler.HandleFunc, 2)
}

//
// Bad Injection

type badInjectionSQSEventHandler struct {
	ServiceA *A `service:"A"`
}

func (i *badInjectionSQSEventHandler) Handle(ctx context.Context, messages []events.SQSMessage, logger nacelle.Logger) error {
	return nil
}

type badInjectionSQSMessageHandler struct {
	ServiceA *A `service:"A"`
}

func (i *badInjectionSQSMessageHandler) Handle(ctx context.Context, message events.SQSMessage, logger nacelle.Logger) error {
	return nil
}
