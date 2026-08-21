package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//args := os.Args[1]
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	file2 := strings.Split(string(file), "\n")
	cara := []rune{'a', 'b', 'c'}
	for i := 0; i <= 8; i++ {
		for e := 0; e < len(cara); e++ {
			poslettre := (cara[e]-' ')*8 + cara[e] - ' '
			fmt.Print(file2[i+int(poslettre)])
		}
		fmt.Println()
	}
}
