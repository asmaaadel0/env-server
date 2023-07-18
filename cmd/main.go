package main

import (
	"flag"
	"log"

	"github.com/codescalersinternships/envserver-Asmaa/internal"
)

func main() {

	var port int
	flag.IntVar(&port, "p", 8080, "port")
	flag.Parse()

	app, err := internal.NewApp(port)
	if err != nil {
		log.Fatalf("Error creating App:%s", internal.ErrorOutOfRange)
	}

	app.Run()
}
