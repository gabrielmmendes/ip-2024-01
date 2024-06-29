package main

import (
	"fmt"
)

func ordenaCrescente(v []int) {
	for i := 0; i < len(v); i++ {
		for j := 0; j < len(v) - i - 1; j++ {
			if v[j] > v[j + 1] {
				x := v[j]
				v[j] = v[j + 1]
				v[j + 1] = x
			}
		}
	}
}

func main() {
	v := []int{22, 63, 10, 2, 5, 92}

	ordenaCrescente(v)

	fmt.Println(v)
}
