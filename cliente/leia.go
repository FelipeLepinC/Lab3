package main

import (
	"log"
	"golang.org/x/net/context"
	"lab3/lab3"
	"google.golang.org/grpc"
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main(){
	var planeta string
	var ciudad string
	var lectura int
	var servidor int32 
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000",grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := lab3.NewStarwarsClient(conn)
	reader := bufio.NewReader(os.Stdin)

	for{
		fmt.Println("¿Desea ejecutar un comando?")
		fmt.Println("1.-Si")
		fmt.Println("2.-No")
		fmt.Scanln(&lectura)
		if lectura == 2 {
			break
		} else if lectura== 1 {
			fmt.Println("ingrese el comando que desea ejecutar")
			comando, _ := reader.ReadString('\n')
			delimitador := " "
			nombresComoArreglo := strings.Split(comando, delimitador)
			fmt.Println(nombresComoArreglo)
			if nombresComoArreglo[0] == "GetNumberRebelds" {
				planeta = nombresComoArreglo[1]
				ciudad = nombresComoArreglo[2]
				count, ok := mapaDos[planeta]
				if ok == false {
					servidor = 0
				}else{
					servidor = int32(count.servidor)
					fmt.Println(servidor)
				}
				response, err := c.Leia(context.Background(),&lab3.L{Planeta:planeta,Ciudad:ciudad,Servidor:servidor})
				if err != nil {
					log.Fatalf("Error when calling Leia: %s", err)
				}
				log.Printf("Response from server: %s",response.Planeta)
				count, ok = mapaDos[response.Planeta]
				if ok == false {
					if response.Planeta != "" {
						fmt.Println("el elemento no estaba", count)
						var r1 Reloj
						r1.x = 0
						r1.y = 1
						r1.z = 0
						r1.servidor = int(response.Servidor)
						mapaDos[response.Planeta] = r1
						fmt.Println(mapaDos)
					}
				} else {
					var r1 Reloj
					r1 = mapaDos[response.Planeta] 
					r1.y = r1.y + 1
					r1.servidor = int(response.Servidor)
					mapaDos[response.Planeta] = r1
				} 
			}
		}
	}

}