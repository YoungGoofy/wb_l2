package echo

import (
	"fmt"
)

func Run(args []string) {
	fmt.Println(echo(args))
}

func echo(args []string) string {
	var offer string
	for i := 1; i < len(args); i++ {
		offer += args[i] + " "
	}
	return offer
}
