package main

import (
	"fmt"

	"github.com/rtfb/pinp/lexer"
	"github.com/rtfb/pinp/token"
)

var prog = `program foo;
begin
	WriteLn('hello');
end;`

func main() {
	f := lexer.NewFileFromString(prog)
	l := lexer.New(f)
	tok := l.NextToken()
	for tok.Type != token.EOF {
		fmt.Println(tok)
		tok = l.NextToken()
	}
}
