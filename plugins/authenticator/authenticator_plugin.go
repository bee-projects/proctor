package commons

import (
	"log"
	"net/rpc"
	"os"
	"os/exec"

	"github.com/hashicorp/go-hclog"
	plugin "github.com/hashicorp/go-plugin"
)

var (
	pluginMap = map[string]plugin.Plugin{
		"authenticator": &AuthenticatorPlugin{},
	}
)

// AuthenticatorPlugin has two methods:
// Server must return an RPC server for this plugin type.
// We construct a AuthenticatorRPCServer for this.
//
// Client must return an implementation of our interface that communicates
// over an RPC client. We return AuthenticatorRPC for this.
type AuthenticatorPlugin struct {
	// Impl injection
	Impl Authenticator
}

// Server is the RPC server that serves the Authenticator Plugin
func (p *AuthenticatorPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &AuthenticatorRPCServer{Impl: p.Impl}, nil
}

// Client is the RPC client that communicated with the RPC server
func (AuthenticatorPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &AuthenticatorRPC{client: c}, nil
}

// LaunchPlugin launches a plugin
func LaunchPlugin(authenticator Authenticator, name string) {
	// Create an plugin with the provided implementation
	myplugin := &AuthenticatorPlugin{
		Impl: authenticator,
	}

	// Define a map of plugins we can dispense.
	// The name of the plugin will be used by the host process
	var pluginMap = map[string]plugin.Plugin{
		name: myplugin,
	}

	// Run the plugin
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: HandshakeConfig,
		Plugins:         pluginMap,
	})
}

// DispenseAuthenticator loads an Authenticator plug-in into memory
func DispenseAuthenticator(pluginPath string, level hclog.Level) (Authenticator, *plugin.Client) {

	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "plugin",
		Output: os.Stdout,
		Level:  hclog.Info,
	})

	// Create a plugin client
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: HandshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command(pluginPath),
		Logger:          logger,
	})

	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
	}

	// Request an authenticator client from the plugin
	raw, err := rpcClient.Dispense("authenticator")
	if err != nil {
		log.Fatal(err)
	}
	authenticator := raw.(Authenticator)

	// Return the autenticator client and plugin client
	return authenticator, client
}
