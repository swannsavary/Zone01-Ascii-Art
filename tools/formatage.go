package tools

import (
	"errors"
	"os"
)

// map qui lie le nom d'un fichier au fichier .txt correspondant
var banners = map[string]string{
	"standard":   "standard.txt",
	"shadow":     "shadow.txt",
	"thinkertoy": "thinkertoy.txt",
}

// prend les arguments donner et renvoie chaque option et le bon fichier a exploiter
func Formatage(args []string) ([]byte, string, string, string, string, error) {
	args2 := args
	var output, color, motacolorer, align string
	file, inArg, err := option_file(args2)
	if err != nil {
		return file, "", "", "", "", err
	}
	if inArg {
		args2 = args2[:len(args2)-1]
	}
	for index, option := range args2[:len(args2)-1] {
		if hasPrefixe(option, "--output=") {
			output = option[9:]
		} else if hasPrefixe(option, "--color=") {
			color, motacolorer = option_color(args2, index)
		} else if hasPrefixe(option, "--align=") {
			align = option[8:]
		}
	}
	return file, color, motacolorer, output, align, nil
}

// renvoie le bon fichier a exploiter
func option_file(args []string) ([]byte, bool, error) {
	nomDuFichier := "standard.txt"
	if len(args) == 0 {
		return nil, false, errors.New("need 1 or more arguments")
	}
	if len(args) > 1 {
		if chemin, exist := banners[args[len(args)-1]]; exist {
			nomDuFichier = chemin
			file, err := os.ReadFile(nomDuFichier)
			return file, true, err
		}
	}
	file, err := os.ReadFile(nomDuFichier)
	return file, false, err
}

// renvoie la bonne couleur a utiliser et le mot a colorer si il y en a
func option_color(args []string, index int) (string, string) {
	color := ""
	motAColorer := ""
	if len(args) > 1 {
		if hasPrefixe(args[index], "--color=") {
			color = args[index][8:]
		}
		if len(args)-1 > index+1 && !hasPrefixe(args[index+1], "--") {
			motAColorer = args[index+1]
		}
	}
	return color, motAColorer
}
