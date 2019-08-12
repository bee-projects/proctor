package commons

import (
	"fmt"
	"net/rpc"
)

// AuthenticatorRPC is an Authenticator implementation that talks over RPC
type AuthenticatorRPC struct{ client *rpc.Client }

// Authenticate will call the Authenticate method of the plugin
func (a *AuthenticatorRPC) Authenticate(dataIn map[string]string) (dataOut map[string]string) {

	fmt.Println("Data received:", dataIn)
	err := a.client.Call("Plugin.Authenticate", &dataIn, &dataOut)
	if err != nil {
		panic(err)
	}
	return dataOut
}

// AuthenticatorRPCServer is the RPC server that AuthenticatorRPC talks to,  conforming to
// net/rpc reqjuirements
type AuthenticatorRPCServer struct {
	//This is the real implementation
	Impl Authenticator
}

// Authenticate ...
//func (s *AuthenticatorRPCServer) Authenticate(args interface{}, resp *map[string]string) error {
func (s *AuthenticatorRPCServer) Authenticate(args map[string]string, resp *map[string]string) error {
	*resp = s.Impl.Authenticate(args)
	return nil
}
