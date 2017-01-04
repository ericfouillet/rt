// Simple REST client

package rt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// New creates a new request from the provided arguments.
// The payload (optional) is provided as a string.
func New(endpoint, method, payload string) *Request {
	return &Request{endpoint, method, payload}
}

// NewWithFile creates a new request from the provided arguments.
// The payload (optional) is provided as a filename.
func NewWithFile(endpoint, method, payloadFile string) (*Request, error) {
	payload, err := ioutil.ReadFile(payloadFile)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not read file %v", payloadFile)
	}
	return &Request{endpoint, method, string(payload)}, nil
}

// Execute sends the REST request and returns the response.
func (r *Request) Execute() (string, interface{}, error) {
	var body bytes.Buffer
	_, err := body.Write([]byte(r.payload))
	if err != nil {
		return "", nil, err
	}
	req, err := http.NewRequest(r.method, r.endpoint, &body)
	if err != nil {
		return "", nil, err
	}
	req.Header.Set("User-Agent", "rt")
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return resp.Status, nil, fmt.Errorf("Request was not successful: %v", resp.Status)
	}
	var t interface{}
	err = json.NewDecoder(resp.Body).Decode(&t)
	if err != nil {
		return "", nil, err
	}
	return resp.Status, t, nil
}
