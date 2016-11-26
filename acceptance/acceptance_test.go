package acceptance

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"testing"
)

var api string

func TestMain(m *testing.M) {
	flag.StringVar(&api, "url", "", "url to run acceptance tests")
	flag.Parse()

	os.Exit(m.Run())
}

func TestServerExists(t *testing.T) {
	resp, err := http.Get(api)
	checkHttpErr(err, t)
	checkHttpCode(resp, 200, t)
}

func TestAddingAndGettingMessage(t *testing.T) {
	user := "sam"
	messageEndpoint := fmt.Sprintf("%s/users/%s/message", api, user)

	msg := message{
		Text: "foo bar",
	}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(msg)

	resp, err := http.Post(messageEndpoint, "application/json", b)
	checkHttpErr(err, t)
	checkHttpCode(resp, 201, t)

	resp, err = http.Get(messageEndpoint)
	checkHttpErr(err, t)
	checkHttpCode(resp, 200, t)

	var responseMsgs []message
	err = json.NewDecoder(resp.Body).Decode(&responseMsgs)

	if reflect.DeepEqual(responseMsgs, []message{msg}) == false {
		t.Fatalf("Repsonses not equal: %+v vs %+v", msg, responseMsgs)
	}

}

type message struct {
	Text string `json:message`
}

func checkHttpErr(err error, t *testing.T) {
	if err != nil {
		t.Fatalf("error contacting server: %v", err)
	}
}

func checkHttpCode(resp *http.Response, expectedCode int, t *testing.T) {
	if resp.StatusCode != expectedCode {
		t.Fatalf("expected %d code, got code: %d", expectedCode, resp.StatusCode)
	}
}
