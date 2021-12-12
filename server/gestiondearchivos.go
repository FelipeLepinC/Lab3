/*TODO:
- Si hace addCity, verificar que la ciudad no exista										DONE
- Si hace updateName, verificar que el planeta exista										DONE
- Si hace updateNumber, verificar que el planeta exista									DONE
- Si hace deleteCity, verificar que el planeta exista										DONE
- AddCity este puede contener o no el nuevo-valor (Ciudad o soldados?)
- (Opcional) Si hace updateName, verificar que la ciudad exista					(Si no hay match, pasa de largo)
- (Opcional) Si hace updateNumber, verificar que la ciudad exista				(Si no hay match, pasa de largo)
- (Opcional) Si hace deleteCity, verificar que la ciudad exista					(Si no hay match, pasa de largo)
*/

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//-----------------------------------------------------------------------------
//Gestión de planetas y ciudades
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

//-----------------------------------------------------------------------------
//Gestión de registro de servidores fulcrum
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

//Función principal
func main() {
	//AddCity("Tatooine", "Valparaíso", "35")
	//ReadPlanet("Tatooine")
	//UpdateName("Tatooine", "Valparaíso", "Talcahuano")
	//UpdateNumber("Tatooine", "Valparaíso", "50")
	//DeleteCity("Tatooine", "Talcahuano")
	//AddRegistro("Fulcrum-1", "AddCity Tatooine Valparaíso 35")
	//ClearRegistro("Fulcrum-1")
}
