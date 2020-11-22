package lexer

import (
	"io/ioutil"
	"path/filepath"
)

// File represents one file of a parsed source code. If the source code is
// being parsed from a string (like in tests), the filename will be a stub.
type File struct {
	name     string
	fullPath string
	input    string
	pos      Cursor
	ch       byte // current char under examination
}

// Cursor holds together several variables representing different aspects of
// file position (like line number, column number and absolute position).
type Cursor struct {
	position     int // current position in input (points to current char)
	readPosition int // current reading position in input (after current char)
	line         int // line number, 1-based
	col          int // column number, 1-based
}

func NewFile(filename string) (File, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return File{}, err
	}
	absPath, err := filepath.Abs(filename)
	if err != nil {
		return File{}, err
	}
	return File{
		name:     filepath.Base(filename),
		fullPath: absPath,
		input:    string(bytes),
		pos: Cursor{
			line: 1,
			col:  1,
		},
	}, nil
}

func NewFileFromString(input string) File {
	return File{
		name:  "<input>",
		input: input,
		pos: Cursor{
			line: 1,
			col:  1,
		},
	}
}

func (f *File) readChar() {
	if f.pos.readPosition >= len(f.input) {
		f.ch = 0
	} else {
		f.ch = f.input[f.pos.readPosition]
	}
	f.pos.position = f.pos.readPosition
	f.pos.readPosition++
	f.pos.col++
}

func (f *File) peekChar() byte {
	if f.pos.readPosition >= len(f.input) {
		return 0
	}
	return f.input[f.pos.readPosition]
}

func (f *File) skipWhitespace() {
	for {
		switch {
		case f.ch == ' ':
		case f.ch == '\t':
		case f.ch == '\n' || f.ch == '\r': // TODO: this will not work correctly with Windows-style EOLs
			f.pos.line++
			f.pos.col = 1
		default:
			return
		}
		f.readChar()
	}
}
