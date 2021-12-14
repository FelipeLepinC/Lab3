package main

import (
	"log"
	"golang.org/x/net/context"
	"lab3/lab3"
	"google.golang.org/grpc"
	"fmt"
	"strings"
	"strconv"
	"bufio"
	"os"
)

type Reloj struct{   ///////////////////////////
	x int
	y int
	z int
	servidor int
}

var mapaDos = make(map[string] Reloj)

func main(){
	// conexion con el server broker.go
	var Accion string
	var Planeta string
	var Ciudad string
	var Intvalue int32
	var  Svalue  string
	var Servidor int32
	var num string
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

	var lectura int
	//var comando string 
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("¿Desea ejecutar un comando?")
		fmt.Println("1.-Si")
		fmt.Println("2.-No")
		fmt.Scanln(&lectura)
		if lectura == 2 {
			break
		} else if lectura== 1 {
			fmt.Println("ingrese el comando que desea ejecutar")
			comando, _ := reader.ReadString('\n')
			//fmt.Scanln(&comando)
			//fmt.Println(comando)
			delimitador := " "
			nombresComoArreglo := strings.Split(comando, delimitador)
			fmt.Println(nombresComoArreglo)
			if nombresComoArreglo[0] == "AddCity" {
				Accion = nombresComoArreglo[0]
				Planeta = nombresComoArreglo[1]
				if len(nombresComoArreglo)== 4{
					Ciudad = nombresComoArreglo[2]
					num = strings.Replace(nombresComoArreglo[3], "\n", "", -1)
					r1, _ := strconv.Atoi(num)
					Intvalue = int32(r1)
					Svalue = ""
				}else if len(nombresComoArreglo) == 3{
					Ciudad = strings.Replace(nombresComoArreglo[2], "\n", "", -1)
					Intvalue = int32(-1)
					Svalue = ""
				}
				fmt.Println(Accion,Planeta,Ciudad,Intvalue,Svalue)
			} else if nombresComoArreglo[0] == "DeleteCity"{
				Accion = nombresComoArreglo[0] 
				Planeta = nombresComoArreglo[1]
				Ciudad = strings.Replace(nombresComoArreglo[2], "\n", "", -1)
				Intvalue = 0
				Svalue = ""
			} else if nombresComoArreglo[0] == "UpdateName"{
			    Accion = nombresComoArreglo[0] 
				Planeta = nombresComoArreglo[1]
				Ciudad = nombresComoArreglo[2]
				Svalue = strings.Replace(nombresComoArreglo[3], "\n", "", -1)
				Intvalue = 0
			} else if nombresComoArreglo[0] == "UpdateNumber"{
				Accion = nombresComoArreglo[0] 
				Planeta = nombresComoArreglo[1]
				Ciudad = nombresComoArreglo[2]
				num = strings.Replace(nombresComoArreglo[3], "\n", "", -1)
				r1, _ := strconv.Atoi(num)
				Intvalue = int32(r1)
				Svalue = ""
			} 
			count, ok := mapaDos[Planeta]
			if ok == false {
				Servidor = 0
			}else{
				Servidor = int32(count.servidor)
				fmt.Println(Servidor)
			}
			response, err := c.Alertabroken(context.Background(),&lab3.Operacion{Accion:Accion,Planeta:Planeta,Ciudad:Ciudad,Intvalue:Intvalue,Svalue:Svalue,Servidor:Servidor})
			if err != nil {
				log.Fatalf("Error when calling Enviarinfo: %s", err)
			}
			fmt.Println("Respuesta del broker : ",response.Nserver)
			if response.Nserver == 1{
				response1, err := c1.Fulcrum(context.Background(),&lab3.Operacion{Accion:Accion,Planeta:Planeta,Ciudad:Ciudad,Intvalue:Intvalue,Svalue:Svalue,Servidor:response.Nserver})
				if err != nil {
					log.Fatalf("Error when calling Enviarinfo: %s", err)
				}
				log.Printf("Respuesta del fulcrum1 : %s %d %d %d", response1.Planeta, response1.X,response1.Y,response1.Z)
				count, ok := mapaDos[response1.Planeta]
				if ok == false {
					if response1.Planeta != "" {
						fmt.Println("el elemento no estaba", count)
						var r1 Reloj
						r1.x = 1
						r1.y = 0
						r1.z = 0
						r1.servidor = int(response.Nserver)
						mapaDos[response1.Planeta]=r1
						fmt.Println(mapaDos)
					}
				} else {
					var r1 Reloj
					r1 = mapaDos[response1.Planeta]
					r1.x = r1.x + 1
					r1.servidor = int(response.Nserver)
					mapaDos[response1.Planeta] = r1
				} 
				if Accion == "DeleteCity" {
					delete(mapaDos, response1.Planeta)
				}
				if response1.Merge == true{
					for key,value := range mapaDos {
						var r1 Reloj
						r1 = mapaDos[response1.Planeta]
						r1.servidor = 0
						mapaDos[response1.Planeta] = r1
					}
				}
				fmt.Println(mapaDos[response1.Planeta])
			}else if response.Nserver == 2{
				response1, err := c2.Fulcrum(context.Background(),&lab3.Operacion{Accion:Accion,Planeta:Planeta,Ciudad:Ciudad,Intvalue:Intvalue,Svalue:Svalue,Servidor:response.Nserver})
				if err != nil {
					log.Fatalf("Error when calling Enviarinfo: %s", err)
				}
				log.Printf(" Respuesta del fulcrum2 : %s %d %d %d", response1.Planeta, response1.X,response1.Y,response1.Z)
				count, ok := mapaDos[response1.Planeta]
				if ok == false {
					if response1.Planeta != "" {
						fmt.Println("el elemento no estaba", count)
						var r1 Reloj
						r1.x = 0
						r1.y = 1
						r1.z = 0
						r1.servidor = int(response.Nserver)
						mapaDos[response1.Planeta]=r1
						fmt.Println(mapaDos)
					}
				} else {
					var r1 Reloj
					r1 = mapaDos[response1.Planeta] 
					r1.y = r1.y + 1
					r1.servidor = int(response.Nserver)
					mapaDos[response1.Planeta] = r1
				} 
				if Accion == "DeleteCity" {
					delete(mapaDos, response1.Planeta)
				}
				if response1.Merge == true{
					for key,value := range mapaDos {
						var r1 Reloj
						r1 = mapaDos[response1.Planeta]
						r1.servidor = 0
						mapaDos[response1.Planeta] = r1
					}
				}
				fmt.Println(mapaDos[response1.Planeta])
			}else if response.Nserver == 3{
				response1, err := c3.Fulcrum(context.Background(),&lab3.Operacion{Accion:Accion,Planeta:Planeta,Ciudad:Ciudad,Intvalue:Intvalue,Svalue:Svalue,Servidor:response.Nserver})
				if err != nil {
					log.Fatalf("Error when calling Enviarinfo: %s", err)
				}
				log.Printf("Respuesta del fulcrum3 : %s %d %d %d", response1.Planeta, response1.X,response1.Y,response1.Z)
				count, ok := mapaDos[response1.Planeta]
				if ok == false {
					if response1.Planeta != "" {
						fmt.Println("el elemento no estaba", count)
						var r1 Reloj
						r1.x = 0
						r1.y = 0
						r1.z = 1
						r1.servidor = int(response.Nserver)
						mapaDos[response1.Planeta]=r1
						fmt.Println(mapaDos)
					}
				} else {
					var r1 Reloj
					r1 = mapaDos[response1.Planeta]
					r1.z = r1.z + 1
					r1.servidor = int(response.Nserver)
					mapaDos[response1.Planeta] = r1
				} 
				if Accion == "DeleteCity" {
					delete(mapaDos, response1.Planeta)
				}
				if response1.Merge == true{
					for key,value := range mapaDos {
						var r1 Reloj
						r1 = mapaDos[response1.Planeta]
						r1.servidor = 0
						mapaDos[response1.Planeta] = r1
					}
				}
				fmt.Println(mapaDos[response1.Planeta])
			}else{
				log.Fatalf("Otro caso")
			}
		}	
	}
}