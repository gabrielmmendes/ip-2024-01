package main

import "fmt"

func main() {
	var num int

	fmt.Scan(&num)

	if num%3 == 0 && num%5 == 0 {
		fmt.Printf("O NUMERO E DIVISIVEL \n")
	} else {
		fmt.Printf("O NUMERO NAO E DIVISIVEL \n")
	}
}
