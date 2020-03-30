package main

import (
	"fmt"
	"strings"
)

// UnpackString etc
func UnpackString(s string) string {

	var sb strings.Builder
	var prev rune
	var special bool

	if len(s) < 2 {
		return s // nothing to unpack, return same
	}

	for _, c := range s {

		r := rune(c)
		esc := prev == '\\'

		switch {

		case r == '\\':
			if esc {
				sb.WriteRune(r)
				special = true
			} else {
				special = false
			}

		case '0' <= r && r <= '9':

			if !esc || special {
				if '0' <= prev && prev <= '9' && !special {
					return "Incorrect format!"
				}
				for i := 1; i < int(r-'0'); i++ {
					sb.WriteRune(prev)
				}
				special = false
			} else {
				sb.WriteRune(r)
				special = true
			}

		default:
			// any other symbol
			sb.WriteRune(r)
		}
		prev = c
	}

	return sb.String()
}

func main() {
	fmt.Println(UnpackString("a4bc2d5e"))
	fmt.Println(UnpackString("abcd"))
	fmt.Println(UnpackString("45"))
	fmt.Println(UnpackString(`qwe\4\5`))
	fmt.Println(UnpackString(`qwe\45`))
	fmt.Println(UnpackString(`qwe\\5`))
}
