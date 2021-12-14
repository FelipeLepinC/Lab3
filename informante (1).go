package main

import (
	"log"
	"golang.org/x/net/context"
	"lab3/lab3"
	"google.golang.org/grpc"
	"fmt"
	"strconv"
	"strings"

)

type Reloj struct{   ///////////////////////////
	x int
	y int
	z int
	servidor int
}

var mapaDos = make(map[string] Reloj)

func main(){

	var Accion string
	var Planeta string
	var Ciudad string
	var Intvalue int32
	var  Svalue  string
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
	/* conexion con el server fulcrum 2
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
	*/
	var lectura int
	var comando string 
    for {
		fmt.Println("¿Desea ejecutar un comando?")
		fmt.Println("1.-Si")
		fmt.Println("2.-No")
		fmt.Scanln(&lectura)
		if lectura == 2 {
			break
		} else if lectura== 1 {		

			response, err := c.Alertabroken(context.Background(),&lab3.Operacion{Accion:"AddCity",Planeta:"Tatooine",Ciudad:"Mos Eisley",Intvalue:5})
			if err != nil {
				log.Fatalf("Error when calling Enviarinfo: %s", err)
			}
			fmt.Println("ingrese el comando que desea ejecutar")
			fmt.Scanln(&comando)
			fmt.Println(comando)
			delimitador := " "
			nombresComoArreglo := strings.Split(comando, delimitador)
			fmt.Println(nombresComoArreglo)
			if nombresComoArreglo[0] == "AddCity" {
				Accion = nombresComoArreglo[0]
				Planeta = nombresComoArreglo[1]
				Ciudad = nombresComoArreglo[2]
				r1, _ := strconv.Atoi(nombresComoArreglo[3])
				Intvalue = int32(r1)
			} else if nombresComoArreglo[0] == "DeleteCity"{
				Accion = nombresComoArreglo[0] 
				Planeta = nombresComoArreglo[1]
				Ciudad = nombresComoArreglo[2]

			} else if nombresComoArreglo[0] == "UpdateName"{
			    Accion = nombresComoArreglo[0] 
				Planeta = nombresComoArreglo[1]
				Ciudad = nombresComoArreglo[2]
				Svalue = nombresComoArreglo[3]
			} else if nombresComoArreglo[0] == "UpdateNumber"{
				Accion = nombresComoArreglo[0] 
				Planeta = nombresComoArreglo[1]
				Ciudad = nombresComoArreglo[2]
				Svalue = nombresComoArreglo[3]
			} 
			if response.Nserver == 1{
				response1, err := c1.Fulcrum(context.Background(),&lab3.Operacion{Accion:Accion,Planeta:Planeta,Ciudad:Ciudad,Intvalue:Intvalue, Svalue: Svalue, Servidor : 1})
				if err != nil {
					log.Fatalf("Error when calling Enviarinfo: %s", err)
				}
				fmt.Println("Response : %s %d %d %d", response1.Planeta, response1.X,response1.Y,response1.Z)
				count, ok := mapaDos[response1.Planeta]
				if ok == false {
					fmt.Println("el elemento no estaba", count)
					var r1 Reloj
					r1.x = 1
					r1.y = 0
					r1.z = 0
					r1.servidor = 1
					mapaDos[response1.Planeta]=r1
					fmt.Println(mapaDos)
				} else {
					var r1 Reloj
					r1 = mapaDos[response1.Planeta] 
					r1.x = r1.x + 1
					mapaDos[response1.Planeta] = r1
					fmt.Println(mapaDos)
				} 
			}else if response.Nserver == 2{
				response1, err := c1.Fulcrum(context.Background(),&lab3.Operacion{Accion:Accion,Planeta:Planeta,Ciudad:Ciudad,Intvalue:Intvalue, Svalue: Svalue, Servidor : 2})
				if err != nil {
					log.Fatalf("Error when calling Enviarinfo: %s", err)
				}
				log.Fatalf("Response : %s %d %d %d", response1.Planeta, response1.X,response1.Y,response1.Z)
				count, ok := mapaDos[response1.Planeta]
				if ok == false {
					fmt.Println("el elemento no estaba", count)
					var r1 Reloj
					r1.x = 1
					r1.y = 0
					r1.z = 0
					r1.servidor = 1
					mapaDos[response1.Planeta]=r1
					fmt.Println(mapaDos)
				} else {
					var r1 Reloj
					mapaDos[response1.Planeta] = r1
					r1.x = r1.x + 1
					mapaDos[response1.Planeta] = r1
				} 
			}else if response.Nserver == 3{
				response1, err := c1.Fulcrum(context.Background(),&lab3.Operacion{Accion:Accion,Planeta:Planeta,Ciudad:Ciudad,Intvalue:Intvalue, Svalue: Svalue, Servidor : 3})
				if err != nil {
					log.Fatalf("Error when calling Enviarinfo: %s", err)
				}
				log.Fatalf("Response : %s %d %d %d", response1.Planeta, response1.X,response1.Y,response1.Z)
				count, ok := mapaDos[response1.Planeta]
				if ok == false {
					fmt.Println("el elemento no estaba", count)
					var r1 Reloj
					r1.x = 1
					r1.y = 0
					r1.z = 0
					r1.servidor = 1
					mapaDos[response1.Planeta]=r1
					fmt.Println(mapaDos)
				} else {
					var r1 Reloj
					mapaDos[response1.Planeta] = r1
					r1.x = r1.x + 1
					mapaDos[response1.Planeta] = r1
				} 
			}else{
				log.Fatalf("Otro caso")
			}
		}
	}
}