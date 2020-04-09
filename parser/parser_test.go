package parser

import (
	"github.com/riadafridishibly/spi/lexer"
	"testing"
)

func TestExpr(t *testing.T) {
	testcase := []struct {
		text string
		want int64
	}{
		{"(1 + 2) * 3", 9},
		{"9 - 2", 7},
		{"(1 + 3) / 4 * 3", 3},
	}

	for _, tt := range testcase {

		lx := lexer.NewLexer(tt.text, 0)
		prsr := NewParser(lx)

		value := prsr.Expr()

		if value != tt.want {
			t.Errorf("TestExpr: got [%+v] want [%+v]\n", value, tt.want)
		}
	}

}
