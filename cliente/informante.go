package main

import (
	"log"
	"fmt"
	"golang.org/x/net/context"
	"lab3/lab3"
	"google.golang.org/grpc"
)

type reloj struct {
	Planeta string
	x int32
	y int32
	z int32
}

var Relojes map[string][]*reloj

func store(Reloj reloj) {
    Relojes[Reloj.Planeta] = append(Relojes[Reloj.Planeta], &Reloj)
}

func main(){
	// conexion con el server broker.go
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

	//response, err := c.Enviarinfo(context.Background(),&lab3.Info{Planeta: "Tatooine",Ciudad: "Mos_Eisley", Soldados: 5})
	Relojes = make(map[string][]*reloj)
	response, err := c.Alertabroken(context.Background(),&lab3.Operacion{Accion:"AddCity",Planeta:"Tatooine",Ciudad:"Mos Eisley",Valor:5})

	reloj1 := reloj{Planeta:"Tatooine",x:0,y:0,z:0}
	reloj2 := reloj{Planeta:"Tierra",x:1,y:0,z:0}
	store(reloj1)
	store(reloj2)

	if err != nil {
		log.Fatalf("Error when calling Enviarinfo: %s", err)
	}
	if response.Nserver == 1{
		Ful, err2 := grpc.Dial(":9001",grpc.WithInsecure())
		if err2 != nil {
			log.Fatalf("did not connect: %s", err2)
		}
		defer Ful.Close()
		
		c := lab3.NewStarwarsClient(Ful)
	}else if response.Nserver == 2{
		Ful, err2 := grpc.Dial(":9002",grpc.WithInsecure())
		if err2 != nil {
			log.Fatalf("did not connect: %s", err2)
		}
		defer Ful.Close()
		
		c := lab3.NewStarwarsClient(Ful)
	}else if response.Nserver == 3{
		Ful, err2 := grpc.Dial(":9003",grpc.WithInsecure())
		if err2 != nil {
			log.Fatalf("did not connect: %s", err2)
		}
		defer Ful.Close()
		
		c := lab3.NewStarwarsClient(Ful)
	}else{
		log.Printf("Response from server: %d",response.Nserver)
	}
}