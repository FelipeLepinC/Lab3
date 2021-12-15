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

type reloj struct{   ///////////////////////////
	x int
	y int
	z int
	servidor int
}

var mp = make(map[string] reloj)

func main(){
	// conexion con el server broker.go
	var Accion string
	var Planeta string
	var Ciudad string
	var Intvalue int32
	var  Svalue  string
	var Servidor int32
	var num string
	var s bool 
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

	var lectura int
	//var comando string 
	reader := bufio.NewReader(os.Stdin)
	s = false
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
			count, ok := mp[Planeta]
			if ok == false {
				Servidor = 0
			}else{
				Servidor = int32(count.servidor)
				fmt.Println(Servidor)
			}
			response, err := c.Alertabroken(context.Background(),&lab3.Operacion{Accion:Accion,Planeta:Planeta,Ciudad:Ciudad,Intvalue:Intvalue,Svalue:Svalue,Servidor:Servidor})
			if err != nil {
				log.Fatalf("Error when calling Alertabroken: %s", err)
			}
			if s == true{
				for key,_ := range mp {
					var r1 reloj
					r1 = mp[key]
					r1.servidor = int(response.Nserver)
					mp[key] = r1
				}
				s = false
			}
			fmt.Println("Respuesta del broker : ",response.Nserver)
			if response.Nserver == 1{
				response1, err := c1.Fulcrum(context.Background(),&lab3.Operacion{Accion:Accion,Planeta:Planeta,Ciudad:Ciudad,Intvalue:Intvalue,Svalue:Svalue,Servidor:response.Nserver})
				if err != nil {
					log.Fatalf("Error when calling Fulcrum1: %s", err)
				}
				log.Printf("Respuesta del fulcrum1 : %s %d %d %d", response1.Planeta, response1.X,response1.Y,response1.Z)
				count, ok := mp[response1.Planeta]
				if ok == false {
					if response1.Planeta != "" {
						fmt.Println("el elemento no estaba", count)
						var r1 reloj
						r1.x = 1
						r1.y = 0
						r1.z = 0
						r1.servidor = int(response.Nserver)
						mp[response1.Planeta]=r1
						fmt.Println(mp)
					}
				} else {
					var r1 reloj
					r1.x = int(response1.X)
					r1.y = int(response1.Y)
					r1.z = int(response1.Z)
					r1.servidor = int(response.Nserver)
					mp[response1.Planeta] = r1
				} 
				if Accion == "DeleteCity" {
					delete(mp, response1.Planeta)
				}
				fmt.Println("")
				if response1.Merge == true{
					for key,_ := range mp {
						var r1 reloj
						r1 = mp[key]
						r1.servidor = 0
						mp[key] = r1
					}
					s = true
				}
				fmt.Println(mp[response1.Planeta])
			}else if response.Nserver == 2{
				response1, err := c2.Fulcrum(context.Background(),&lab3.Operacion{Accion:Accion,Planeta:Planeta,Ciudad:Ciudad,Intvalue:Intvalue,Svalue:Svalue,Servidor:response.Nserver})
				if err != nil {
					log.Fatalf("Error when calling Fulcrum2: %s", err)
				}
				log.Printf(" Respuesta del fulcrum2 : %s %d %d %d", response1.Planeta, response1.X,response1.Y,response1.Z)
				count, ok := mp[response1.Planeta]
				if ok == false {
					if response1.Planeta != "" {
						fmt.Println("el elemento no estaba", count)
						var r1 reloj
						r1.x = 0
						r1.y = 1
						r1.z = 0
						r1.servidor = int(response.Nserver)
						mp[response1.Planeta]=r1
						fmt.Println(mp)
					}
				} else {
					var r1 reloj
					r1.x = int(response1.X)
					r1.y = int(response1.Y)
					r1.z = int(response1.Z)
					r1.servidor = int(response.Nserver)
					mp[response1.Planeta] = r1
				} 
				if Accion == "DeleteCity" {
					delete(mp, response1.Planeta)
				}
				if response1.Merge == true{
					for key,_ := range mp {
						var r1 reloj
						r1 = mp[key]
						r1.servidor = 0
						mp[key] = r1
					}
					s = true
				}
				fmt.Println(mp[response1.Planeta])
			}else if response.Nserver == 3{
				response1, err := c3.Fulcrum(context.Background(),&lab3.Operacion{Accion:Accion,Planeta:Planeta,Ciudad:Ciudad,Intvalue:Intvalue,Svalue:Svalue,Servidor:response.Nserver})
				if err != nil {
					log.Fatalf("Error when calling Fulcrum3: %s", err)
				}
				log.Printf("Respuesta del fulcrum3 : %s %d %d %d", response1.Planeta, response1.X,response1.Y,response1.Z)
				count, ok := mp[response1.Planeta]
				if ok == false {
					if response1.Planeta != "" {
						fmt.Println("el elemento no estaba", count)
						var r1 reloj
						r1.x = 0
						r1.y = 0
						r1.z = 1
						r1.servidor = int(response.Nserver)
						mp[response1.Planeta]=r1
						fmt.Println(mp)
					}
				} else {
					var r1 reloj
					r1.x = int(response1.X)
					r1.y = int(response1.Y)
					r1.z = int(response1.Z)
					r1.servidor = int(response.Nserver)
					mp[response1.Planeta] = r1
				} 
				if Accion == "DeleteCity" {
					delete(mp, response1.Planeta)
				}
				if response1.Merge == true{
					for key,_ := range mp {
						var r1 reloj
						r1 = mp[key]
						r1.servidor = 0
						mp[key] = r1
					}
					s = true
				}
				fmt.Println(mp[response1.Planeta])
			}else{
				log.Fatalf("Otro caso")
			}
		}	
	}
}