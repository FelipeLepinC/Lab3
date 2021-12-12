package main

import (
	"log"
	"golang.org/x/net/context"
	"lab3/lab3"
	"google.golang.org/grpc"
	"time"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9001",grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := lab3.NewStarwarsClient(conn)

	var Ful2 *grpc.ClientConn
	Ful2, err2 := grpc.Dial(":9002",grpc.WithInsecure())
	if err2 != nil {
		log.Fatalf("did not connect: %s", err2)
	}
	defer Ful2.Close()
	c2 := lab3.NewStarwarsClient(Ful2)

	var Ful3 *grpc.ClientConn
	Ful3, err3 := grpc.Dial(":9003",grpc.WithInsecure())
	if err3 != nil {
		log.Fatalf("did not connect: %s", err3)
	}
	defer Ful3.Close()
	c3 := lab3.NewStarwarsClient(Ful3)

	for {
        time.Sleep(5 * time.Second)
		response, err := c2.Intermediario(context.Background(),&lab3.Message{Nserver:2})
		if err != nil {
			log.Fatalf("Error when calling Enviarinfo: %s", err)
		}
		log.Printf("Response from server: %d",response.Nserver)

		response1, e1 := c3.Intermediario(context.Background(),&lab3.Message{Nserver:3})
		if e1 != nil {
			log.Fatalf("Error when calling Enviarinfo: %s", e1)
		}
		log.Printf("Response from server: %d",response1.Nserver)

		response2, e2 := c.Intermediario(context.Background(),&lab3.Message{Nserver:1})
		if e2 != nil {
			log.Fatalf("Error when calling Enviarinfo: %s", e2)
		}
		log.Printf("Response from server: %d",response2.Nserver)
    }
}