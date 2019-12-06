package main

import (
	"fmt"
	"os"
	"os/user"
	"github.com/ozw-sei/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello, Programing Language Monkey! %s \n", user.Username)

	repl.Start(os.Stdin, os.Stdout)
}
