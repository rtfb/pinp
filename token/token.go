package token

// A set of token types.
const (
	Illegal = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and literals
	Ident  = "IDENT" // add, foo, bar, x, y, etc
	Num    = "NUM"   // numeric literal
	String = "STRING"

	// Operators
	Assign    = "="
	Plus      = "+"
	Minus     = "-"
	Bang      = "!"
	Asterisk  = "*"
	Slash     = "/"
	Modulo    = "%"
	LT        = "<"
	GT        = ">"
	Equals    = "=="
	NotEquals = "!="

	// Delimiters
	Comma     = ","
	Semicolon = ";"
	LParen    = "("
	RParen    = ")"
	LBrace    = "{"
	RBrace    = "}"
	LBracket  = "["
	RBracket  = "]"
	Colon     = ":"

	// Keywords
	Function = "FUNCTION"
	Var      = "VAR"
	True     = "TRUE"
	False    = "FALSE"
	If       = "IF"
	Else     = "ELSE"
	Return   = "RETURN"
	Begin    = "BEGIN"
	End      = "END"
	Do       = "DO"
	While    = "WHILE"
)

var keywords = map[string]Type{
	"function": Function,
	"var":      Var,
	"true":     True,
	"false":    False,
	"if":       If,
	"else":     Else,
	"return":   Return,
	"begin":    Begin,
	"end":      End,
	"do":       Do,
	"while":    While,
}

// Type identifies a token type.
type Type string

// Token represents a token.
type Token struct {
	Type    Type
	Literal string
}

// LookupIdent returns an appropriate token type for identifier: if it's a
// reserved keyword, it will return that keyword's type, otherwise Ident.
func LookupIdent(ident string) Type {
	if tokType, ok := keywords[ident]; ok {
		return tokType
	}
	return Ident
}
