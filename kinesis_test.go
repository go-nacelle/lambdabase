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

var testKinesisPayload = `{
	"Records": [
		{
			"eventID": "ev1",
			"kinesis": {
				"PartitionKey": "foo",
				"Data": "WyJ4MSIsICJ5MSIsICJ6MSJdCg=="
			}
		},
		{
			"eventID": "ev2",
			"kinesis": {
				"PartitionKey": "bar",
				"Data": "WyJ4MiIsICJ5MiIsICJ6MiJdCg=="
			}
		},
		{
			"eventID": "ev3",
			"kinesis": {
				"PartitionKey": "baz",
				"Data": "WyJ4MyIsICJ5MyIsICJ6MyJdCg=="
			}
		}
	]
}`

var testKinesisRecords = []events.KinesisEventRecord{
	{
		EventID: "ev1",
		Kinesis: events.KinesisRecord{
			PartitionKey: "foo",
			Data:         []byte{91, 34, 120, 49, 34, 44, 32, 34, 121, 49, 34, 44, 32, 34, 122, 49, 34, 93, 10},
		},
	},
	{
		EventID: "ev2",
		Kinesis: events.KinesisRecord{
			PartitionKey: "bar",
			Data:         []byte{91, 34, 120, 50, 34, 44, 32, 34, 121, 50, 34, 44, 32, 34, 122, 50, 34, 93, 10},
		},
	},
	{
		EventID: "ev3",
		Kinesis: events.KinesisRecord{
			PartitionKey: "baz",
			Data:         []byte{91, 34, 120, 51, 34, 44, 32, 34, 121, 51, 34, 44, 32, 34, 122, 51, 34, 93, 10},
		},
	},
}

func TestKinesisEventInit(t *testing.T) {
	handler := NewMockKinesisEventHandlerInitializer()
	outer := &kinesisEventHandler{
		handler:  handler,
		Logger:   nacelle.NewNilLogger(),
		Services: nacelle.NewServiceContainer(),
	}

	config := nacelle.NewConfig(nacelle.NewTestEnvSourcer(nil))
	err := outer.Init(config)
	assert.Nil(t, err)
	mockassert.CalledOnceMatching(t, handler.InitFunc, func(t assert.TestingT, call interface{}) bool {
		return call.(KinesisEventHandlerInitializerInitFuncCall).Arg0 == config // TODO - ergonomics
	})
}

func TestKinesisEventBadInjection(t *testing.T) {
	handler := &badInjectionKinesisEventHandler{}
	outer := &kinesisEventHandler{
		handler:  handler,
		Logger:   nacelle.NewNilLogger(),
		Services: makeBadContainer(),
	}

	config := nacelle.NewConfig(nacelle.NewTestEnvSourcer(nil))
	err := outer.Init(config)
	assert.Contains(t, err.Error(), "ServiceA")
}

func TestKinesisEventInitError(t *testing.T) {
	handler := NewMockKinesisEventHandlerInitializer()
	handler.InitFunc.SetDefaultReturn(fmt.Errorf("oops"))
	outer := &kinesisEventHandler{
		handler:  handler,
		Logger:   nacelle.NewNilLogger(),
		Services: makeBadContainer(),
	}

	config := nacelle.NewConfig(nacelle.NewTestEnvSourcer(nil))
	err := outer.Init(config)
	assert.EqualError(t, err, "oops")
}

func TestKinesisRecordInit(t *testing.T) {
	handler := NewMockKinesisRecordHandlerInitializer()
	outer := &kinesisRecordHandler{
		handler:  handler,
		Logger:   nacelle.NewNilLogger(),
		Services: nacelle.NewServiceContainer(),
	}

	config := nacelle.NewConfig(nacelle.NewTestEnvSourcer(nil))
	err := outer.Init(config)
	assert.Nil(t, err)
	mockassert.CalledOnceMatching(t, handler.InitFunc, func(t assert.TestingT, call interface{}) bool {
		return call.(KinesisRecordHandlerInitializerInitFuncCall).Arg0 == config // TODO - ergonomics
	})
}

func TestKinesisRecordBadInjection(t *testing.T) {
	handler := &badInjectionKinesisRecordHandler{}
	outer := &kinesisRecordHandler{
		handler:  handler,
		Logger:   nacelle.NewNilLogger(),
		Services: makeBadContainer(),
	}

	config := nacelle.NewConfig(nacelle.NewTestEnvSourcer(nil))
	err := outer.Init(config)
	assert.Contains(t, err.Error(), "ServiceA")
}

func TestKinesisRecordInitError(t *testing.T) {
	handler := NewMockKinesisRecordHandlerInitializer()
	handler.InitFunc.SetDefaultReturn(fmt.Errorf("oops"))
	outer := &kinesisRecordHandler{
		handler:  handler,
		Logger:   nacelle.NewNilLogger(),
		Services: makeBadContainer(),
	}

	config := nacelle.NewConfig(nacelle.NewTestEnvSourcer(nil))
	err := outer.Init(config)
	assert.EqualError(t, err, "oops")
}

func TestKinesisEventInvoke(t *testing.T) {
	handler := NewMockKinesisEventHandlerInitializer()
	outer := &kinesisEventHandler{
		handler: handler,
		Logger:  nacelle.NewNilLogger(),
	}

	response, err := outer.Invoke(context.Background(), []byte(testKinesisPayload))
	assert.Nil(t, err)
	assert.Nil(t, response)
	mockassert.CalledOnceMatching(t, handler.HandleFunc, func(t assert.TestingT, call interface{}) bool {
		return assert.ObjectsAreEqual(call.(KinesisEventHandlerInitializerHandleFuncCall).Arg1, testKinesisRecords) // TODO - ergonomics
	})
}

func TestKinesisEventInvokeError(t *testing.T) {
	handler := NewMockKinesisEventHandlerInitializer()
	outer := &kinesisEventHandler{
		handler: handler,
		Logger:  nacelle.NewNilLogger(),
	}

	handler.HandleFunc.SetDefaultReturn(fmt.Errorf("oops"))
	_, err := outer.Invoke(context.Background(), []byte(testKinesisPayload))
	assert.EqualError(t, err, "failed to process Kinesis event (oops)")
}

func TestKinesisRecordHandle(t *testing.T) {
	handler := NewMockKinesisRecordHandlerInitializer()
	outer := &kinesisRecordHandler{handler: handler}

	err := outer.Handle(context.Background(), testKinesisRecords, nacelle.NewNilLogger())
	assert.Nil(t, err)

	for _, record := range testKinesisRecords {
		mockassert.CalledOnceMatching(t, handler.HandleFunc, func(t assert.TestingT, call interface{}) bool {
			return assert.ObjectsAreEqual(call.(KinesisRecordHandlerInitializerHandleFuncCall).Arg1, record) // TODO - ergonomics
		})
	}
}

func TestKinesisRecordHandleError(t *testing.T) {
	handler := NewMockKinesisRecordHandlerInitializer()
	handler.HandleFunc.PushReturn(nil)
	handler.HandleFunc.PushReturn(fmt.Errorf("oops"))
	outer := &kinesisRecordHandler{handler: handler}

	err := outer.Handle(context.Background(), testKinesisRecords, nacelle.NewNilLogger())
	assert.EqualError(t, err, "failed to process Kinesis record ev2 (oops)")
	mockassert.CalledN(t, handler.HandleFunc, 2)
}

//
// Bad Injection

type badInjectionKinesisEventHandler struct {
	ServiceA *A `service:"A"`
}

func (i *badInjectionKinesisEventHandler) Handle(ctx context.Context, records []events.KinesisEventRecord, logger nacelle.Logger) error {
	return nil
}

type badInjectionKinesisRecordHandler struct {
	ServiceA *A `service:"A"`
}

func (i *badInjectionKinesisRecordHandler) Handle(ctx context.Context, record events.KinesisEventRecord, logger nacelle.Logger) error {
	return nil
}
