package main

import (
	"fmt"

	aerospike "github.com/aerospike/aerospike-client-go/v8"
)

const (
	namespace = "test"
	set       = "test"
	host      = "localhost" //127.0.0.1
	port      = 3000
)

func main() {
	client, err := aerospike.NewClient(host, port)
	if err != nil {
		fmt.Println("Cannot connect to Aerospike")
	}
	defer client.Close()

	key, err := aerospike.NewKey(namespace, set, 1)
	if err != nil {
		fmt.Println("Key is not created")
	}
	key1, err := aerospike.NewKey(namespace, set, 2)
	if err != nil {
		fmt.Println("Key is not created")
	}
	insert := aerospike.BinMap{
		"name": "Sachin",
		"Page": 25,
	}
	insert1 := aerospike.BinMap{
		"name": "Sahil",
		"Page": 26,
	}
	client.Put(nil, key, insert)
	client.Put(nil, key1, insert1)
	rec, err := client.ScanAll(nil, namespace, set)
	if err != nil {
		fmt.Println("Cannot scan the record")
	}
	for record := range rec.Results() {
		fmt.Println(record.Record.Bins)
	}
	rec.Close()
}
