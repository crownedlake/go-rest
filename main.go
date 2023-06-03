package main

import (
	_ "github.com/lib/pq"
	"goProject/api"
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	apiHandler, err := api.GetNewAPIHandler()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", apiHandler.HandlePaths)

	log.Println("attempting go server, port:", port)
	err = http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatal(err)
	}
}
