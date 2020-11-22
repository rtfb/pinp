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
	l := lexer.New(prog)
	tok := l.NextToken()
	for tok.Type != token.EOF {
		fmt.Println(tok)
		tok = l.NextToken()
	}
}
