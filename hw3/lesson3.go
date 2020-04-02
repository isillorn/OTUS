package main

/*
Распаковка строки
Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:

* "a4bc2d5e" => "aaaabccddddde"
* "abcd" => "abcd"
* "45" => "" (некорректная строка)

Дополнительное задание: поддержка escape - последовательности
* `qwe\4\5` => `qwe45` (*)
* `qwe\45` => `qwe44444` (*)
* `qwe\\5` => `qwe\\\\\` (*)
*/

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
		esc := prev == '\\' // previous symbol was "\" ?

		switch {

		case r == '\\':
			if esc {
				sb.WriteRune(r)
				special = true // start special seq ex \\5
			} else {
				special = false // if not "\\" then cancel special sequence
			}

		case '0' <= r && r <= '9':

			if !esc || special {
				// if special seq or no "\" in previous symbol detected
				if '0' <= prev && prev <= '9' && !special {
					return "Incorrect format!" // two conseq. digits wo special sequence
				}
				for i := 1; i < int(r-'0'); i++ {
					sb.WriteRune(prev) // simply retries previous symbol R times
				}
				special = false
			} else {
				// if detects "\" in prev symbol, start special sequence
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
