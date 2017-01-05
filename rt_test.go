package rt

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var methods = [...]string{http.MethodPut, http.MethodPost, http.MethodDelete}

func TestGet(t *testing.T) {
	ts := newTestServer()
	defer ts.Close()
	rt := &Request{ts.URL, http.MethodGet, ""}
	status, resp, err := rt.Execute()
	if !assert.NoError(t, err, "request should have been successful") {
		return
	}
	if !assert.Contains(t, status, "200", "Status should be 200") {
		return
	}
	if !assert.NotNil(t, resp, "Response should not be empty") {
		return
	}
}

func TestOther(t *testing.T) {
	ts := newTestServer()
	defer ts.Close()
	for _, m := range methods {
		t.Run("method="+m, func(t *testing.T) {
			testMethod(t, ts.URL, m, "{'request': 'my request', 'param': 124}")
		})
	}
}

func testMethod(t *testing.T, url, method, payload string) {
	rt := &Request{url, method, payload}
	status, resp, err := rt.Execute()
	if !assert.NoError(t, err, "request should have been successful") {
		return
	}
	if !assert.Contains(t, status, "200", "Status should be 200") {
		return
	}
	if !assert.NotNil(t, resp, "Response should not be empty") {
		return
	}
}

func newTestServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := struct {
			Response string
		}{"test"}
		json.NewEncoder(w).Encode(&t)
	}))
}
