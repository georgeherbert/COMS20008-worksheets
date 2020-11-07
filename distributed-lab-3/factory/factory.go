package main

import (
	"flag"
	"fmt"
	"net"
	"pairbroker/stubs"
	"net/rpc"
	"strconv"
)

type Factory struct {}

//TODO: Define a Multiply function to be accessed via RPC.
//Check the previous weeks' examples to figure out how to do this.
func (f *Factory) Multiply(req stubs.Pair, res *stubs.JobReport) (err error) {
	res.Result = req.X * req.Y
	fmt.Println(strconv.Itoa(req.X) + " * " + strconv.Itoa(req.Y) + " = " + strconv.Itoa(res.Result))
	return
}

func (f *Factory) Divide(req stubs.Pair, res *stubs.JobReport) (err error) {
	res.Result = req.X / req.Y
	fmt.Println(strconv.Itoa(req.X) + " / " + strconv.Itoa(req.Y) + " = " + strconv.Itoa(res.Result))
	return
}

func main(){
	pAddr := flag.String("ip", "127.0.0.1:8050", "IP and port to listen on")
	brokerAddr := flag.String("broker","127.0.0.1:8030", "Address of broker instance")
	flag.Parse()
	//TODO: You'll need to set up the RPC server, and subscribe to the running broker instance.

	client, _ := rpc.Dial("tcp", *brokerAddr)
	defer client.Close()

	subscribeMul := stubs.Subscription{Topic: "multiply", FactoryAddress: *pAddr, Callback: "Factory.Multiply"}
	statusMul := new(stubs.StatusReport)
	client.Go(stubs.Subscribe, subscribeMul, statusMul, nil)

	subscribeDiv := stubs.Subscription{Topic: "divide", FactoryAddress: *pAddr, Callback: "Factory.Divide"}
	statusDiv := new(stubs.StatusReport)
	client.Go(stubs.Subscribe, subscribeDiv, statusDiv, nil)

	rpc.Register(&Factory{})
	listener, _ := net.Listen("tcp", *pAddr)
	defer listener.Close()
	rpc.Accept(listener)
}
