package tools

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
