package gateway_test

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"log"
	"net/http"
	"testing"

	"github.com/Drafteame/gateway/v2"
	"github.com/tj/assert"
)

func Example() {
	http.HandleFunc("/", hello)
	log.Fatal(gateway.ListenAndServe(nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "Hello World from Go")
}

func TestGateway_Invoke(t *testing.T) {
	evt := events.APIGatewayProxyRequest{
		Path:       "/pets/luna",
		HTTPMethod: "POST",
	}

	gw := gateway.NewGateway(http.HandlerFunc(hello))

	payload, err := gw.Invoke(context.Background(), evt)
	assert.NoError(t, err)
	assert.JSONEq(t, `{"body":"Hello World from Go\n", "headers":{"Content-Type":"text/plain; charset=utf8"}, "multiValueHeaders":{}, "statusCode":200}`, string(payload.([]byte)))
}
