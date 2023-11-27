package main

import (
	"flag"
	"log"
	"simple-server/api"
	"simple-server/storage"
)


func main() {
	listendAddr := flag.String("listenAddr", ":3000", "the server address")
	flag.Parse()

	store := storage.NewMemoryStorage()

	server := api.NewServer(*listendAddr, store)

	log.Println("Starting server on port: ", *listendAddr)

	if err := server.Start(); err != nil{
		log.Fatal(err)
	}
}