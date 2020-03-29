package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

func main() {

	if time, err := ntp.Time("ru.pool.ntp.org"); err != nil {
		fmt.Fprint(os.Stderr, err, " (error code = 1)\n")
	} else {
		fmt.Printf("%s", time)
	}
}
