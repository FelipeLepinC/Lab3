package main

import (
	"log"
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
	// conexion con el server fulcrum 2
	var Ful2 *grpc.ClientConn
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
	c3 := lab3.NewStarwarsClient(Ful3)

	//response, err := c.Enviarinfo(context.Background(),&lab3.Info{Planeta: "Tatooine",Ciudad: "Mos_Eisley", Soldados: 5})
	Relojes = make(map[string][]*reloj)
	response, err := c.Alertabroken(context.Background(),&lab3.Operacion{Accion:"AddCity",Planeta:"Tatooine",Ciudad:"Mos Eisley",Valor:5})
	if err != nil {
		log.Fatalf("Error when calling Enviarinfo: %s", err)
	}

	reloj1 := reloj{Planeta:"Tatooine",x:0,y:0,z:0}
	reloj2 := reloj{Planeta:"Tierra",x:1,y:0,z:0}
	store(reloj1)
	store(reloj2)

	if response.Nserver == 1{
		response1, err := c1.Fulcrum(context.Background(),&lab3.Operacion{Accion:"AddCity",Planeta:"Tatooine",Ciudad:"Mos Eisley",Valor:5})
		if err != nil {
			log.Fatalf("Error when calling Enviarinfo: %s", err)
		}
		log.Fatalf("Response : %s %d %d %d", response1.Planeta, response1.X,response1.Y,response1.Z)
	}else if response.Nserver == 2{
		response1, err := c2.Fulcrum(context.Background(),&lab3.Operacion{Accion:"AddCity",Planeta:"Tatooine",Ciudad:"Mos Eisley",Valor:5})
		if err != nil {
			log.Fatalf("Error when calling Enviarinfo: %s", err)
		}
		log.Fatalf("Response : %s %d %d %d", response1.Planeta, response1.X,response1.Y,response1.Z)
	}else if response.Nserver == 3{
		response1, err := c3.Fulcrum(context.Background(),&lab3.Operacion{Accion:"AddCity",Planeta:"Tatooine",Ciudad:"Mos Eisley",Valor:5})
		if err != nil {
			log.Fatalf("Error when calling Enviarinfo: %s", err)
		}
		log.Fatalf("Response : %s %d %d %d", response1.Planeta, response1.X,response1.Y,response1.Z)
	}else{
		log.Fatalf("Otro caso")
	}
}