package admin

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthchck(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(healthcheck))
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Fatalf("Error querying admin endpoint, %s", err.Error())
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Error reading in response body, %s", err.Error())
	}

	buf := bytes.NewBuffer(respBody)
	if buf.String() != "pong" {
		t.Fatal(`admin healthcheck should return "pong"`)
	}
}
