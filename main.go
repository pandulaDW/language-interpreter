package main

import (
	"fmt"
	"github.com/pandulaDW/language-interpreter/repl"
	"os"
	"os/user"
)

func main() {
	cUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s!. This is the Monkey Programming Language. \n", cUser.Username)
	fmt.Print("Feel free to type in commands\n")

	repl.Start(os.Stdin, os.Stdout)
}
