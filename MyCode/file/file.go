package file

import (
	"fmt"
	"os"

	"github.com/Aditya-Chowdhary/Monkey-Interpreter/compiler"
	"github.com/Aditya-Chowdhary/Monkey-Interpreter/lexer"
	"github.com/Aditya-Chowdhary/Monkey-Interpreter/parser"
	"github.com/Aditya-Chowdhary/Monkey-Interpreter/vm"
)

func RunCode(fileCode string) {
	l := lexer.New(fileCode)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(p.Errors())
		os.Exit(1)
	}

	comp := compiler.New()
	err := comp.Compile(program)
	if err != nil {
		fmt.Printf("Woops! Compilation failed: \n %s\n", err)
		os.Exit(1)
	}

	code := comp.Bytecode()

	machine := vm.New(code)

	err = machine.Run()
	if err != nil {
		fmt.Printf("Woops! Executing bytecode failed:\n %s\n", err)
		os.Exit(1)
	}
}

const MONKEY_FACE = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

func printParserErrors(errors []string) {
	fmt.Print(MONKEY_FACE)
	fmt.Println("Whoops! We ran into some monkey business here!")
	for _, msg := range errors {
		fmt.Println("\t" + msg + "\n")
	}
}
