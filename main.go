package main

import (
	"ascii-art/tools"
	"os"
)

func main() {
	args := os.Args[1:]
	posOfWrite := 0
	file, color, motAColorer, output, err := tools.Formatage(args)
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		return
	}
	if color != "" {
		posOfWrite++
	}
	if motAColorer != "" {
		posOfWrite++
	}
	if output != "" {
		posOfWrite++
	}
	caractere := tools.Splitcara(args[posOfWrite])
	file2 := tools.SplitByRune(string(file), '\n')
	tools.PrintCara(file2, caractere, color, motAColorer, output)
}
