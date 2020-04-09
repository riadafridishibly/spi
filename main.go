// golang implementation of simple pascal interpreter
package main

import (
	"bufio"
	"fmt"
	"github.com/riadafridishibly/spi/lexer"
	"github.com/riadafridishibly/spi/parser"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	out := func() bool {
		fmt.Print("spi> ")
		return true
	}

	for out() && scanner.Scan() {

		txt := scanner.Text()
		if strings.TrimSpace(txt) == "" {
			continue
		}

		lx := lexer.NewLexer(txt, 0)
		prsr := parser.NewParser(lx)
		res := prsr.Expr()

		fmt.Println(res)
	}
}
