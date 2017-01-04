package rt

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := struct {
			Response string
		}{"test"}
		json.NewEncoder(w).Encode(&t)
	}))
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
