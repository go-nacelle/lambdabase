package lambdabase

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/rpc"

	"github.com/aphistic/sweet"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambda/messages"
	"github.com/go-nacelle/nacelle"
	. "github.com/onsi/gomega"
)

type ServerSuite struct{}

var testConfig = nacelle.NewConfig(nacelle.NewTestEnvSourcer(map[string]string{
	"_lambda_server_port": "0",
}))

var testHandler = LambdaHandlerFunc(func(ctx context.Context, payload []byte) ([]byte, error) {
	data := []string{}
	if err := json.Unmarshal(payload, &data); err != nil {
		return nil, fmt.Errorf("malformed input")
	}

	for i, value := range data {
		data[i] = fmt.Sprintf("%s:%s", value, GetRequestID(ctx))
	}

	serialized, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return serialized, nil
})

func (s *ServerSuite) TestServeAndStop(t sweet.T) {
	server := makeLambdaServer(testHandler)
	err := server.Init(testConfig)
	Expect(err).To(BeNil())

	go server.Start()
	defer server.Stop()

	// Hack internals to get the dynamic port (don't bind to one on host)
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", getDynamicPort(server.listener)))
	Expect(err).To(BeNil())

	client := rpc.NewClient(conn)
	defer client.Close()

	request := &messages.InvokeRequest{
		Payload:   []byte(`["foo", "bar", "baz"]`),
		RequestId: "bonk",
	}

	response := &messages.InvokeResponse{}
	err = client.Call("Function.Invoke", request, &response)
	Expect(err).To(BeNil())
	Expect(string(response.Payload)).To(Equal(`["foo:bonk","bar:bonk","baz:bonk"]`))

	request = &messages.InvokeRequest{
		Payload:   []byte(`[123, 456, 789]`),
		RequestId: "bonk",
	}

	err = client.Call("Function.Invoke", request, &response)
	Expect(err).To(BeNil())
	Expect(response.Error.Message).To(Equal("malformed input"))
}

func (s *ServerSuite) TestBadInjection(t sweet.T) {
	server := NewServer(&badInjectionLambdaHandler{})
	server.Logger = nacelle.NewNilLogger()
	server.Services = makeBadContainer()
	server.Health = nacelle.NewHealth()

	err := server.Init(testConfig)
	Expect(err.Error()).To(ContainSubstring("ServiceA"))
}

func (s *ServerSuite) TestInitError(t sweet.T) {
	server := NewServer(&badInitLambdaHandler{})
	server.Logger = nacelle.NewNilLogger()
	server.Services = makeBadContainer()
	server.Health = nacelle.NewHealth()

	err := server.Init(testConfig)
	Expect(err).To(MatchError("oops"))
}

//
// Helpers

type wrappedHandler struct {
	lambda.Handler
}

func (h *wrappedHandler) Init(config nacelle.Config) error {
	return nil
}

func makeLambdaServer(handler lambda.Handler) *Server {
	server := NewServer(&wrappedHandler{Handler: handler})
	server.Logger = nacelle.NewNilLogger()
	server.Services = nacelle.NewServiceContainer()
	server.Health = nacelle.NewHealth()
	return server
}

func getDynamicPort(listener net.Listener) int {
	return listener.Addr().(*net.TCPAddr).Port
}

//
// Bad Injection

type A struct{ X int }
type B struct{ X float64 }

type badInjectionLambdaHandler struct {
	ServiceA *A `service:"A"`
}

func (i *badInjectionLambdaHandler) Init(nacelle.Config) error {
	return nil
}

func (i *badInjectionLambdaHandler) Invoke(context.Context, []byte) ([]byte, error) {
	return nil, nil
}

func makeBadContainer() nacelle.ServiceContainer {
	container := nacelle.NewServiceContainer()
	container.Set("A", &B{})
	return container
}

//
// Bad Init

type badInitLambdaHandler struct{}

func (i *badInitLambdaHandler) Init(nacelle.Config) error {
	return fmt.Errorf("oops")
}

func (i *badInitLambdaHandler) Invoke(context.Context, []byte) ([]byte, error) {
	return nil, nil
}
