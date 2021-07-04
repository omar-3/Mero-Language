package repl

import (
	"compiler/compiler"
	"compiler/lexer"
	"compiler/object"
	"compiler/parser"
	"compiler/vm"
	"fmt"
	"io"
	"io/ioutil"
)

const PROMPT = ">> "

func Start(in string, out io.Writer) {
	inputFileData, _ := ioutil.ReadFile(in)

	constants := []object.Object{}
	globals := make([]object.Object, vm.GlobalsSize)

	symbolTable := compiler.NewSymbolTable()
	for i, v := range object.Builtins {
		symbolTable.DefineBuiltin(i, v.Name)
	}

	l := lexer.New(string(inputFileData))
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(out, p.Errors())
	}

	comp := compiler.NewWithState(symbolTable, constants)
	err := comp.Compile(program)
	if err != nil {
		fmt.Fprintf(out, "Woops! Compilation failed:\n %s\n", err)
	}

	code := comp.Bytecode()
	constants = code.Constants

	machine := vm.NewWithGlobalsStore(code, globals)
	err = machine.Run()
	if err != nil {
		fmt.Fprintf(out, "Woops! Executing bytecode failed:\n %s\n", err)
	}

	lastPopped := machine.LastPoppedStackElem()
	io.WriteString(out, lastPopped.Inspect())
	io.WriteString(out, "\n")

	io.WriteString(out, "----------------------------\n\n")

}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
