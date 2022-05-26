package main

import (
	"fmt"
)

func main() {
	palavra := [3]int{1, 2, 3}
	fmt.Println(palavra)
	tes(palavra)

	fmt.Println("variavel original ", palavra)

}

func tes(pala [3]int) {
	var palavra2 = pala
	palavra2[0] = 88
	fmt.Println("variavel criada ", palavra2)
}
