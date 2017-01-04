// Simple REST client

package rt

// Request is a REST request as specified by the user.
type Request struct {
	endpoint string
	method   string
	payload  string
}

// Executor executes REST requests.
type Executor interface {
	Execute()
}
