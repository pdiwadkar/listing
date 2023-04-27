package main

import (
	"fmt"
	"log"
)

func main() {

	server := NewRestServer()
	asClient, err := server.Init()
	if err != nil {
		log.Fatal(err)
	}
	server.ListenAndServe()
	fmt.Println(asClient.GetNodeNames())

}
