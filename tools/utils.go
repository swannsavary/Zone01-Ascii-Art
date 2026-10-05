package tools

// split un string en fonction d'une rune en []string
func SplitByRune(str string, r rune) []string {
	var result []string
	var mot []rune
	for _, cara := range str {
		if cara == r {
			result = append(result, string(mot))
			mot = []rune{}
		} else {
			mot = append(mot, cara)
		}
	}
	return result
}

func hasPrefixe(mot string, prefix string) bool {
	if prefix > mot {
		return false
	}
	for index, c := range prefix {
		if c != rune(mot[index]) {
			return false
		}
	}
	return true
}
