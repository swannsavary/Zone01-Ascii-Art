package tools

import (
	"os"
)

// separe les caracteres d'un string et le separe en un slice de slice de rune a chaque \n
func Splitcara(s string) [][]rune {
	var result [][]rune
	ligne := []rune{}
	s2 := []rune(s)
	for i := 0; i < len(s2); i++ {
		if i != len(s2)-1 && s2[i] == '\\' && s2[i+1] == 'n' {
			result = append(result, ligne)
			ligne = []rune{}
			i++
		} else {
			ligne = append(ligne, s2[i])
		}
	}
	if len(ligne) != 0 {
		result = append(result, ligne)
	}
	return result
}

// affiche les caracteres du fichier en mode art
func PrintCara(file []string, cara [][]rune, color string, caraAColorer string) {
	var result string
	var fin []int
	var debut []int
	for j := 0; j < len(cara); j++ {
		if len(cara[j]) == 0 {
			result += "\n"
			continue
		}
		for i := 0; i < 9; i++ {
			for e := 0; e < len(cara[j]); e++ {
				poslettre := (cara[j][e] - ' ') * 9
				if caraAColorer == "" {
					result += colors[color] + file[i+int(poslettre)] + colors["reset"]
					continue
				}

				// À déterminer : est-ce que le caractère e
				// appartient à une occurrence de caraAColorer ?
				if e+len([]rune(caraAColorer)) < len(cara[j]) && string(cara[j][e:e+len([]rune(caraAColorer))]) == caraAColorer {
					debut = append(debut, e)
					fin = append(fin, e+len([]rune(caraAColorer)))
					result += colors[color] + file[i+int(poslettre)]
				} else {
					if e > debut[len(debut)-1] && e < fin[len(fin)-1] {
						result += file[i+int(poslettre)]
					} else {
						result += file[i+int(poslettre)] + colors["reset"]
					}

				}
			}

			result += "\n"
		}
	}

	os.Stdout.WriteString(result)
}
