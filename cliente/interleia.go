package main

import (
	"log"
	"golang.org/x/net/context"
	"lab3/lab3"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":9000",grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	
	c := lab3.NewStarwarsClient(conn)

	// conexion con el server fulcrum 1
	var Ful1 *grpc.ClientConn
	Ful1, err1 := grpc.Dial(":9001",grpc.WithInsecure())
	if err1 != nil {
		log.Fatalf("did not connect: %s", err1)
	}
	defer Ful1.Close()
	c1 := lab3.NewStarwarsClient(Ful1)
	// conexion con el server fulcrum 2
	/* var Ful2 *grpc.ClientConn
	Ful2, err2 := grpc.Dial(":9002",grpc.WithInsecure())
	if err2 != nil {
		log.Fatalf("did not connect: %s", err2)
	}
	defer Ful2.Close()
	c2 := lab3.NewStarwarsClient(Ful2)
	// conexion con el server fulcrum 3
	var Ful3 *grpc.ClientConn
	Ful3, err3 := grpc.Dial(":9003",grpc.WithInsecure())
	if err3 != nil {
		log.Fatalf("did not connect: %s", err3)
	}
	defer Ful3.Close()
	c3 := lab3.NewStarwarsClient(Ful3) */

	planeta := "Tatooine"
	ciudad := "Mos Eisley"
	servidor := int32(1)

	for{
		response, err := c.Interleia(context.Background(),&lab3.L{Planeta:planeta,Ciudad:ciudad,Servidor:servidor})
		if err != nil {
			log.Fatalf("Error when calling Interleia: %s", err)
		}
		log.Printf("Response from server: %s",response.Planeta)

		response2, err2 := c1.Leiafulcrum(context.Background(),&lab3.L{Planeta:response.Planeta,Ciudad:response.Ciudad,Servidor:response.Servidor})
		if err2 != nil {
			log.Fatalf("Error when calling Leiafulcrum: %s", err2)
		}
		log.Printf("Response from server: %s",response2.Planeta)
		planeta = response2.Planeta
		ciudad = response2.Ciudad
		servidor = response2.Servidor
	}
}