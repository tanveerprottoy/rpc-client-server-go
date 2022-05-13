package app

import (
	"errors"
	"net"
	"net/rpc"
)

type Response struct {
	Message string
}

type Request struct {
	Name string
}

type RPC struct{}

func (h *RPC) Execute(req Request, res *Response) (err error) {
	if req.Name == "" {
		err = errors.New("A name must be specified")
		return
	}
	res.Message = "Hello " + req.Name
	return
}

func StartRPCServer() {
	rpc.Register(&RPC{})

	// Create a TCP listener that will listen on `Port`
	listener, _ := net.Listen("tcp", ":3000")

	// Close the listener whenever we stop
	defer listener.Close()

	// Wait for incoming connections
	rpc.Accept(listener)
}
