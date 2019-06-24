// Code generated by github.com/efritz/go-mockgen 0.1.0; DO NOT EDIT.
// This file was generated by robots at
// 2019-06-21T14:12:31-05:00
// using the command
// $ go-mockgen -f github.com/go-nacelle/lambdabase -i kinesisEventHandlerInitializer -i kinesisRecordHandlerInitializer -o kinesis_mock_test.go

package lambdabase

import (
	"context"
	events "github.com/aws/aws-lambda-go/events"
	config "github.com/go-nacelle/config"
	log "github.com/go-nacelle/log"
	"sync"
)

// MockKinesisEventHandlerInitializer is a mock implementation of the
// kinesisEventHandlerInitializer interface (from the package
// github.com/go-nacelle/lambdabase) used for unit testing.
type MockKinesisEventHandlerInitializer struct {
	// HandleFunc is an instance of a mock function object controlling the
	// behavior of the method Handle.
	HandleFunc *KinesisEventHandlerInitializerHandleFunc
	// InitFunc is an instance of a mock function object controlling the
	// behavior of the method Init.
	InitFunc *KinesisEventHandlerInitializerInitFunc
}

// NewMockKinesisEventHandlerInitializer creates a new mock of the
// kinesisEventHandlerInitializer interface. All methods return zero values
// for all results, unless overwritten.
func NewMockKinesisEventHandlerInitializer() *MockKinesisEventHandlerInitializer {
	return &MockKinesisEventHandlerInitializer{
		HandleFunc: &KinesisEventHandlerInitializerHandleFunc{
			defaultHook: func(context.Context, []events.KinesisEventRecord, log.Logger) error {
				return nil
			},
		},
		InitFunc: &KinesisEventHandlerInitializerInitFunc{
			defaultHook: func(config.Config) error {
				return nil
			},
		},
	}
}

// surrogateMockKinesisEventHandlerInitializer is a copy of the
// kinesisEventHandlerInitializer interface (from the package
// github.com/go-nacelle/lambdabase). It is redefined here as it is
// unexported in the source package.
type surrogateMockKinesisEventHandlerInitializer interface {
	Handle(context.Context, []events.KinesisEventRecord, log.Logger) error
	Init(config.Config) error
}

// NewMockKinesisEventHandlerInitializerFrom creates a new mock of the
// MockKinesisEventHandlerInitializer interface. All methods delegate to the
// given implementation, unless overwritten.
func NewMockKinesisEventHandlerInitializerFrom(i surrogateMockKinesisEventHandlerInitializer) *MockKinesisEventHandlerInitializer {
	return &MockKinesisEventHandlerInitializer{
		HandleFunc: &KinesisEventHandlerInitializerHandleFunc{
			defaultHook: i.Handle,
		},
		InitFunc: &KinesisEventHandlerInitializerInitFunc{
			defaultHook: i.Init,
		},
	}
}

// KinesisEventHandlerInitializerHandleFunc describes the behavior when the
// Handle method of the parent MockKinesisEventHandlerInitializer instance
// is invoked.
type KinesisEventHandlerInitializerHandleFunc struct {
	defaultHook func(context.Context, []events.KinesisEventRecord, log.Logger) error
	hooks       []func(context.Context, []events.KinesisEventRecord, log.Logger) error
	history     []KinesisEventHandlerInitializerHandleFuncCall
	mutex       sync.Mutex
}

// Handle delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockKinesisEventHandlerInitializer) Handle(v0 context.Context, v1 []events.KinesisEventRecord, v2 log.Logger) error {
	r0 := m.HandleFunc.nextHook()(v0, v1, v2)
	m.HandleFunc.appendCall(KinesisEventHandlerInitializerHandleFuncCall{v0, v1, v2, r0})
	return r0
}

// SetDefaultHook sets function that is called when the Handle method of the
// parent MockKinesisEventHandlerInitializer instance is invoked and the
// hook queue is empty.
func (f *KinesisEventHandlerInitializerHandleFunc) SetDefaultHook(hook func(context.Context, []events.KinesisEventRecord, log.Logger) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Handle method of the parent MockKinesisEventHandlerInitializer instance
// invokes the hook at the front of the queue and discards it. After the
// queue is empty, the default hook function is invoked for any future
// action.
func (f *KinesisEventHandlerInitializerHandleFunc) PushHook(hook func(context.Context, []events.KinesisEventRecord, log.Logger) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *KinesisEventHandlerInitializerHandleFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(context.Context, []events.KinesisEventRecord, log.Logger) error {
		return r0
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *KinesisEventHandlerInitializerHandleFunc) PushReturn(r0 error) {
	f.PushHook(func(context.Context, []events.KinesisEventRecord, log.Logger) error {
		return r0
	})
}

func (f *KinesisEventHandlerInitializerHandleFunc) nextHook() func(context.Context, []events.KinesisEventRecord, log.Logger) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *KinesisEventHandlerInitializerHandleFunc) appendCall(r0 KinesisEventHandlerInitializerHandleFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of
// KinesisEventHandlerInitializerHandleFuncCall objects describing the
// invocations of this function.
func (f *KinesisEventHandlerInitializerHandleFunc) History() []KinesisEventHandlerInitializerHandleFuncCall {
	f.mutex.Lock()
	history := make([]KinesisEventHandlerInitializerHandleFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// KinesisEventHandlerInitializerHandleFuncCall is an object that describes
// an invocation of method Handle on an instance of
// MockKinesisEventHandlerInitializer.
type KinesisEventHandlerInitializerHandleFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 []events.KinesisEventRecord
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 log.Logger
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c KinesisEventHandlerInitializerHandleFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c KinesisEventHandlerInitializerHandleFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// KinesisEventHandlerInitializerInitFunc describes the behavior when the
// Init method of the parent MockKinesisEventHandlerInitializer instance is
// invoked.
type KinesisEventHandlerInitializerInitFunc struct {
	defaultHook func(config.Config) error
	hooks       []func(config.Config) error
	history     []KinesisEventHandlerInitializerInitFuncCall
	mutex       sync.Mutex
}

// Init delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockKinesisEventHandlerInitializer) Init(v0 config.Config) error {
	r0 := m.InitFunc.nextHook()(v0)
	m.InitFunc.appendCall(KinesisEventHandlerInitializerInitFuncCall{v0, r0})
	return r0
}

// SetDefaultHook sets function that is called when the Init method of the
// parent MockKinesisEventHandlerInitializer instance is invoked and the
// hook queue is empty.
func (f *KinesisEventHandlerInitializerInitFunc) SetDefaultHook(hook func(config.Config) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Init method of the parent MockKinesisEventHandlerInitializer instance
// invokes the hook at the front of the queue and discards it. After the
// queue is empty, the default hook function is invoked for any future
// action.
func (f *KinesisEventHandlerInitializerInitFunc) PushHook(hook func(config.Config) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *KinesisEventHandlerInitializerInitFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(config.Config) error {
		return r0
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *KinesisEventHandlerInitializerInitFunc) PushReturn(r0 error) {
	f.PushHook(func(config.Config) error {
		return r0
	})
}

func (f *KinesisEventHandlerInitializerInitFunc) nextHook() func(config.Config) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *KinesisEventHandlerInitializerInitFunc) appendCall(r0 KinesisEventHandlerInitializerInitFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of KinesisEventHandlerInitializerInitFuncCall
// objects describing the invocations of this function.
func (f *KinesisEventHandlerInitializerInitFunc) History() []KinesisEventHandlerInitializerInitFuncCall {
	f.mutex.Lock()
	history := make([]KinesisEventHandlerInitializerInitFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// KinesisEventHandlerInitializerInitFuncCall is an object that describes an
// invocation of method Init on an instance of
// MockKinesisEventHandlerInitializer.
type KinesisEventHandlerInitializerInitFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 config.Config
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c KinesisEventHandlerInitializerInitFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c KinesisEventHandlerInitializerInitFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// MockKinesisRecordHandlerInitializer is a mock implementation of the
// kinesisRecordHandlerInitializer interface (from the package
// github.com/go-nacelle/lambdabase) used for unit testing.
type MockKinesisRecordHandlerInitializer struct {
	// HandleFunc is an instance of a mock function object controlling the
	// behavior of the method Handle.
	HandleFunc *KinesisRecordHandlerInitializerHandleFunc
	// InitFunc is an instance of a mock function object controlling the
	// behavior of the method Init.
	InitFunc *KinesisRecordHandlerInitializerInitFunc
}

// NewMockKinesisRecordHandlerInitializer creates a new mock of the
// kinesisRecordHandlerInitializer interface. All methods return zero values
// for all results, unless overwritten.
func NewMockKinesisRecordHandlerInitializer() *MockKinesisRecordHandlerInitializer {
	return &MockKinesisRecordHandlerInitializer{
		HandleFunc: &KinesisRecordHandlerInitializerHandleFunc{
			defaultHook: func(context.Context, events.KinesisEventRecord, log.Logger) error {
				return nil
			},
		},
		InitFunc: &KinesisRecordHandlerInitializerInitFunc{
			defaultHook: func(config.Config) error {
				return nil
			},
		},
	}
}

// surrogateMockKinesisRecordHandlerInitializer is a copy of the
// kinesisRecordHandlerInitializer interface (from the package
// github.com/go-nacelle/lambdabase). It is redefined here as it is
// unexported in the source package.
type surrogateMockKinesisRecordHandlerInitializer interface {
	Handle(context.Context, events.KinesisEventRecord, log.Logger) error
	Init(config.Config) error
}

// NewMockKinesisRecordHandlerInitializerFrom creates a new mock of the
// MockKinesisRecordHandlerInitializer interface. All methods delegate to
// the given implementation, unless overwritten.
func NewMockKinesisRecordHandlerInitializerFrom(i surrogateMockKinesisRecordHandlerInitializer) *MockKinesisRecordHandlerInitializer {
	return &MockKinesisRecordHandlerInitializer{
		HandleFunc: &KinesisRecordHandlerInitializerHandleFunc{
			defaultHook: i.Handle,
		},
		InitFunc: &KinesisRecordHandlerInitializerInitFunc{
			defaultHook: i.Init,
		},
	}
}

// KinesisRecordHandlerInitializerHandleFunc describes the behavior when the
// Handle method of the parent MockKinesisRecordHandlerInitializer instance
// is invoked.
type KinesisRecordHandlerInitializerHandleFunc struct {
	defaultHook func(context.Context, events.KinesisEventRecord, log.Logger) error
	hooks       []func(context.Context, events.KinesisEventRecord, log.Logger) error
	history     []KinesisRecordHandlerInitializerHandleFuncCall
	mutex       sync.Mutex
}

// Handle delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockKinesisRecordHandlerInitializer) Handle(v0 context.Context, v1 events.KinesisEventRecord, v2 log.Logger) error {
	r0 := m.HandleFunc.nextHook()(v0, v1, v2)
	m.HandleFunc.appendCall(KinesisRecordHandlerInitializerHandleFuncCall{v0, v1, v2, r0})
	return r0
}

// SetDefaultHook sets function that is called when the Handle method of the
// parent MockKinesisRecordHandlerInitializer instance is invoked and the
// hook queue is empty.
func (f *KinesisRecordHandlerInitializerHandleFunc) SetDefaultHook(hook func(context.Context, events.KinesisEventRecord, log.Logger) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Handle method of the parent MockKinesisRecordHandlerInitializer instance
// invokes the hook at the front of the queue and discards it. After the
// queue is empty, the default hook function is invoked for any future
// action.
func (f *KinesisRecordHandlerInitializerHandleFunc) PushHook(hook func(context.Context, events.KinesisEventRecord, log.Logger) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *KinesisRecordHandlerInitializerHandleFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(context.Context, events.KinesisEventRecord, log.Logger) error {
		return r0
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *KinesisRecordHandlerInitializerHandleFunc) PushReturn(r0 error) {
	f.PushHook(func(context.Context, events.KinesisEventRecord, log.Logger) error {
		return r0
	})
}

func (f *KinesisRecordHandlerInitializerHandleFunc) nextHook() func(context.Context, events.KinesisEventRecord, log.Logger) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *KinesisRecordHandlerInitializerHandleFunc) appendCall(r0 KinesisRecordHandlerInitializerHandleFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of
// KinesisRecordHandlerInitializerHandleFuncCall objects describing the
// invocations of this function.
func (f *KinesisRecordHandlerInitializerHandleFunc) History() []KinesisRecordHandlerInitializerHandleFuncCall {
	f.mutex.Lock()
	history := make([]KinesisRecordHandlerInitializerHandleFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// KinesisRecordHandlerInitializerHandleFuncCall is an object that describes
// an invocation of method Handle on an instance of
// MockKinesisRecordHandlerInitializer.
type KinesisRecordHandlerInitializerHandleFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 events.KinesisEventRecord
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 log.Logger
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c KinesisRecordHandlerInitializerHandleFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c KinesisRecordHandlerInitializerHandleFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// KinesisRecordHandlerInitializerInitFunc describes the behavior when the
// Init method of the parent MockKinesisRecordHandlerInitializer instance is
// invoked.
type KinesisRecordHandlerInitializerInitFunc struct {
	defaultHook func(config.Config) error
	hooks       []func(config.Config) error
	history     []KinesisRecordHandlerInitializerInitFuncCall
	mutex       sync.Mutex
}

// Init delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockKinesisRecordHandlerInitializer) Init(v0 config.Config) error {
	r0 := m.InitFunc.nextHook()(v0)
	m.InitFunc.appendCall(KinesisRecordHandlerInitializerInitFuncCall{v0, r0})
	return r0
}

// SetDefaultHook sets function that is called when the Init method of the
// parent MockKinesisRecordHandlerInitializer instance is invoked and the
// hook queue is empty.
func (f *KinesisRecordHandlerInitializerInitFunc) SetDefaultHook(hook func(config.Config) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Init method of the parent MockKinesisRecordHandlerInitializer instance
// invokes the hook at the front of the queue and discards it. After the
// queue is empty, the default hook function is invoked for any future
// action.
func (f *KinesisRecordHandlerInitializerInitFunc) PushHook(hook func(config.Config) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *KinesisRecordHandlerInitializerInitFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(config.Config) error {
		return r0
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *KinesisRecordHandlerInitializerInitFunc) PushReturn(r0 error) {
	f.PushHook(func(config.Config) error {
		return r0
	})
}

func (f *KinesisRecordHandlerInitializerInitFunc) nextHook() func(config.Config) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *KinesisRecordHandlerInitializerInitFunc) appendCall(r0 KinesisRecordHandlerInitializerInitFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of KinesisRecordHandlerInitializerInitFuncCall
// objects describing the invocations of this function.
func (f *KinesisRecordHandlerInitializerInitFunc) History() []KinesisRecordHandlerInitializerInitFuncCall {
	f.mutex.Lock()
	history := make([]KinesisRecordHandlerInitializerInitFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// KinesisRecordHandlerInitializerInitFuncCall is an object that describes
// an invocation of method Init on an instance of
// MockKinesisRecordHandlerInitializer.
type KinesisRecordHandlerInitializerInitFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 config.Config
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c KinesisRecordHandlerInitializerInitFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c KinesisRecordHandlerInitializerInitFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}