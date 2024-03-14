package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/Aditya-Chowdhary/Monkey-Interpreter/ast"
	"github.com/Aditya-Chowdhary/Monkey-Interpreter/compiler"
	"github.com/Aditya-Chowdhary/Monkey-Interpreter/evaluator"
	"github.com/Aditya-Chowdhary/Monkey-Interpreter/file"
	"github.com/Aditya-Chowdhary/Monkey-Interpreter/lexer"
	"github.com/Aditya-Chowdhary/Monkey-Interpreter/object"
	"github.com/Aditya-Chowdhary/Monkey-Interpreter/parser"
	"github.com/Aditya-Chowdhary/Monkey-Interpreter/vm"
)

var engine = flag.String("engine", "", "use `vm` or `eval`. Leave blank to return comparison")

var input = `
	let fibonacci = fn(x) {
	if (x == 0) {
	0
	} else {
	if (x == 1) {
	return 1;
	} else {
	fibonacci(x - 1) + fibonacci(x - 2);
	}
	}
	};
	fibonacci(35);
	`

type benchResults struct {
	engine   string
	result   object.Object
	duration time.Duration
}

func main() {
	flag.Parse()

	if flag.Arg(0) != "" {
		code, err := file.ReadCode(flag.Arg(0))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		input = code
	}

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	if *engine == "vm" {
		timeVM(program)
	} else if *engine == "eval" {
		timeEval(program)
	} else {
		results := make(chan benchResults)
		go func() {
			fmt.Println("Running eval")
			result := timeEval(program)
			results <- result
			}()
			go func() {
			fmt.Println("Running vm")
			result := timeVM(program)
			results <- result
		}()

		printResult(<-results)
		printResult(<-results)
	}
}

func timeVM(program *ast.Program) benchResults {
	comp := compiler.New()
	err := comp.Compile(program)
	if err != nil {
		fmt.Printf("compiler error: %s", err)
		return benchResults{}
	}

	machine := vm.New(comp.Bytecode())

	start := time.Now()

	err = machine.Run()
	if err != nil {
		fmt.Printf("vm error: %s", err)
		return benchResults{}
	}

	duration := time.Since(start)
	result := machine.LastPoppedStackElem()
	return benchResults{engine: "vm", result: result, duration: duration}
}

func timeEval(program *ast.Program) benchResults {
	env := object.NewEnvironment()
	start := time.Now()
	result := evaluator.Eval(program, env)
	duration := time.Since(start)
	return benchResults{engine: "eval", result: result, duration: duration}
}

func printResult(result benchResults) {
	fmt.Printf("--> engine=%s, result=%s, duration=%s\n",
		result.engine,
		result.result.Inspect(),
		result.duration)
}
