package main

import (
	"github.com/chinaboard/habada/service"
	"log"
	"net/http"
)

func main() {
	routersInit := service.InitRouter()

	server := &http.Server{
		Addr:    ":3333",
		Handler: routersInit,
	}

	log.Fatalln(server.ListenAndServe())

}
