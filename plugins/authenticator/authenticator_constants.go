package authenticator

import (
	plugin "github.com/hashicorp/go-plugin"
)

var (
	// HandshakeConfig is prevents users from executing bad plugins
	// from the command line and be launched as a plugin only.
	// It is a UX feature, not a security feature.
	HandshakeConfig = plugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "BASIC_PLUGIN",
		MagicCookieValue: "hello",
	}
)
