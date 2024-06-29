package main

import (
	"fmt"
)

func troca(v []int, x int, y int) {
	z := v[x]
	v[x] = v[y]
	v[y] = z
}

func trocaOpostosSeMenor(v []int, n int) {
	for i := 0; i < n; i++ {
		if i > n - i {
			break
		}

		if v[i] < v[n - 1] {
			troca(v, i, n - 1)
		}
	}
}

func main() {
	v := []int{22, 63, 10, 2, 5, 92}

	trocaOpostosSeMenor(v, 6)

	fmt.Println(v)
}
