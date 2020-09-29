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

var testDynamoDBPayload = `{
	"Records": [
		{
			"eventID": "ev1",
			"eventName": "INSERT",
			"dynamodb": {
				"NewImage": {
					"PK": {"S": "foo"},
					"SK": {"S": "bonk"}
				}
			}
		},
		{
			"eventID": "ev2",
			"eventName": "INSERT",
			"dynamodb": {
				"NewImage": {
					"PK": {"S": "bar"},
					"SK": {"S": "quux"}
				}
			}
		},
		{
			"eventID": "ev3",
			"eventName": "INSERT",
			"dynamodb": {
				"NewImage": {
					"PK": {"S": "baz"},
					"SK": {"S": "honk"}
				}
			}
		}
	]
}`

var testDynamoDBRecords = []events.DynamoDBEventRecord{
	{
		EventID:   "ev1",
		EventName: "INSERT",
		Change: events.DynamoDBStreamRecord{
			NewImage: map[string]events.DynamoDBAttributeValue{
				"PK": events.NewStringAttribute("foo"),
				"SK": events.NewStringAttribute("bonk"),
			},
		},
	},
	{
		EventID:   "ev2",
		EventName: "INSERT",
		Change: events.DynamoDBStreamRecord{
			NewImage: map[string]events.DynamoDBAttributeValue{
				"PK": events.NewStringAttribute("bar"),
				"SK": events.NewStringAttribute("quux"),
			},
		},
	},
	{
		EventID:   "ev3",
		EventName: "INSERT",
		Change: events.DynamoDBStreamRecord{
			NewImage: map[string]events.DynamoDBAttributeValue{
				"PK": events.NewStringAttribute("baz"),
				"SK": events.NewStringAttribute("honk"),
			},
		},
	},
}

func TestDynamoDBEventInit(t *testing.T) {
	handler := NewMockDynamoDBEventHandlerInitializer()
	outer := &dynamoDBEventHandler{
		handler:  handler,
		Logger:   nacelle.NewNilLogger(),
		Services: nacelle.NewServiceContainer(),
	}

	config := nacelle.NewConfig(nacelle.NewTestEnvSourcer(nil))
	err := outer.Init(config)
	assert.Nil(t, err)
	mockassert.CalledOnceMatching(t, handler.InitFunc, func(t assert.TestingT, call interface{}) bool {
		return assert.ObjectsAreEqual(call.(DynamoDBEventHandlerInitializerInitFuncCall).Arg0, config) // TODO - ergonomics
	})
}

func TestDynamoDBEventBadInjection(t *testing.T) {
	handler := &badInjectionDynamoDBEventHandler{}
	outer := &dynamoDBEventHandler{
		handler:  handler,
		Logger:   nacelle.NewNilLogger(),
		Services: makeBadContainer(),
	}

	config := nacelle.NewConfig(nacelle.NewTestEnvSourcer(nil))
	err := outer.Init(config)
	assert.Contains(t, err.Error(), "ServiceA")
}

func TestDynamoDBEventInitError(t *testing.T) {
	handler := NewMockDynamoDBEventHandlerInitializer()
	handler.InitFunc.SetDefaultReturn(fmt.Errorf("oops"))
	outer := &dynamoDBEventHandler{
		handler:  handler,
		Logger:   nacelle.NewNilLogger(),
		Services: nacelle.NewServiceContainer(),
	}

	config := nacelle.NewConfig(nacelle.NewTestEnvSourcer(nil))
	err := outer.Init(config)
	assert.EqualError(t, err, "oops")
}

func TestDynamoDBRecordInit(t *testing.T) {
	handler := NewMockDynamoDBRecordHandlerInitializer()
	outer := &dynamoDBRecordHandler{
		handler:  handler,
		Services: nacelle.NewServiceContainer(),
	}

	config := nacelle.NewConfig(nacelle.NewTestEnvSourcer(nil))
	err := outer.Init(config)
	assert.Nil(t, err)
	mockassert.CalledOnceMatching(t, handler.InitFunc, func(t assert.TestingT, call interface{}) bool {
		return call.(DynamoDBRecordHandlerInitializerInitFuncCall).Arg0 == config // TODO - ergonomics
	})
}

func TestDynamoDBRecordBadInjection(t *testing.T) {
	handler := &badInjectionDynamoDBRecordHandler{}
	outer := &dynamoDBRecordHandler{
		handler:  handler,
		Services: makeBadContainer(),
	}

	config := nacelle.NewConfig(nacelle.NewTestEnvSourcer(nil))
	err := outer.Init(config)
	assert.Contains(t, err.Error(), "ServiceA")
}

func TestDynamoDBRecordInitError(t *testing.T) {
	handler := NewMockDynamoDBRecordHandlerInitializer()
	handler.InitFunc.SetDefaultReturn(fmt.Errorf("oops"))
	outer := &dynamoDBRecordHandler{
		handler:  handler,
		Services: nacelle.NewServiceContainer(),
	}

	config := nacelle.NewConfig(nacelle.NewTestEnvSourcer(nil))
	err := outer.Init(config)
	assert.EqualError(t, err, "oops")
}

func TestDynamoDBEventInvoke(t *testing.T) {
	handler := NewMockDynamoDBEventHandlerInitializer()
	outer := &dynamoDBEventHandler{
		handler: handler,
		Logger:  nacelle.NewNilLogger(),
	}

	response, err := outer.Invoke(context.Background(), []byte(testDynamoDBPayload))
	assert.Nil(t, err)
	assert.Nil(t, response)
	mockassert.CalledOnceMatching(t, handler.HandleFunc, func(t assert.TestingT, call interface{}) bool {
		return assert.ObjectsAreEqual(call.(DynamoDBEventHandlerInitializerHandleFuncCall).Arg1, testDynamoDBRecords) // TODO - ergonomics
	})
}

func TestDynamoDBEventInvokeError(t *testing.T) {
	handler := NewMockDynamoDBEventHandlerInitializer()
	outer := &dynamoDBEventHandler{
		handler: handler,
		Logger:  nacelle.NewNilLogger(),
	}

	handler.HandleFunc.SetDefaultReturn(fmt.Errorf("oops"))
	_, err := outer.Invoke(context.Background(), []byte(testDynamoDBPayload))
	assert.EqualError(t, err, "failed to process DynamoDB event (oops)")
}

func TestDynamoDBRecordHandle(t *testing.T) {
	handler := NewMockDynamoDBRecordHandlerInitializer()
	outer := &dynamoDBRecordHandler{handler: handler}

	err := outer.Handle(context.Background(), testDynamoDBRecords, nacelle.NewNilLogger())
	assert.Nil(t, err)

	for _, record := range testDynamoDBRecords {
		mockassert.CalledOnceMatching(t, handler.HandleFunc, func(t assert.TestingT, call interface{}) bool {
			return assert.ObjectsAreEqual(call.(DynamoDBRecordHandlerInitializerHandleFuncCall).Arg1, record) // TODO - ergonomics
		})
	}
}

func TestDynamoDBRecordHandleError(t *testing.T) {
	handler := NewMockDynamoDBRecordHandlerInitializer()
	handler.HandleFunc.PushReturn(nil)
	handler.HandleFunc.PushReturn(fmt.Errorf("oops"))
	outer := &dynamoDBRecordHandler{handler: handler}

	err := outer.Handle(context.Background(), testDynamoDBRecords, nacelle.NewNilLogger())
	assert.EqualError(t, err, "failed to process DynamoDB record ev2 (oops)")
	mockassert.CalledN(t, handler.HandleFunc, 2)
}

//
// Bad Injection

type badInjectionDynamoDBEventHandler struct {
	ServiceA *A `service:"A"`
}

func (i *badInjectionDynamoDBEventHandler) Handle(ctx context.Context, records []events.DynamoDBEventRecord, logger nacelle.Logger) error {
	return nil
}

type badInjectionDynamoDBRecordHandler struct {
	ServiceA *A `service:"A"`
}

func (i *badInjectionDynamoDBRecordHandler) Handle(ctx context.Context, record events.DynamoDBEventRecord, logger nacelle.Logger) error {
	return nil
}
