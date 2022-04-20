package main

import (
	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
)

// it represents to config parameters into the config.yml
type Config struct {
	Apikey string
}

func New() interface{} {
	return &Config{}
}

var Version = "0.3"
var Priority = 1

func main() {
	server.StartServer(New, Version, Priority)
}

func (conf Config) Access(kong *pdk.PDK) {
	key, err := kong.Request.GetQueryArg("key")
	apiKey := conf.Apikey

	if err != nil {
		kong.Log.Err(err.Error())
	}

	//it adjusts the header parameters in this way.
	x := make(map[string][]string)
	x["Content-Type"] = append(x["Content-Type"], "application/json")

	//If the key of the consumer is not equal to the claimed key, kong doesn't ensure the proxy
	if apiKey != key {
		kong.Response.Exit(403, "forbidden", x)
	}
}
