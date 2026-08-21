package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("need 1 argument")
		return
	}
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	caractere := splitcara(args[0])
	fmt.Println(caractere)
	file2 := strings.Split(string(file), "\n")
	for i := 0; i <= 8; i++ {
		for e := 0; e < len(caractere); e++ {
			poslettre := (caractere[e]-' ')*8 + caractere[e] - ' '
			fmt.Print(file2[i+int(poslettre)])
		}
		fmt.Println()
	}
}

func splitcara(s string) []rune {
	var result []rune
	s2 := []rune(s)
	for i := 0; i < len(s2); i++ {
		if s2[i] == '\\' && s2[i+1] == 'n' {
			result = append(result, '\n')
		} else {
			result = append(result, s2[i])
		}
	}
	return result
}
