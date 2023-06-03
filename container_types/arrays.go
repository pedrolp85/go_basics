/*
Arrays
Todos los valores contenidos son del mismo tipo
cada elemento de la lista ocupa un indice como en Python 0:lenght-1

Se almacenan en partes contiguas de memoria (en caso de los arrays la
parte directa) los slices y maps tiene una parte indirecta ?

La complejidad de acceso a cualquier valor mediante su clave es O(1)
los accesos en array y slice son más rápidos que en los maps

Por defecto, todos los elementos del array son inicializados con el
equivalente al 0 del tipo de los valores
int sería 0, string ""

Representacion: array types: [N]T
Donde T es un tipo arbitrario
y N la longitud del array

La longitud es parte del tipo del array, y como en Golang los tipos no pueden
cambiarse, un array no puede redimensionarse ni igualar 2 arrays de diferente longitud

Muy importante, los arrays se pasan por valor y no por referencia, por lo que cuando le
pasas un array a una función, se crea una copia

*/

package main

import (
	"fmt"
)

func main() {

	const Size = 32

	type Person struct {
		name string
		age  int
	}

	var x [5]int // An array of 5 integers
	fmt.Println(x)

	var y [8]string // An array of 8 strings
	fmt.Println(y)

	// Element type is a struct type: Person
	var array_personalized_type [5]Person
	fmt.Println(array_personalized_type)

	// Acceso a valores por índice

	x[0] = 100
	x[1] = 101
	x[3] = 103
	x[4] = 105

	fmt.Printf("x[0] = %d, x[1] = %d, x[2] = %d\n", x[0], x[1], x[2])
	fmt.Println("x = ", x)

	// Declaring and initializing an array at the same time
	var a = [5]int{2, 4, 6, 8, 10}
	// The expression on the right-hand side of the above statement is called an array literal.
	fmt.Println("a=", a)

	// Short hand declaration
	b := [5]int{2, 4, 6, 8, 10}
	fmt.Println("b=", b)

	// Letting Go compiler infer the length of the array
	c := [...]int{3, 5, 7, 9, 11, 13, 17}
	fmt.Println("c= ", c)

	//iterating an array with a for loop and index integers
	for i := 0; i < len(c); i++ {
		fmt.Println("iteracion ", c[i])
	}

	// interating an array with range
	daysOfWeek := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	// Esto es como el enumerate de Python
	for index, value := range daysOfWeek {
		fmt.Printf("Day %d of week = %s\n", index, value)
	}
}

// The followings array composite literals
// are equivalent to each other.

/*
asigned_array := [4]bool{false, true, true, false}
[4]bool{0: false, 1: true, 2: true, 3: false}
[4]bool{1: true, true}
[4]bool{2: true, 1: true}
[...]bool{false, true, true, false}
[...]bool{3: false, 1: true, true}
*/
