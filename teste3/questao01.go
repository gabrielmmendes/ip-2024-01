package main

import (
	"fmt"
	"strings"
)

func main() {
	var frase string

	for {
		var x string
		fmt.Scan(&x)
		frase += " " + x

		if strings.Contains(x, "#") {
			break
		}
	}

	frase = strings.Replace(frase, "#", "", len(frase))
	frase = strings.TrimSpace(frase)

	var novaFrase string
	for i := len(frase) - 1; i >= 0; i-- {
		var x string
		x = fmt.Sprintf("%c", frase[i])
		novaFrase += x
	}

	fmt.Println(novaFrase)
}
