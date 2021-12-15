package main

import (
	"log"
	"golang.org/x/net/context"
	"lab3/lab3"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial("10.6.40.253:9001",grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	
	c := lab3.NewStarwarsClient(conn)

	// conexion con el server fulcrum 1
	var Ful1 *grpc.ClientConn
	Ful1, err1 := grpc.Dial("10.6.40.250:9001",grpc.WithInsecure())
	if err1 != nil {
		log.Fatalf("did not connect: %s", err1)
	}
	defer Ful1.Close()
	c1 := lab3.NewStarwarsClient(Ful1)
	// conexion con el server fulcrum 2
	 var Ful2 *grpc.ClientConn
	Ful2, err2 := grpc.Dial("10.6.40.251:9001",grpc.WithInsecure())
	if err2 != nil {
		log.Fatalf("did not connect: %s", err2)
	}
	defer Ful2.Close()
	c2 := lab3.NewStarwarsClient(Ful2)
	// conexion con el server fulcrum 3
	var Ful3 *grpc.ClientConn
	Ful3, err3 := grpc.Dial("10.6.40.252:9001",grpc.WithInsecure())
	if err3 != nil {
		log.Fatalf("did not connect: %s", err3)
	}
	defer Ful3.Close()
	c3 := lab3.NewStarwarsClient(Ful3) 

	planeta := "Tatooine"
	valor := int32(0)
	x := int32(0)
	y := int32(0)
	z := int32(0)
	servidor := int32(1)
	merge := false

	for{
		response, err := c.Interleia(context.Background(),&lab3.Lresponse{Planeta:planeta,Valor:valor,X:x,Y:y,Z:z,Servidor:servidor,Merge:merge})
		if err != nil {
			log.Fatalf("Error when calling Interleia: %s", err)
		}
		log.Printf("Response from server: %s",response.Planeta)

		if response.Servidor == 1{
			response2, err2 := c1.Leiafulcrum(context.Background(),&lab3.L{Planeta:response.Planeta,Ciudad:response.Ciudad,Servidor:response.Servidor})
			if err2 != nil {
				log.Fatalf("Error when calling Leiafulcrum: %s", err2)
			}
			log.Printf("Response from server: %s",response2.Planeta)
			planeta = response2.Planeta
			valor = response2.Valor
			x = response2.X
			y = response2.Y
			z = response2.Z
			servidor = response2.Servidor
			merge = response2.Merge
		}else if response.Servidor == 2{
			response2, err2 := c2.Leiafulcrum(context.Background(),&lab3.L{Planeta:response.Planeta,Ciudad:response.Ciudad,Servidor:response.Servidor})
			if err2 != nil {
				log.Fatalf("Error when calling Leiafulcrum: %s", err2)
			}
			log.Printf("Response from server: %s",response2.Planeta)
			planeta = response2.Planeta
			valor = response2.Valor
			x = response2.X
			y = response2.Y
			z = response2.Z
			servidor = response2.Servidor
			merge = response2.Merge
		}else{
			response2, err2 := c3.Leiafulcrum(context.Background(),&lab3.L{Planeta:response.Planeta,Ciudad:response.Ciudad,Servidor:response.Servidor})
			if err2 != nil {
				log.Fatalf("Error when calling Leiafulcrum: %s", err2)
			}
			log.Printf("Response from server: %s",response2.Planeta)
			planeta = response2.Planeta
			valor = response2.Valor
			servidor = response2.Servidor
		}
	}
}