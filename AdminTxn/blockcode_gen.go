// program to generate random block code data for Unit testing.
// Only blockcode,logo,idapp,idtxn fields are considered as thers fields are
// not used.
/*
Docker command to run aerospike
docker run -tid --name aerospike -p 3000-3002:3000-3002 -v D:/Work/Aerospike/Keys:/aerospike/etc -e "FEAURE_KEY_FILE=/aerospike/etc/trial-features.conf" -e "NAMESPACE=hdfc" aerospike/aerospike-server-enterprise
a3e709fccabed63eaec539aa3825c1e409ba516a02b1d608abcf83c8af8c82d7
*/
package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	as "github.com/aerospike/aerospike-client-go/v6"
	"github.com/aerospike/aerospike-client-go/v6/types"
)

func main() {
	/*	fmt.Println("main")
		data := GenerateRandomList()
		for _, val := range data {
			fmt.Println(val)
		}*/
	//PrintTestRecords()
	FilterRecords()
}

func FilterRecords() {
	/*
		map[BIN:405028 BLOCKDOE:P IDAPP:RS IDTXN:ACI LOGO:400]
		map[BIN:524368 BLOCKDOE:B IDAPP:RS IDTXN:VST LOGO:835]
	*/
	client, err := as.NewClient("127.0.0.1", 3000)
	if err != nil {

		if err.Matches(types.NETWORK_ERROR) || err.Matches(types.INVALID_NODE_ERROR) {
			fmt.Println("Aerospike connectiviy error:: try after sometime ")
			return
		}
	}
	defer client.Close()
	queryPolicy := as.NewQueryPolicy()
	queryPolicy.MaxRecords = 5
	queryPolicy.FilterExpression = as.ExpLet(
		as.ExpDef("bin", as.ExpStringBin("BIN")),
		as.ExpDef("logo", as.ExpStringBin("LOGO")),
		as.ExpDef("code", as.ExpStringBin("BLOCKDOE")),
		as.ExpAnd(
			as.ExpEq(
				as.ExpVar("bin"), as.ExpStringVal("524368")),
			as.ExpEq(as.ExpVar("logo"), as.ExpStringVal("212")),
			as.ExpEq(as.ExpVar("code"), as.ExpStringVal("Q")),
		),
	)
	stmt := as.NewStatement("hdfc", "cctxn")
	result, err := client.Query(queryPolicy, stmt)
	if err != nil {
		log.Fatal(err)
	}
	for recs := range result.Results() {
		if recs != nil {
			fmt.Println(recs.Record.Bins)
		}
	}
	key, err := as.NewKey("hdfc", "cctxn", 45)
	if err != nil {
		log.Fatal(err)
	}
	py := as.NewPolicy()
	rec, _ := client.Get(py, key)
	fmt.Println(rec.Bins)
}
func PrintTestRecords() {
	client, err := as.NewClient("127.0.0.1", 3000)
	fmt.Println("Connected ", client.GetNodeNames())
	defer client.Close()
	if err != nil {
		log.Fatal(err)
	}
	policy := as.NewPolicy()
	key, err := as.NewKey("hdfc", "cctxn", 45)
	if err != nil {
		log.Fatal(err)
	}
	rec, _ := client.Get(policy, key)
	fmt.Println(rec.Bins)

	/*
		select idTxn from hdfc.cctxn where
		logo = X AND bin = Y AND blockcode = Z
	*/
}
func ConnectAndInsert() {
	//only once..
	//keys range 1 to 25000
	client, err := as.NewClient("127.0.0.1", 3000)
	if err != nil {
		log.Fatal(err)
	}
	
	writepolicy := as.NewWritePolicy(0, 0)
	writepolicy.SendKey = true
	data := GenerateRandomList()
	for i, record := range data {
		//insert records here.
		key, err := as.NewKey("hdfc", "cctxn", i+1)
		if err != nil {
			fmt.Println("Error while ceating the key")
		}
		blockCode := as.NewBin("BLOCKDOE", record.blockcode)
		bin := as.NewBin("BIN", record.bin)
		logo := as.NewBin("LOGO", record.logo)
		idApp := as.NewBin("IDAPP", record.idapp)
		idTxn := as.NewBin("IDTXN", record.idtxn)
		err = client.PutBins(writepolicy, key, blockCode, bin, logo, idApp, idTxn)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(i)
	}

	defer client.Close()
	fmt.Println("connected")

}

type CCTxn struct {
	blockcode string
	bin       string
	logo      string
	idapp     string
	idtxn     string
}

func GenerateRandomList() []CCTxn {
	codes := []string{"H", "P", "H", "Q", "J", "Q", "P", "J", "B", "Q", "B", "B", "A", "A", "A", "B", " ",
		"B", "B", "C", "B", "B", "C", "C", "J", "Q", "P", "P", "P", " ", "J", "J"}
	bins := []string{"540536", "405028", "461786", "405028", "461786", "553162", "461786", "405028", "461786",
		"461786", "405028", "461786", "524368", "524368", "540536"}
	idtxns := []string{"PCR", "UNB", "APD", "APR", "VST", "ACI", "CSD", "LDT", "DSR", "DAE"}
	logos := []string{"415", "821", "228", "422", "231", "408", "277", "500", "835", "103", "238", "204", "328",
		"400", "400", "302", "813", "785", "422", "212", "220", "818", "680", "754",
		"902", "302", "601", "601"}
	rand.Seed(time.Now().UnixMilli())
	arr := make([]CCTxn, 0)
	for t := 0; t < 25000; t++ {
		id1 := rand.Intn(len(codes))
		id2 := rand.Intn(len(bins))
		id3 := rand.Intn(len(idtxns))
		id4 := rand.Intn(len(logos))
		c1 := CCTxn{blockcode: codes[id1], bin: bins[id2], idapp: "RS", logo: logos[id4], idtxn: idtxns[id3]}
		arr = append(arr, c1)
	}
	return arr
}
