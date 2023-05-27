package main

/*
Preguntar como inicializar int32 en shorthand
comentar el return con el typing
*/

import (
	"log"
	"strconv"
)

func main() {

	var res string

	for number := 1; number <= 10; number++ {

		res, _ = EvenOrOdd(number)
		log.Println("number " + strconv.FormatInt(int64(number), 10) + " is " + res)
	}

}

func EvenOrOdd(number int) (ever_or_odd string, quotient int) {

	var remainder int

	remainder = number % 2
	quotient = number / 2

	if remainder == 0 {
		ever_or_odd = "Even"
	} else {
		ever_or_odd = "Odd"
	}

	return
}
