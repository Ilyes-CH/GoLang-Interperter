package main

import (
	"fmt"
	"ilymar/repl"
	"os"
	"os/user"
)



func main(){

	user, err := user.Current()
	if err != nil{
		panic(err)
	}
	fmt.Printf("Hello %s! This is IlymarLang Programming Language\n",user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout) //instructing the func to read from the console and prints to the console
}

