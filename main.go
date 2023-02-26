package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"kinolang/repl"
	"kinolang/lexer"
	"kinolang/parser"
	"kinolang/evaluator"
	"kinolang/object"
)

func main() {
	if len(os.Args) < 2 {
		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Hello %s! This is the Kinolang programming language!\n",
			user.Username)
		fmt.Printf("Feel free to type in commands\n")
		repl.Start(os.Stdin, os.Stdout)
	} else {
		// ファイルをオープン
		absPath, err := filepath.Abs(os.Args[1])
		file, err := os.Open(absPath)
		if err != nil {
			fmt.Println(os.Args[1] + ": no such file or directory")
		}
		defer file.Close()

		// ファイルからソースを読み込む
		source, err := ioutil.ReadAll(file)

		e := filepath.Ext(absPath)
		if e != ".kn" {
			fmt.Printf("different extensions. want='.kn', got='%s'\n", e)
			return
		}

		env := object.NewEnvironment()
		l := lexer.New(string(source))
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(os.Stdout, p.Errors())
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil && evaluated.Type() != object.NULL_OBJ {
			io.WriteString(os.Stdout, evaluated.Inspect())
			io.WriteString(os.Stdout, "\n")
		}
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

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}