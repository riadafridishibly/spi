package parser

// TODO: Update test for the `ast.Node`

// import (
//     "github.com/riadafridishibly/spi/ast"
//     "github.com/riadafridishibly/spi/lexer"
//     "testing"
// )

// func TestFactor(t *testing.T) {
//     testcase := []struct {
//         text string
//         want int64
//     }{
//         {"4", 4},
//     }
//
//     for _, tt := range testcase {
//
//         lx := lexer.NewLexer(tt.text, 0)
//         prsr := NewParser(lx)
//         tree := prsr.Parse()
//
//         v := ast.Walk(tree)
//
//         t.Error(v)
//
//         // if value.Value.(int64) != tt.want {
//         //     t.Errorf("TestExpr: got [%+v] want [%+v]\n", value, tt.want)
//         // }
//     }
//
// }

// func TestExpr(t *testing.T) {
//     testcase := []struct {
//         text string
//         want int64
//     }{
//         {"(1 + 2) * 3", 9},
//         {"9 - 2", 7},
//         {"(1 + 3) / 4 * 3", 3},
//     }
//
//     for _, tt := range testcase {
//
//         lx := lexer.NewLexer(tt.text, 0)
//         prsr := NewParser(lx)
//
//         value := prsr.Expr()
//
//         if value != tt.want {
//             t.Errorf("TestExpr: got [%+v] want [%+v]\n", value, tt.want)
//         }
//     }
//
// }
