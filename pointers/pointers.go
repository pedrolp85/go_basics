package main

/*
GO soporta punteros, lo que nos permite pasar referencias a variables o registros
Para tipar firmas de funciones que aceptan punteros de entrada: func example (iptr *int)

*/

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
