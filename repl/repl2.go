package repl

import (
	"compiler/evaluator"
	"compiler/lexer"
	"compiler/object"
	"compiler/parser"
	"io"
	"io/ioutil"
)

const PROMPT2 = ">> "

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Start2(inFile string, out io.Writer) {
	env := object.NewEnvironment()
	inputFileData, err := ioutil.ReadFile(inFile)
	check(err)
	l := lexer.New(string(inputFileData))
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors2(out, p.Errors())
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		io.WriteString(out, evaluated.Inspect())
		io.WriteString(out, "\n")
	} else {
		panic("Evaluation Error")
	}

}

func printParserErrors2(out io.Writer, errors []string) {
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
