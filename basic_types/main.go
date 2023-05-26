/*
FICHERO RAIZ
*/

package main

import (
	"log"
	_ "strconv"
)

// Main de la aplicacion.
func main() {

	/*
		uint8 // (tinyint),
		uint32 // (integer),
		uint64 // (bigint)
		int8,
		int32,
		int64,


			Size	Range
		float32	32 bits	-3.4e+38 to 3.4e+38.
		float64	64 bits	-1.7e+308 to +1.7e+308.

	*/

	//nombreInferido := "Gonzalo"
	a := 2
	b := 2
	_, err := dividir(a, b)

	if err != nil {
		// Error
		log.Println("Error:" + err.Error())
	}

	log.Println("Finalizamos.")

}

// Division entre dos numeros.
func dividir(a int, b int) (int, error) {
	defer func() {
		log.Println("Cerramos conexion BBDD.")
		log.Println("Cerramos Streams con Files.")
	}() //log.Println("Hemos terminado la division.")
	log.Println("Hola.")
	return a / b, nil
}
