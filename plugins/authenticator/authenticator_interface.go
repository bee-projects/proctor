package authenticator

// Authenticator is the interface exposed as a plugin
type Authenticator interface {
	//Authenticate(interface{}) map[string]string
	Authenticate(map[string]string) map[string]string
}
