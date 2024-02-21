package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/Aditya-Chowdhary/Monkey-Interpreter/file"
	"github.com/Aditya-Chowdhary/Monkey-Interpreter/repl"
)

func main() {
	if len(os.Args) == 1 {
		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Hello %s! This is the Monkey Programming Language!\n", user.Username)
		fmt.Printf("Feel free to type in commands!\n")
		repl.Start(os.Stdin, os.Stdout)
	} else if len(os.Args) == 2 {
		fileName := os.Args[1]

		fileContent, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Println("Error: Invalid file name")
			os.Exit(1)
		}
		if !strings.HasSuffix(fileName, ".mky") {
			fmt.Println("Error: Invalid file name")
			os.Exit(1)
		}
		file.RunCode(string(fileContent))
	} else {
		fmt.Println("Error: Too many arguments!")
		fmt.Println(`Usage: "./monkey [file-path]"`)
	}
}
