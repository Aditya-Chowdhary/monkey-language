package main

import (
	"fmt"
	"os"

	"github.com/Aditya-Chowdhary/Monkey-Interpreter/file"
	"github.com/Aditya-Chowdhary/Monkey-Interpreter/repl"
)

func main() {
	if len(os.Args) == 1 {
		repl.Start(os.Stdin, os.Stdout)
	} else if len(os.Args) == 2 {
		fileName := os.Args[1]

		fileContent, err := file.ReadCode(fileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		file.RunCode(fileContent)

	} else {
		fmt.Println("Error: Too many arguments!")
		fmt.Println(`Usage: "./monkey [file-path]"`)
	}
}
