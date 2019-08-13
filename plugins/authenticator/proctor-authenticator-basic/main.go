package main

import (
	"fmt"

	hclog "github.com/hashicorp/go-hclog"
	plugin "github.com/hashicorp/go-plugin"

	authenticator "github.com/bee-projects/proctor/plugins/authenticator"
)

var (
	pluginMap = map[string]plugin.Plugin{
		"authenticator": &authenticator.AuthenticatorPlugin{},
	}
)

func main() {

	// Get Authenticator client from plugin
	basicAuth, client := authenticator.DispenseAuthenticator("./plugin", hclog.Info)
	defer client.Kill()

	// Prepare to pass some data
	dataIn := make(map[string]string)
	dataIn["user"] = "user"
	dataIn["pass"] = "blahblahblah"

	// Invoke authenticator plugin and output response
	headers := basicAuth.Authenticate(dataIn)
	fmt.Printf("The response was: %v \n", headers["Authorization"])
}
