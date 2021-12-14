package main

import (
	"log"
	"golang.org/x/net/context"
	"lab3/lab3"
	"google.golang.org/grpc"
)

type Reloj struct{   ///////////////////////////
	x int
	y int
	z int
	servidor int
}

var mapaDos = make(map[string] Reloj)

func main(){
	var Planeta string
	var Ciudad string
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
					Servidor = 0
				}else{
					Servidor = int32(count.servidor)
					fmt.Println(Servidor)
				}
				response, err := c.Leia(context.Background(),&lab3.L{Planeta:planeta,Ciudad:ciudad,Servidor:Servidor})
				if err != nil {
					log.Fatalf("Error when calling Leia: %s", err)
				}
				log.Printf("Response from server: %s",response.Planeta)
				

			}
		}
	}

}