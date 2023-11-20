package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/MBATheGamer/mba_lang/repl"
)

func main() {
	var user, err = user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the MBA Programming Language!\n",
		user.Username,
	)

	fmt.Printf("Feel free to type in commands\n")

	repl.Start(os.Stdin, os.Stdout)
}
