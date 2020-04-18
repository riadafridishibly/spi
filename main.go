// golang implementation of simple pascal interpreter
package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/riadafridishibly/spi/ast"
	"github.com/riadafridishibly/spi/lexer"
	"github.com/riadafridishibly/spi/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./spi <program-name>")
		os.Exit(1)
	}

	filename := os.Args[1]
	filedata, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	txt := string(filedata)

	lx := lexer.NewLexer(txt, 0)
	prsr := parser.NewParser(lx)
	tree := prsr.Parse()
	ast.Walk(tree)
}
