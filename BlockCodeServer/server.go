package main

import (
	"context"
	"fmt"
	"log"
	"net"

	as "github.com/aerospike/aerospike-client-go/v6"
	"github.com/aerospike/aerospike-client-go/v6/types"
	pb "github.com/pdiwadkar/listing/blockcodes"
	"google.golang.org/grpc"
)

const (
	port = ":50001"
)

/*
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative .\blockcodes\blockcode.proto
*/

type CodeServer struct {
	pb.UnimplementedBlockCodeTransactionServer
	client *as.Client
}

func NewCodeServer() *CodeServer {
	client, err := ConnectAeroCache()
	fmt.Println("New code server  ", client != nil)
	if err != nil {
		log.Fatal("Unable to connect to Aerocache.")
	}
	server := CodeServer{client: client}
	return &server
}

func (cs *CodeServer) GetTransaction(ctx context.Context, in *pb.Txn) (*pb.IdTxn, error) {

	log.Println("Received:: ", in.Bin)
	queryPolicy := as.NewQueryPolicy()
	queryPolicy.MaxRecords = 5
	queryPolicy.FilterExpression = as.ExpLet(
		as.ExpDef("bin", as.ExpStringBin("BIN")),
		as.ExpDef("logo", as.ExpStringBin("LOGO")),
		as.ExpDef("code", as.ExpStringBin("BLOCKDOE")),
		as.ExpAnd(
			as.ExpEq(
				as.ExpVar("bin"), as.ExpStringVal(in.Bin)),
			as.ExpEq(as.ExpVar("logo"), as.ExpStringVal(in.Logo)),
			as.ExpEq(as.ExpVar("code"), as.ExpStringVal(in.Code)),
		),
	)
	stmt := as.NewStatement("hdfc", "cctxn")
	fmt.Println("clint ", cs.client != nil)
	fmt.Println("CS  is ", cs != nil)
	//client, err := as.NewClient("127.0.0.1", 3000)

	result, err := cs.client.Query(queryPolicy, stmt)
	if err != nil {
		log.Fatal(err)
	}
	var tx string
	for recs := range result.Results() {
		if recs != nil {
			fmt.Println(recs.Record.Bins["IDTXN"])
			tx = "2002"
		}
	}
	return &pb.IdTxn{Bin: in.Bin, Logo: in.Logo, Code: in.Code, Txn: tx}, nil
}
func main() {

	codeServer := NewCodeServer()
	codeServer.ListenAndServe(codeServer.client)
}
func (cs *CodeServer) ListenAndServe(client *as.Client) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	ser := grpc.NewServer()
	pb.RegisterBlockCodeTransactionServer(ser, cs)
	log.Println("Server listening at ", lis.Addr())
	if err := ser.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
func ConnectAeroCache() (*as.Client, error) {
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
func (cs *CodeServer) Shutdown() {
	//shutdown using os.Signal
}
