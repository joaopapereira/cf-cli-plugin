package main

import (
	"code.cloudfoundry.org/cli/plugin"
	"log"
)

func main() {
	con := plugin.NewCliConnection("10")
	version, err := con.ApiVersion()
	if err != nil {
		log.Fatal("Failed to get API version")
	}
	log.Printf("API version: %s", version)
}
