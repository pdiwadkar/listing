package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/pdiwadkar/listing/blockcodes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50001"
)

func main() {
	//bconfig := grpc.BackoffConfig{MaxDelay:time.Second*2,}
	/*retryPolicy := `{
		"method_config": [{
			"name" :[{"service":""}],
			"waitForReady" :true,
			"retryPolicy" : {
				"MaxAttempts":3
			}
		}
		]
	}`*/
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	cli := pb.NewBlockCodeTransactionClient(conn)
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()
	IdTxn, err := cli.GetTransaction(ctx, &pb.Txn{Bin: "405028", Logo: "400", Code: "P"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(IdTxn)
}
