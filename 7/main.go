package main

import "fmt"

func main() {
	salarios := map[int]string{1: "alvaro", 2: "carlos", 3: "marlon"}

	fmt.Printf(salarios[21])
	for id, nome := range salarios {
		fmt.Printf("O nome de %d é %s", id, nome)
	}

	// blank identifier (white value)
	
	for _, nome := range salarios {
		fmt.Printf("O nome é %s", nome)
	}
}
