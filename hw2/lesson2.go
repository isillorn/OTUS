package main

import (
	"fmt"
	"strings"
)

func itoa(i int) (s string) {

	var symbols []byte
	var sb strings.Builder

	for tmp := i; tmp > 0; {
		symbols = append(symbols, byte(tmp%10))
		tmp = (tmp - tmp%10) / 10
		//		fmt.Printf("%v\n", symbols)
	}

	for ind := len(symbols) - 1; ind >= 0; ind-- {
		fmt.Fprintf(&sb, "%d", symbols[ind])
	}

	s = sb.String()
	return s
}

func itoa2(i int) (s string) {

	var sb strings.Builder

	if i > 0 {
		//		fmt.Printf("i = %d\n", i)
		sb.WriteString(itoa2((i - i%10) / 10))
		fmt.Fprintf(&sb, "%d", i%10)
	}

	s = sb.String()
	return s
}

func main() {
	fmt.Println(itoa(513))
	fmt.Println(itoa2(513))
}
