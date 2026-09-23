package tools

import (
	"errors"
	"os"
)

// structure qui lie le nom d'un fichier au fichier .txt correspondant
var banners = map[string]string{
	"standard":   "standard.txt",
	"shadow":     "shadow.txt",
	"thinkertoy": "thinkertoy.txt",
}

func Formatage(args []string) ([]byte, string, string, error) {
	args2 := args
	file, inArg, err := optionfile(args2)
	if err != nil {
		return file, "", "", err
	}
	if inArg {
		args2 = args[:len(args)-1]
	}
	color, motacolorer := optioncolor(args2)
	return file, color, motacolorer, nil
}

func optionfile(args []string) ([]byte, bool, error) {
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

func optioncolor(args []string) (string, string) {
	color := ""
	motAColorer := ""
	if len(args) > 1 {
		if args[0][:8] == "--color=" {
			color = args[0][8:]
		}
		if len(args) == 3 {
			motAColorer = args[1]
		}
	}
	return color, motAColorer
}
