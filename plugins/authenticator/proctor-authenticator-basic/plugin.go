package main

import (
	"encoding/base64"
	"os"

	authenticator "github.com/bee-projects/proctor/plugins/authenticator"
	"github.com/hashicorp/go-hclog"
)

// BasicAuth ...
type BasicAuth struct {
	logger hclog.Logger
}

// Authenticate returns a basic auth header srting given a username and password
// Authenticate The Authentication implementation for this plugin.
func (a *BasicAuth) Authenticate(dataIn map[string]string) (dataOut map[string]string) {

	// Get username and password from dataIn
	username := dataIn["user"]
	password := dataIn["pass"]
	auth := username + ":" + password

	// Create auth header and add to dataOut
	authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	dataOut = make(map[string]string)
	dataOut["Authorization"] = authHeader
	return dataOut
}

// Launch Plugin
func main() {

	// Instantiate your Authenticator Plugin
	authPlugin := &BasicAuth{
		logger: hclog.New(&hclog.LoggerOptions{
			Level:      hclog.Trace,
			Output:     os.Stderr,
			JSONFormat: true,
		}),
	}

	// Launch the plugin
	authenticator.LaunchPlugin(authPlugin, "authenticator")

}
