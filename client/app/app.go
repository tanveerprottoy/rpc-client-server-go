package app

import (
	"fmt"
	"log"
	"net/rpc"
)

func StartRPCClient() {
	client, err := rpc.DialHTTP("tcp", ":3000")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply string
	err = client.Call("RPC.Execute", "clientCall", &reply)
	if err != nil {
		log.Fatal("RPC error:", err)
	}
	fmt.Print("RPC:", reply)
}
