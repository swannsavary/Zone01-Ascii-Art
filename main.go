package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	file, err := whoFile(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	caractere := splitcara(args[0])
	file2 := strings.Split(string(file), "\n")
	printCara(file2, caractere)
}

// separe les caractere d'un string et le separe en un slice de slice de rune a chaque \n
func splitcara(s string) [][]rune {
	var result [][]rune
	s2 := []rune(s)
	s3 := []rune{}
	for i := 0; i < len(s2); i++ {
		if s2[i] == '\\' && s2[i+1] == 'n' {
			if len(s3) != 0 {
				result = append(result, s3)
			}
			s3 = []rune{}
		} else if i != 0 && s2[i-1] == '\\' && s2[i] == 'n' {
			continue
		} else {
			s3 = append(s3, s2[i])
		}
	}
	if len(s3) != 0 {
		result = append(result, s3)
	}
	return result
}

// affiche les caracteres du fichier
func printCara(file []string, cara [][]rune) {
	for j := 0; j < len(cara); j++ {
		for i := 0; i <= 8; i++ {
			for e := 0; e < len(cara[j]); e++ {
				poslettre := (cara[j][e] - ' ') * 9
				fmt.Print(file[i+int(poslettre)])
			}
			if i != 0 {
				fmt.Println()
			}
		}
	}
}

// structure qui lie le numero optionnel au fichier correspondant
var banners = map[string]string{
	"1": "standard.txt",
	"2": "shadow.txt",
	"3": "thinkertoy.txt",
}

// renvoie le bonne fichier a exploiter en fonction du numero optionnel
func whoFile(args []string) ([]byte, error) {
	nomDuFichier := "standard.txt"
	if len(args) == 2 {
		if chemin, exist := banners[args[1]]; exist {
			nomDuFichier = chemin
		}
	} else if len(args) != 1 {
		return nil, errors.New("need 1 or 2 argument")

	}
	return os.ReadFile(nomDuFichier)
}
