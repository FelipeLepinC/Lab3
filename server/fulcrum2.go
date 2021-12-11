package main

import (
	"fmt"
	"log"
	"net"
	"lab3/lab3"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Bienvenidos al server GRPC!!!")
	lis,err := net.Listen("tcp", ":9002")
	if err != nil {
		log.Fatal(err)
	}
	l := lab3.Server{}

	grpcServer := grpc.NewServer()

	lab3.RegisterStarwarsServer(grpcServer,&l)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}