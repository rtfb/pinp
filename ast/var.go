package ast

import (
	"bytes"

	"github.com/rtfb/pinp/token"
)

// VarStatement is the AST subtree containing a let statement.
type VarStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *VarStatement) statementNode() {
}

// TokenLiteral implements Node.
func (ls *VarStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// String implements Node.
func (ls *VarStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}
