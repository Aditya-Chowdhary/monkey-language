This is a project to write an interpreter and a compiler for a made up language Monkey. The repository follows the book "Writing an Interpreter in Go" until Tag `Book1End`, and its sequel "Writing a Compiler in Go" until Tag `Book2End`. All further commits are my own contribution.

- https://interpreterbook.com/
- https://compilerbook.com/

# How to run

### Linux

1. Clone the repository
2. Run the command `cd MyCode/` to change directories.
3. To run the REPL, run the command `./build/monkey`.
4. Alternatively, create a file with extension `.mky` and run the command `./build/monkey [file]` to compile and run the file

A pre-written file `test.mky` is also present in the repository

# Running benchmarks

### Linux
1. Clone the repository
2. Run the command `cd MyCode/` to change directories.
3. To run using compiler and vm, run the command `./build/fibonacci -engine=vm`
4. To run using interpreter, run the command `./build/fibonacci -engine=eval`

- Currently the benchmark is based on a pre-written fibonacci recursion program.
- The fibonacci number can be changed by going into `./benchmark/main.go Ln 30`, and editing the number in the function call.