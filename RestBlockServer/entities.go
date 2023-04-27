package main

type Txn struct {
	Bin  string `json:"bin"`
	Logo string `json:"logo"`
	Code string `json:"code"`
}
type IdTxn struct {
	Bin  string `json:"bin"`
	Logo string `json:"logo"`
	Code string `json:"code"`
	Txn  string `json:"txn"`
}
