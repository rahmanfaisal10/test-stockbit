package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findFirstStringInBracket1("ometa)(gText"))
}

func findFirstStringInBracket1(str string) string {
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
