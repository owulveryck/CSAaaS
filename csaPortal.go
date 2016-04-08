package main

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/owulveryck/CSAaaS/server"
	"log"
	"net/http"
)

type Configuration struct {
	Debug       bool
	Scheme      string
	Port        int
	Address     string
	CsaBackend  string
	PrivateKey  string
	Certificate string
}

var config Configuration

func main() {
	// Default values
	config.Port = 8080
	config.Scheme = "https"
	config.Address = "0.0.0.0"
	config.Debug = false
	config.CsaBackend = "https://localhost:8888/csa"
	config.PrivateKey = "ssl/server.key"
	config.Certificate = "ssl/server.pem"

	err := envconfig.Process("csaportal", &config)
	if err != nil {
		log.Fatal(err.Error())

	}
	format := "Infos:\n  Debug: %v\n  URL: %s://%s:%d\n  CSA backend: %s\n"
	_, err = fmt.Printf(format, config.Debug, config.Scheme, config.Address, config.Port, config.CsaBackend)
	if err != nil {
		log.Fatal(err.Error())

	}
	router := server.NewRouter()

	addr := fmt.Sprintf("%v:%v", config.Address, config.Port)
	if config.Scheme == "https" {
		log.Fatal(http.ListenAndServeTLS(addr, config.Certificate, config.PrivateKey, router))

	} else {
		log.Fatal(http.ListenAndServe(addr, router))
	}
}
