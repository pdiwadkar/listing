package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	as "github.com/aerospike/aerospike-client-go/v6"
	"github.com/aerospike/aerospike-client-go/v6/types"
	"github.com/gorilla/mux"
)

var AeroClient *as.Client
var err error

type RestServer struct {
}

func NewRestServer() *RestServer {
	return &RestServer{}
}
func (rs *RestServer) Shutdown() {
	if AeroClient != nil {
		AeroClient.Close()
		fmt.Println("Aerospike connection closed.")
	}
}

func (rs *RestServer) Init() (*as.Client, error) {
	//code to initialize Aerospike client.
	AeroClient, err = rs.connectAeroCache()
	if err != nil {
		return nil, err
	}
	if AeroClient != nil {
		return AeroClient, nil
	}
	return nil, errors.New("init failed.")
}
func (rs *RestServer) ListenAndServe() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/check", GetValidTransactions)

	log.Fatal(http.ListenAndServe(":12000", router))
}

func (rs *RestServer) connectAeroCache() (*as.Client, error) {
	/*
		Connect to Aerospike server.
		if Error-network or node not available.. return the error
		if connected return client.
	*/
	client, err := as.NewClient("127.0.0.1", 3000)
	if err != nil {
		if err.Matches(types.NETWORK_ERROR) || err.Matches(types.INVALID_NODE_ERROR) {
			fmt.Println("Aerospike connectiviy error:: try after sometime ")
			return nil, err
		}
	}
	return client, nil
}
