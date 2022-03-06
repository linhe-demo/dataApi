package tools

import "strings"

func InsertStringSpecialCharacter(target string, symbol string) (out string) {
	tmpLen := len(target)
	for i := 0; i < tmpLen; i += 3 {
		out += target[i:i+3] + symbol
	}
	out = strings.TrimRight(out, symbol)
	return out
}
