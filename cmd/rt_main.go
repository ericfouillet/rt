package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/ericfouillet/rt"
)

var payloadFile = flag.String("f", "", "specify a file containing the JSON payload you wish to use")
var payloadString = flag.String("p", "", "specify a string containing the JSON payload you wish to use")

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		usage()
		return
	}
	request, err := makeRequest(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, resp, err := request.Execute()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
}

// makeRequest creates a new request from command line arguments
func makeRequest(args []string) (*rt.Request, error) {
	endpoint := args[0]
	var method string
	switch args[1] {
	case http.MethodGet:
	case http.MethodPost:
	case http.MethodPut:
	case http.MethodDelete:
		method = args[1]
	default:
		return nil, fmt.Errorf("Unsupported method %v", args[1])
	}
	fmt.Println("arguments", endpoint, method, *payloadString, *payloadFile)
	if *payloadFile != "" && *payloadString != "" {
		return nil, fmt.Errorf("The -f and -p options are mutually exclusive. Use only one.")
	}
	if *payloadString != "" {
		return rt.New(endpoint, method, *payloadString), nil
	}
	if *payloadFile != "" {
		req, err := rt.NewWithFile(endpoint, method, *payloadFile)
		if err != nil {
			return nil, err
		}
		return req, nil
	}
	return rt.New(endpoint, method, ""), nil
}

func usage() {
	fmt.Print("Usage: rt <options> endpoint method\n",
		"\t<options>: optional - Use -f filename or -p <json string> to set a payload for the request\n",
		"\tendpoint: mandatory - the endpoint of the REST API\n",
		"\tmethod: mandatory - GET, POST, PUT, DELETE\n")
}
