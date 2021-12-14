package lab3

import (
	"log"
	"golang.org/x/net/context"
	"math/rand"
	"time"
	"os"
	"strings"
	"bufio"
	"fmt"
	"io/ioutil"
	"strconv"
)

type Server struct {
	process bool
	planeta string
	ciudad string
	soldados int32
	intvalue int32
    servidor int32
}

type reloj struct{   ///////////////////////////
	x int
	y int
	z int
	servidor int
}

type rj struct{   ///////////////////////////
	x int32
	y int32
	z int32
}

var mapaDos = make(map[string] reloj)

func ReadPlanet(planeta string) {
	fileName := planeta + ".txt"

	//Revisamos si el archivo del planeta existe
	if _, err := os.Stat(fileName); err == nil {
		//Si existe, lo abrimos
		data, err := ioutil.ReadFile(fileName)

		//En caso de error, lo printea
		if err != nil {
			log.Panicf("failed reading data from file: %s\n", err)
		}

		//Imprimimos el contenido del archivo
		//fmt.Printf("\nFile Name: %s", fileName)
		//fmt.Printf("\nSize: %d bytes", len(data))
		fmt.Printf("Leyendo contenido de planeta %s\n", planeta)
		fmt.Printf("%s", data)
	} else {
		fmt.Printf("Planeta %s no existe\n", planeta)
	}
}

func AddCity(planeta string, ciudad string, soldados string) {
	fileName := planeta + ".txt"
	var existe = false

	//Revisamos si el archivo del planeta existe
	if _, err := os.Stat(fileName); err == nil {
		//Si existe, lo abrimos
		file, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		//Leemos línea por línea
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			split := strings.Split(scanner.Text(), " ")
			if split[1] == ciudad {
				//Si existe la ciudad, lo marcamos
				existe = true
			}
		}

		//Si existe la ciudad, no se puede agregar
		if existe {
			fmt.Printf("Ciudad %s ya existe\n", ciudad)
			return
		}
	}

	/*Si el archivo del planeta no existe, lo creamos y escribimos
	Si el archivo del planeta existe, pero la ciudad no estaba, sólo escribimos*/
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	//Agregamos la ciudad
	file.WriteString(planeta + " " + ciudad + " " + soldados + "\n")
	fmt.Printf("Registro [%s, %s, %s] creado\n", planeta, ciudad, soldados)
}

func ChangeInfo(tipo string, planeta string, ciudad string, dato_new string) {
	oldFileName := planeta + ".txt"
	tempFileName := planeta + "_temp.txt"

	//Revisamos si el archivo del planeta existe
	if _, err := os.Stat(oldFileName); err == nil {
		//Si existe, lo abrimos
		oldFile, err := os.Open(oldFileName)
		if err != nil {
			log.Fatal(err)
		}
		defer oldFile.Close()

		//Creamos un archivo temporal
		tempFile, err := os.Create(tempFileName)
		if err != nil {
			log.Fatal(err)
		}
		defer tempFile.Close()

		//Buscamos la ciudad linea por linea
		var existe = false
		scanner := bufio.NewScanner(oldFile)
		for scanner.Scan() {
			split := strings.Split(scanner.Text(), " ")

			if tipo == "DeleteCity" {
				//Si es una operación de tipo DeleteCity, escribimos sólo las otras ciudades
				if split[1] != ciudad {
					tempFile.WriteString(scanner.Text() + "\n")
				} else {
					existe = true
					fmt.Printf("Ciudad %s eliminada\n", ciudad)
				}
			} else if tipo == "UpdateName" {
				//Si es una operación de tipo UpdateName, actualizamos el nombre y escribimos todas las ciudades
				if split[1] == ciudad {
					tempFile.WriteString(split[0] + " " + dato_new + " " + split[2] + "\n")
					existe = true
					fmt.Printf("Nombre de ciudad %s actualizado a %s\n", ciudad, dato_new)
				} else {
					tempFile.WriteString(scanner.Text() + "\n")
				}
			} else if tipo == "UpdateNumber" {
				//Si es una operación de tipo UpdateNumber, actualizamos el valor y escribimos todas las ciudades
				if split[1] == ciudad {
					tempFile.WriteString(split[0] + " " + split[1] + " " + dato_new + "\n")
					existe = true
					fmt.Printf("Cantidad de soldados de ciudad %s actualizada a %s\n", ciudad, dato_new)
				} else {
					tempFile.WriteString(scanner.Text() + "\n")
				}
			}
		}

		//Eliminamos el archivo obsoleto
		os.Remove(oldFileName)

		//Renombramos el archivo temporal
		e := os.Rename(tempFileName, oldFileName)
		if e != nil {
			log.Fatal(e)
		}

		if !existe {
			fmt.Printf("Ciudad %s no existe\n", ciudad)
		}

	} else {
		fmt.Printf("Planeta %s no existe\n", oldFileName)
	}
}

func DeleteCity(planeta string, ciudad string) {
	ChangeInfo("DeleteCity", planeta, ciudad, "")
}

func UpdateName(planeta string, ciudad_old string, ciudad_new string) {
	ChangeInfo("UpdateName", planeta, ciudad_old, ciudad_new)
}

func UpdateNumber(planeta string, ciudad string, soldados_new string) {
	ChangeInfo("UpdateNumber", planeta, ciudad, soldados_new)
}

func AddRegistro(servidor string, operacion string) {
	fileName := servidor + ".txt"

	//Si el archivo no existe, lo crea y escribe. Si existe, escribe
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	//Agregamos la ciudad
	file.WriteString(operacion + "\n")
}

func ClearRegistro(servidor string) {
	fileName := servidor + ".txt"

	//Revisamos si el archivo del servidor existe
	if _, err := os.Stat(fileName); err == nil {
		//Si existe, lo borramos y lo creamos de nuevo
		os.Remove(fileName)
		os.Create(fileName)

		fmt.Printf("Registro del servidor %s limpiado\n", servidor)
	} else {
		fmt.Printf("Registro del servidor no existe\n")
	}
}

func GetPlanet(planeta string) string {
	filename := planeta + ".txt"
	contenido := ""

	//Revisamos si el archivo del planeta existe
	if _, err := os.Stat(filename); err == nil {
		//Si existe, lo abrimos
		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		//Vemos línea por línea
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			contenido += scanner.Text() + "\n"
		}
		return contenido

	} else {
		fmt.Printf("Planeta %s no existe\n", planeta)
		return ""
	}
}

func (s *Server) Enviarinfo(ctx context.Context, in *Info) (*Info, error){
	return &Info{Planeta: "Tatooine",Ciudad: "Mos_Eisley", Soldados: 5},nil
}

func (s *Server) Fulcrum(ctx context.Context, in *Operacion) (*Reloj, error){
	log.Printf("%s %s %s", in.Accion, in.Planeta,in.Ciudad)
	if in.Accion == "AddCity" {
		if in.Intvalue != -1 {
			fmt.Println(in.Intvalue)
			AddCity(in.Planeta, in.Ciudad, strconv.Itoa(int(in.Intvalue)))
		}else{
			AddCity(in.Planeta, in.Ciudad, "")
		}
		count, ok := mapaDos[in.Planeta]
		if ok == false {
			fmt.Println("el elemento no estaba", count)
			var r1 reloj
			if in.Servidor == 1{
				r1.x = 1
				r1.y = 0
				r1.z = 0
			}else if in.Servidor == 2{
				r1.x = 0
				r1.y = 1
				r1.z = 0
			}else{
				r1.x = 0
				r1.y = 0
				r1.z = 1
			}
			r1.servidor = int(in.Servidor)
			mapaDos[in.Planeta] = r1
			fmt.Println(mapaDos)
		} else {
			var r1 reloj
			r1 = mapaDos[in.Planeta] 
			if in.Servidor == 1{
				r1.x = r1.x + 1
			}else if in.Servidor == 2{
				r1.y = r1.y + 1
			}else{
				r1.z = r1.z + 1
			}
			mapaDos[in.Planeta] = r1
		} 
	}else if in.Accion == "DeleteCity"{
		DeleteCity(in.Planeta, in.Ciudad)
		delete(mapaDos, in.Planeta)
	}else if in.Accion == "UpdateName"{
		UpdateName(in.Planeta, in.Ciudad, in.Svalue)
		count, ok := mapaDos[in.Planeta]
		if ok == true {
			var r1 reloj
			r1 = mapaDos[in.Planeta]
			if in.Servidor == 1{
				r1.x = r1.x + 1
			}else if in.Servidor == 2{
				r1.y = r1.y + 1
			}else{
				r1.z = r1.z + 1
			}
			mapaDos[in.Planeta] = r1
		}else{
			fmt.Println("el elemento no estaba", count)
			return &Reloj{Planeta: "",X:int32(0),Y:int32(0),Z:int32(0)},nil
		}
	}else if in.Accion == "UpdateNumber"{
		UpdateNumber(in.Planeta, in.Ciudad, strconv.Itoa(int(in.Intvalue)))
		count, ok := mapaDos[in.Planeta]
		if ok == true {
			var r1 reloj
			r1 = mapaDos[in.Planeta]
			if in.Servidor == 1{
				r1.x = r1.x + 1
			}else if in.Servidor == 2{
				r1.y = r1.y + 1
			}else{
				r1.z = r1.z + 1
			}
			mapaDos[in.Planeta] = r1
		}else{
			fmt.Println("el elemento no estaba", count)
			return &Reloj{Planeta: "",X:int32(0),Y:int32(0),Z:int32(0)},nil
		}
	}else{
		log.Printf("Caso contrario")	
	}
	var r reloj
	r = mapaDos[in.Planeta]
	return &Reloj{Planeta: in.Planeta,X:int32(r.x),Y:int32(r.y),Z:int32(r.z)},nil
}	

func (s *Server) Alertabroken(ctx context.Context, in *Operacion) (*Message, error){
	n := int32(0)
	if in.Accion != "" && in.Servidor == 0 {
		rand.Seed(time.Now().UnixNano())
		n = int32(1 + rand.Intn(3-1+1))
		log.Printf("Servidor elegido es : %d",n)
	}else{
		n = in.Servidor	
	}
	return &Message{Nserver:n},nil
}

func (s *Server) Intermediario(ctx context.Context, in *Message) (*Message, error) {
	log.Printf(" %d ",in.Nserver)
	return &Message{Nserver:in.Nserver},nil
}

func (s *Server) Leia(ctx context.Context, in *L) (*L, error) {
	log.Printf(" %s ",in.Planeta)
	s.planeta = in.Planeta
	s.ciudad = in.Ciudad
	s.servidor = in.Servidor
	s.process = true
	log.Printf("%t",s.process)
	for(s.process) {
		//log.Printf("Esta aqui for leia")
	}
	return &L{Planeta: s.planeta,Ciudad:s.ciudad,Servidor:s.servidor}, nil
}

func (s *Server) Interleia(ctx context.Context, in *L) (*L, error) {
	log.Printf(" %s ",in.Planeta)
	s.planeta = in.Planeta
	s.ciudad = in.Ciudad
	s.servidor = in.Servidor
	s.process = false
	log.Printf("%t",s.process)
	for(!s.process){
		//log.Printf("Esta aqui for Interleia")
	}
	log.Printf("Sale del ciclo")
	return &L{Planeta: s.planeta,Ciudad:s.ciudad,Servidor:s.servidor}, nil
}

func (s *Server) Leiafulcrum(ctx context.Context, in *L)(*L, error){
	log.Printf(" %s ",in.Planeta)
	return &L{Planeta: in.Planeta,Ciudad:in.Ciudad,Servidor:in.Servidor}, nil
}

func (s *Server) Pedirdic(ctx context.Context, in *Message) (*Merge, error) {
	// pasar los relojes que tenemos a formato string -> string -> string
	var r1 reloj
	r1.x = 1
	r1.y = 0
	r1.z = 0
	r1.servidor = 1
	mapaDos["Tierra"]=r1
	var m = make(map[string]string)
	for key, value := range mapaDos {
		fmt.Println("Key:", key, "Value:", value.x)
		m[key] = strconv.Itoa(value.x)+","+strconv.Itoa(value.y)+","+strconv.Itoa(value.z)
		// Agregar funcion josue GetPlanet(planeta string)
	}
	return &Merge{M:m},nil 
}

func (s *Server) Enviardic(ctx context.Context, in *Merge) (*Message, error) {
	for key, _ := range mapaDos {
		delete(mapaDos, key)
	}
	n := 0
	var r1 reloj
	for key, value := range in.M {
		fmt.Println("Key:", key, "Value:", value)
		delimitador := ","
		l := strings.Split(value, delimitador)
		n, _ = strconv.Atoi(l[0])
		r1.x = n
		n, _ = strconv.Atoi(l[1])
		r1.y = n
		n, _ = strconv.Atoi(l[2])
		r1.z = n
		r1.servidor = 1
		mapaDos[key]=r1
		// Falta pasar el string a archivo 
	}
	return &Message{Nserver:1},nil
}