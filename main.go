package main

import (
	"fmt"
	"github.com/bee-projects/proctor/plugins/authenticator"
)

func main() {
	fmt.Println(authenticator.HandshakeConfig.MagicCookieKey)
}