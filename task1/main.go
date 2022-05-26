package main

import (
	"fmt"
	"math"
)

const n int = 3

func main() {
	var resp string
	fmt.Println("Informe qual o método de resolução da matriz")
	fmt.Scanln(resp)
	//matriz1 := [][]int{{1, 4, 9}, {4, 4, 8}, {9, 8, 7}}
	//matriz := [3][3]int{{4, 12, -16}, {12, 37, -43}, {-16, -43, 98}}
	matriz := [3][3]int{{2, -1, -2}, {-4, 6, 3}, {-4, -2, 8}}

	mymatriz := MyMatriz{}.newMyMatriz(matriz)
	mymatriz.printarMatriz()
	//mymatriz.decompsicao_cholesky(matriz)
	mymatriz.decompsicao_lu(matriz)
	mymatriz.printarMatriz()

}

type MyMatriz struct {
	matriz    [3][3]int
	m         int
	n         int
	quadrada  bool
	simetrica bool
	positiva  bool
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
	my.positividade()

	var det int = my.laplace(my.matriz, 0)

	fmt.Println("determinante: ", det)
	fmt.Println("quadrada: ", my.quadrada)
	fmt.Println("simetrica: ", my.simetrica)
	fmt.Println("positiva: ", my.positiva)
	return my
}
func (my MyMatriz) cofactor(mat [3][3]int, temp [3][3]int, p int, q int, n int) {
	var i, j int
	i = 0
	j = 0
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if row != p && col != q {
				temp[i][j+1] = mat[row][col]
				if j == n-1 {
					j = 0
					i++
				}
			}
		}
	}
}
func (my MyMatriz) laplace(matriz [3][3]int, n int) int {
	var det int = 0
	if my.quadrada {
		if my.n == 1 {
			det = my.matriz[0][0]
		} else {
			var temp [3][3]int
			var sign int = 1

			for index_linhas := 0; index_linhas < n; index_linhas++ {
				my.cofactor(matriz, temp, 0, index_linhas, n)
				det += sign * matriz[index_linhas][0] * my.laplace(temp, n-1)

				sign = -sign
			}
		}
		return det
	} else {
		return -1
	}
}

func (my MyMatriz) positividade() {
	var i = true
	fmt.Println("****************")
	for index_coluna, _ := range my.matriz {
		for index_linha, _ := range my.matriz[0] {
			if index_coluna == index_coluna && my.matriz[index_coluna][index_linha] < 0 {
				i = false
				//fmt.Println("false because: ", my.matriz[index_coluna][index_linha], " ", index_coluna, " ", index_linha)
			}
		}
	}
	my.positiva = i
}

func (my MyMatriz) printarMatriz() {
	fmt.Println("\t Matriz")
	for index_coluna, _ := range my.matriz {
		for index_linha, _ := range my.matriz[0] {
			fmt.Print(my.matriz[index_coluna][index_linha], "\t")
		}
		fmt.Println()
	}
}

func (my MyMatriz) decompsicao_cholesky(matriz [3][3]int) {
	if my.quadrada && my.simetrica {
		var n int = len(matriz)
		lower := matriz
		for i := 0; i < n; i++ {
			for j := 0; j <= i; j++ {
				var sum int = 0
				if j == i {
					for k := 0; k < j; k++ {
						sum += int(math.Pow(float64(lower[j][k]), 2))
					}
					lower[j][j] = int(math.Sqrt(float64(matriz[j][j] - sum)))
				} else {
					for k := 0; k < j; k++ {
						sum += (lower[i][k] * lower[j][k])
					}
					lower[i][j] = ((matriz[i][j] - sum) / lower[j][j])
					lower[j][i] = 0
				}
			}
		}

		my.printarMatriz()

		fmt.Println(" Lower Triangular \t transposta ")
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				fmt.Print(lower[i][j], "\t")
			}
			fmt.Print(" ")
			for j := 0; j < n; j++ {
				fmt.Print(lower[j][i], "\t")
			}
			fmt.Println()
		}
	} else {
		fmt.Println("ERROR A MATRIZ NÃO É QUADRADA OU SIMETRICA")
	}
}
func (my MyMatriz) decompsicao_lu(matriz [3][3]int) {
	if my.quadrada {
		l := my.matriz
		u := my.matriz
		var n int = len(my.matriz)
		for i := 0; i < n; i++ {
			for k := i; k < n; k++ {
				var sum int = 0
				for j := 0; j < i; j++ {
					sum += (l[i][j] * u[j][k])
				}
				u[i][k] = my.matriz[i][k] - sum
				if k > i {
					u[k][i] = 0
				}
			}
			for k := i; k < n; k++ {
				if i == k {
					l[i][i] = 1
				} else {
					var sum int = 0
					for j := 0; j < i; j++ {
						sum += (l[k][j] * u[j][i])
					}
					l[k][i] = (my.matriz[k][i] - sum) / u[i][i]
					l[i][k] = 0
				}
			}
		}
		fmt.Println("     Lower Triangular / Upper Triangular")

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				fmt.Print("   ", l[i][j], "\t")
			}
			fmt.Print("\t")
			for j := 0; j < n; j++ {
				fmt.Print("    ", u[i][j], "\t")
			}
			fmt.Print("\n")
		}
	}
}
