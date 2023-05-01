package main

import (
	"flag"
	"log"

	"github.com/rht6226/grpc-machine/client"
	"github.com/rht6226/grpc-machine/machine"
	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("server_addr", "localhost:9111", "The server address in the format of host:port")
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	mClient := machine.NewMachineClient(conn)

	instructionset := machine.InstructionSet{Instructions: []*machine.Instruction{
		{Operand: 5, Operator: "PUSH"},
		{Operand: 6, Operator: "PUSH"},
		{Operator: "MUL"},
		{Operand: 10, Operator: "PUSH"},
		{Operator: "DIV"},
		{Operand: 14, Operator: "PUSH"},
		{Operator: "ADD"},
		{Operator: "FIB"},
	}}

	client.RunExecute(mClient, &instructionset)
}
