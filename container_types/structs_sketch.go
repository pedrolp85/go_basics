package main

import "log"

// Declaracion de tipos de datos.

type SerVivoStruct struct {
	Persona struct {
		Nombre string `json:"name"`
	}
	Animal struct {
		Nombre string `json:"name"`
	}
	Edad int32 `json:"years"`
}

func NewSerVivo() SerVivoStruct {
	return SerVivoStruct{}
}

func (s SerVivoStruct) Latir() {
	if s.Animal.Nombre != "" {
		log.Println("late el corazón de animal.")
		/*... informacion propia del animal ...*/
	}

	if s.Persona.Nombre != "" {
		log.Println("late el corazón de persona.")
		/*... informacion propia de la persona ...*/
	}
}

type PersonaStruct struct {
	Nombre string `json:"name"`
	//Servivo SerVivoStruct `json:"ser_vivo"`
}

type AnimalStruct struct {
	Nombre string `json:"name"`
	//Servivo SerVivoStruct `json:"ser_vivo"`
}

// Declararcion de variables

// funcion init

// funcion main
func main() {

	// Verificacion de que una estructura esté inicializada.
	var personaNulo PersonaStruct

	log.Println(personaNulo)
	log.Println()

	persona := PersonaStruct{
		Persona: Persona{
			Nombre: "dnombre personas",
		},
	}
	persona.Nombre = "pepe"
	persona.Servivo = SerVivoStruct{Edad: 10}
	persona.Servivo.Persona.Nombre

	if personaNulo == persona {
		log.Println("no inicializado")
	} else {
		log.Println(persona)
	}

	// Verificacion de que una cadena se encuentre incializada.
	nombre := "pepe"
	if "" != nombre {
		log.Println("Nombre:", nombre)
	} else {
		log.Println("No inicializado.")
	}

}
