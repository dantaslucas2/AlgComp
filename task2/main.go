package main

import "fmt"

type MyMatriz struct {
	matriz    [3][3]int
	m         int
	n         int
	quadrada  bool
	simetrica bool
	positiva  bool
}

func main() {
	fmt.Println("he")

	matriz := [3][3]int{{2, -1, -2}, {-4, 6, 3}, {-4, -2, 8}}

	mymatriz := MyMatriz{}.newMyMatriz(matriz)
	mymatriz.jacobi()
}

func (my MyMatriz) newMyMatriz(matriz [3][3]int) MyMatriz {
	my.matriz = matriz

	my.m = len(my.matriz)
	my.n = len(my.matriz[0])

	if my.n == my.m {
		my.quadrada = true
	} else {
		my.quadrada = false
	}

	if my.quadrada {
		var i = true
		for index_coluna, _ := range my.matriz {
			for index_linha, _ := range my.matriz[0] {
				if my.matriz[index_coluna][index_linha] != my.matriz[index_linha][index_coluna] {
					i = false
					//fmt.Println("false because: ", my.matriz[index_coluna][index_linha], " ", index_coluna, " ", index_linha)
				}
			}
		}
		my.simetrica = i
	}

	return my
}
func (my MyMatriz) jacobi() {
	if my.simetrica {
		var n int = len(my.matriz)

	}
}
