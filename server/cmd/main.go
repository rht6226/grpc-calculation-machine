package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/rht6226/grpc-machine/machine"
	"github.com/rht6226/grpc-machine/server"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 9111, "Port on which gRPC server should listen TCP conn.")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	machine.RegisterMachineServer(grpcServer, &server.MachineServer{})
	log.Printf("Initializing gRPC server on port %d", *port)
	grpcServer.Serve(lis)
}
