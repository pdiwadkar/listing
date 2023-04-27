package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	as "github.com/aerospike/aerospike-client-go/v6"
)

type sqSet struct {
	input string
	outut string
}

var cn chan Txn
var rs chan []IdTxn

func init() {
	cn = make(chan Txn)
	rs = make(chan []IdTxn)
	fmt.Println("Channel set")
}
func GetValidTransactions(w http.ResponseWriter, req *http.Request) {
	//query
	//ProcessRequest(w, req)
	var txn Txn
	body, _ := ioutil.ReadAll(req.Body)
	//fmt.Println("Body ", string(body))
	json.Unmarshal(body, &txn)
	go WorkerAerospike()
	cn <- txn
	arr := <-rs
	fmt.Println(txn)
	fmt.Println("Total :: ", len(arr))
	for _, tx := range arr {
		fmt.Println(tx.Bin, "  ", tx.Code, "  ", tx.Logo, "   ", tx.Txn)
	}
	fmt.Println("##########################################")
}

func WorkerAerospike() {
	txn := <-cn
	fmt.Println("Transaction received...")
	client := AeroClient
	if client == nil {
		log.Fatal("Error in getting aerospike connection")
	}
	queryPolicy := as.NewQueryPolicy()
	queryPolicy.MaxRecords = 5
	queryPolicy.FilterExpression = as.ExpLet(
		as.ExpDef("bin", as.ExpStringBin("BIN")),
		as.ExpDef("logo", as.ExpStringBin("LOGO")),
		as.ExpDef("code", as.ExpStringBin("BLOCKDOE")),
		as.ExpAnd(
			as.ExpEq(
				as.ExpVar("bin"), as.ExpStringVal(txn.Bin)),
			as.ExpEq(as.ExpVar("logo"), as.ExpStringVal(txn.Logo)),
			as.ExpEq(as.ExpVar("code"), as.ExpStringVal(txn.Code)),
		),
	)
	stmt := as.NewStatement("hdfc", "cctxn")
	result, err := AeroClient.Query(queryPolicy, stmt)
	if err != nil {
		log.Fatal(err)
	}
	resultArr := make([]IdTxn, 0)
	for recs := range result.Results() {
		fmt.Println("POSITIVE")
		if recs.Err == nil {
			bn := fmt.Sprint(recs.Record.Bins["BIN"])
			lg := fmt.Sprint(recs.Record.Bins["LOGO"])
			cd := fmt.Sprint(recs.Record.Bins["BLOCKDOE"])
			idtx := fmt.Sprint(recs.Record.Bins["IDTXN"])
			resultArr = append(resultArr, IdTxn{Bin: bn, Logo: lg, Code: cd, Txn: idtx})
		}
	}
	rs <- resultArr
}

func ProcessRequest(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Request received.")
	client := AeroClient
	if client == nil {
		log.Fatal("Error in getting aerospike connection")
	}
	var txn Txn
	body, _ := ioutil.ReadAll(req.Body)
	fmt.Println("Body ", body)
	json.Unmarshal(body, &txn)
	fmt.Println(txn)
	//result := QueryForTransactions(txn)
	//result := "{\"key\":\"Hello\"}"
	queryPolicy := as.NewQueryPolicy()
	queryPolicy.MaxRecords = 5
	queryPolicy.FilterExpression = as.ExpLet(
		as.ExpDef("bin", as.ExpStringBin("BIN")),
		as.ExpDef("logo", as.ExpStringBin("LOGO")),
		as.ExpDef("code", as.ExpStringBin("BLOCKDOE")),
		as.ExpAnd(
			as.ExpEq(
				as.ExpVar("bin"), as.ExpStringVal(txn.Bin)),
			as.ExpEq(as.ExpVar("logo"), as.ExpStringVal(txn.Logo)),
			as.ExpEq(as.ExpVar("code"), as.ExpStringVal(txn.Code)),
		),
	)
	stmt := as.NewStatement("hdfc", "cctxn")
	result, err := AeroClient.Query(queryPolicy, stmt)
	if err != nil {
		log.Fatal(err)
	}
	resultArr := make([]IdTxn, 0)
	for recs := range result.Results() {

		if recs.Err != nil {
			bn := fmt.Sprint(recs.Record.Bins["BIN"])
			lg := fmt.Sprint(recs.Record.Bins["LOGO"])
			cd := fmt.Sprint(recs.Record.Bins["BLOCKDOE"])
			idtx := fmt.Sprint(recs.Record.Bins["IDTXN"])
			resultArr = append(resultArr, IdTxn{Bin: bn, Logo: lg, Code: cd, Txn: idtx})

		}
	}

	//---------------------------------------------
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultArr)

}

func QueryForTransactions(txn Txn) []IdTxn {
	queryPolicy := as.NewQueryPolicy()
	queryPolicy.MaxRecords = 5
	queryPolicy.FilterExpression = as.ExpLet(
		as.ExpDef("bin", as.ExpStringBin("BIN")),
		as.ExpDef("logo", as.ExpStringBin("LOGO")),
		as.ExpDef("code", as.ExpStringBin("BLOCKDOE")),
		as.ExpAnd(
			as.ExpEq(
				as.ExpVar("bin"), as.ExpStringVal(txn.Bin)),
			as.ExpEq(as.ExpVar("logo"), as.ExpStringVal(txn.Logo)),
			as.ExpEq(as.ExpVar("code"), as.ExpStringVal(txn.Code)),
		),
	)
	stmt := as.NewStatement("hdfc", "cctxn")
	result, err := AeroClient.Query(queryPolicy, stmt)
	if err != nil {
		log.Fatal(err)
	}
	resultArr := make([]IdTxn, 0)
	for recs := range result.Results() {
		if recs != nil {
			bn := fmt.Sprint(recs.Record.Bins["BIN"])
			lg := fmt.Sprint(recs.Record.Bins["LOGO"])
			cd := fmt.Sprint(recs.Record.Bins["BLOCKDOE"])
			idtx := fmt.Sprint(recs.Record.Bins["IDTXN"])
			resultArr = append(resultArr, IdTxn{Bin: bn, Logo: lg, Code: cd, Txn: idtx})

		}
	}
	return resultArr
}
