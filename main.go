package main

import (
	"flag"
	"log"

	"github.com/sebzz2k2/krompton/config"
	"github.com/sebzz2k2/krompton/server"
)

func setupFlags() {
	log.Println("Running at", config.Host, config.Port)
	flag.StringVar(&config.Host, "host", "0.0.0.0", "Host")
	flag.IntVar(&config.Port, "port", 5381, "Port")
	flag.Parse()
}
func main() {
	log.Println("Starting your Krompton server")
	setupFlags()
	server.RunSyncTcp()
}
