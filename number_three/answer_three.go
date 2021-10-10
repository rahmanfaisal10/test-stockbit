package number_three

import "strings"

func FindFirstStringInBracket(str string) string {
	indexFirstBracketFound := strings.Index(str, "(")
	if indexFirstBracketFound >= 0 {
		indexClosingBracketFound := strings.Index(str, ")")
		switch {
		case indexFirstBracketFound > indexClosingBracketFound:
			return ""
		case indexClosingBracketFound >= 0:
			return str[indexFirstBracketFound+1 : indexClosingBracketFound]
		default:
			return ""
		}
	}

	return ""
}
