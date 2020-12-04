package main

import (
	"fmt"

	"github.com/rtfb/pinp/lexer"
	"github.com/rtfb/pinp/parser"
)

var prog = `program foo;
begin
	WriteLn('hello');
end;`

// TODO: this valid no-op program causes errors:
// var prog = `program foo;
// begin
// end;`

func main() {
	f := lexer.NewFileFromString(prog)
	l := lexer.New(f)
	p := parser.New(l)
	ast := p.ParseProgram()
	for _, e := range p.Errors() {
		fmt.Println("E:", e)
	}
	fmt.Println(ast)
}
