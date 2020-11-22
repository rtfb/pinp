package lexer

import "github.com/rtfb/pinp/token"

// Lexer manages the lexical analysis of the input stream.
type Lexer struct {
	f File
}

// New creates a lexer.
func New(file File) *Lexer {
	l := Lexer{f: file}
	l.f.readChar()
	return &l
}

// NextToken reads and returns the next token.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.f.skipWhitespace()
	switch l.f.ch {
	case '=':
		if l.f.peekChar() == '=' {
			ch := l.f.ch
			l.f.readChar()
			literal := string(ch) + string(l.f.ch)
			tok = token.Token{
				Type:    token.Equals,
				Literal: literal,
				Pos:     l.tokPos(),
			}
		} else {
			tok = l.newToken(token.Assign, l.f.ch)
		}
	case ';':
		tok = l.newToken(token.Semicolon, l.f.ch)
	case ':':
		tok = l.newToken(token.Colon, l.f.ch)
	case '(':
		tok = l.newToken(token.LParen, l.f.ch)
	case ')':
		tok = l.newToken(token.RParen, l.f.ch)
	case '{':
		tok = l.newToken(token.LBrace, l.f.ch)
	case '}':
		tok = l.newToken(token.RBrace, l.f.ch)
	case '[':
		tok = l.newToken(token.LBracket, l.f.ch)
	case ']':
		tok = l.newToken(token.RBracket, l.f.ch)
	case ',':
		tok = l.newToken(token.Comma, l.f.ch)
	case '+':
		tok = l.newToken(token.Plus, l.f.ch)
	case '-':
		tok = l.newToken(token.Minus, l.f.ch)
	case '!':
		if l.f.peekChar() == '=' {
			ch := l.f.ch
			l.f.readChar()
			literal := string(ch) + string(l.f.ch)
			tok = token.Token{
				Type:    token.NotEquals,
				Literal: literal,
				Pos:     l.tokPos(),
			}
		} else {
			tok = l.newToken(token.Bang, l.f.ch)
		}
	case '/':
		tok = l.newToken(token.Slash, l.f.ch)
	case '%':
		tok = l.newToken(token.Modulo, l.f.ch)
	case '*':
		tok = l.newToken(token.Asterisk, l.f.ch)
	case '<':
		tok = l.newToken(token.LT, l.f.ch)
	case '>':
		tok = l.newToken(token.GT, l.f.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	case '\'':
		tok.Pos = l.tokPos()
		tok.Type = token.String
		tok.Literal = l.readString()
	default:
		if isLetter(l.f.ch) {
			tok.Pos = l.tokPos()
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.f.ch) {
			tok.Pos = l.tokPos()
			tok.Literal = l.readNumber()
			tok.Type = token.Num
			return tok
		} else {
			tok = l.newToken(token.Illegal, l.f.ch)
		}
	}
	l.f.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.f.pos.position
	for isLetter(l.f.ch) {
		l.f.readChar()
	}
	return l.f.input[position:l.f.pos.position]
}

func (l *Lexer) readNumber() string {
	position := l.f.pos.position
	for isDigit(l.f.ch) {
		l.f.readChar()
	}
	return l.f.input[position:l.f.pos.position]
}

func (l *Lexer) readString() string {
	position := l.f.pos.position + 1
	for {
		l.f.readChar()
		if l.f.ch == '\'' || l.f.ch == 0 {
			break
		}
	}
	return l.f.input[position:l.f.pos.position]
}

func (l *Lexer) newToken(tokType token.Type, ch byte) token.Token {
	return token.Token{
		Type:    tokType,
		Literal: string(ch),
		Pos:     l.tokPos(),
	}
}

func (l *Lexer) tokPos() token.Position {
	pos := l.f.pos
	return token.Position{
		File: l.f.name,
		Line: pos.line,
		Col:  pos.col - 1,
		Pos:  pos.position - 1,
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
