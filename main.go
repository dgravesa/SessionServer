package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/dgravesa/SessionServer/controller"
	"github.com/dgravesa/SessionServer/data"
	"github.com/dgravesa/SessionServer/model"

	_ "github.com/lib/pq"
)

var portNumber = flag.Int("port", 11000, "the port number to listen on")
var configName = flag.String("config", "config.yml", "path to the configuration file")

func main() {
	flag.Parse()

	controller.RegisterAll()
	log.Println("Request handlers registered.")

	log.Println("Initializing SQL data layer...")
	dblayer, err := data.NewDBLayer(*configName)
	if err != nil {
		log.Fatal(err)
	}

	model.SetData(dblayer)
	log.Println("SQL data layer intialized.")

	log.Printf("Listening on port %d\n", *portNumber)
	portStr := fmt.Sprintf(":%d", *portNumber)
	http.ListenAndServe(portStr, nil)
}
