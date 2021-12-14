package main

import (
	"log"
	"golang.org/x/net/context"
	"lab3/lab3"
	"google.golang.org/grpc"
)

func main(){
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000",grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := lab3.NewStarwarsClient(conn)

	response, err := c.Leia(context.Background(),&lab3.L{Planeta:"Tatooine",Ciudad:"Mos_Eisley",Servidor:1})
	if err != nil {
		log.Fatalf("Error when calling Leia: %s", err)
	}
	log.Printf("Response from server: %s",response.Planeta)


}