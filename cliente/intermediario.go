package main

import (
	"log"
	"golang.org/x/net/context"
	"lab3/lab3"
	"google.golang.org/grpc"
	"time"
	"fmt"
	"strings"
	"strconv"
)

func Encontrar(d1 []string, d2 []string, d3 []string,f int ) int{
	maximo := -99999999
	Maxfila := 0
	if f == 3 {
		for column := 1; column <= 3; column++ {
			for i := 0; i < 3; i++ {
				if column == 1 {
					n, _ := strconv.Atoi(d1[i])
					if n >= maximo {
						maximo = n
						Maxfila = 1
					}
				}else if column == 2 {
					n, _ := strconv.Atoi(d2[i])
					if n >= maximo{
						maximo =  n
						Maxfila = 2
					}
				}else{
					n, _ := strconv.Atoi(d3[i])
					if n >= maximo{
						maximo = n
						Maxfila = 3
					}
				}
			}
		}
	}else if f == 2{
		for column := 1; column <= 2; column++ {
			for i := 0; i < 3; i++ {
				if column == 1 {
					n, _ := strconv.Atoi(d1[i])
					if n >= maximo {
						maximo = n
						Maxfila = 1
					}
				}else{
					n, _ := strconv.Atoi(d2[i])
					if n >= maximo{
						maximo =  n
						Maxfila = 2
					}
				}
			}
		}
	}
	return Maxfila
}

func main() {
	var largos [3]int
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("10.6.40.250:9001",grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := lab3.NewStarwarsClient(conn)

	var Ful2 *grpc.ClientConn
	Ful2, err2 := grpc.Dial("10.6.40.251:9001",grpc.WithInsecure())
	if err2 != nil {
		log.Fatalf("did not connect: %s", err2)
	}
	defer Ful2.Close()
	c2 := lab3.NewStarwarsClient(Ful2) 

	var Ful3 *grpc.ClientConn
	Ful3, err3 := grpc.Dial("10.6.40.252:9001",grpc.WithInsecure())
	if err3 != nil {
		log.Fatalf("did not connect: %s", err3)
	}
	defer Ful3.Close()
	c3 := lab3.NewStarwarsClient(Ful3)

	var nuevo = make(map[string]string)
	var temp = make(map[string]string)
	var temp1 = make(map[string]string)
	var temp2 = make(map[string]string)

	for {
        time.Sleep(120 * time.Second)
		response, err := c.Pedirdic(context.Background(),&lab3.Message{Nserver:1})
		if err != nil {
			log.Fatalf("Error when calling Pedirdic: %s", err)
		}
		response1, e1 := c2.Pedirdic(context.Background(),&lab3.Message{Nserver:2})
		if e1 != nil {
			log.Fatalf("Error when calling Pedirdic: %s", e1)
		}

		response2, e2 := c3.Pedirdic(context.Background(),&lab3.Message{Nserver:3})
		if e2 != nil {
			log.Fatalf("Error when calling Pedirdic: %s", e2)
		}
		// Elegir el mejor 
		largos[0] = len(response.M) 
		largos[1] = len(response1.M) 
		largos[2] = len(response2.M)

		max := -999999999
		mindex := -1
		i := 0

		for _,num := range largos {
			if num > max {
				max = num 
				mindex = i
			}
			i = i + 1
		}
		fmt.Println(mindex)

		if mindex == 0{
			temp = response.M
			temp1 = response1.M
			temp2 = response2.M
		}else if mindex == 1{
			temp = response1.M
			temp1 = response.M
			temp2 = response2.M
		}else{
			temp = response2.M
			temp1 = response.M
			temp2 = response1.M
		}

		n := 0
		for key, value := range temp {
			fmt.Println("Key:", key, "Value:", value)
			delimitador := ","
			r1 := strings.Split(value, delimitador)
			count1, ok1 := temp1[key] 
			count2, ok2 := temp2[key]
			if ok1 && ok2 { // Los 3 servers tienen la info 
				r2 := strings.Split(value, delimitador)
				r3 := strings.Split(value, delimitador)
				n = Encontrar(r1, r2, r3,3)
			}else if ok1{ // El server 1 y el server3 tienen el planeta
				r3 := strings.Split(value, delimitador)
				n = Encontrar(r1, r3, r1 ,2)
				if n == 2 {
					n = 3
				}
			}else if ok2{ // El server 1 y el server2  tienen el planeta 
				r2 := strings.Split(value, delimitador)
				n = Encontrar(r1, r2, r1 ,2)
			}else{// solo el server 1 tiene el planeta 
				n = 1
			}
			if n == 1{
				nuevo[key] = value
			}else if n == 2{
				nuevo[key] = count1
			}else{
				nuevo[key] = count2
			}
		}
		fmt.Println("Diccionario nuevo")

		for key, value := range nuevo{
			fmt.Printf(key,value)
		}

		resp1, er1 := c.Enviardic(context.Background(),&lab3.Merge{M:nuevo})
		if er1 != nil {
			log.Fatalf("Error when calling Enviardic: %s", er1)
		}
		resp2, er2 := c2.Enviardic(context.Background(),&lab3.Merge{M:nuevo})
		if er2 != nil {
			log.Fatalf("Error when calling Enviardic: %s", er2)
		}
		resp3, er3 := c3.Enviardic(context.Background(),&lab3.Merge{M:nuevo})
		if er3 != nil {
			log.Fatalf("Error when calling Enviardic: %s", er3)
		}
		fmt.Println(resp1.Nserver,resp2.Nserver,resp3.Nserver)
    }
}