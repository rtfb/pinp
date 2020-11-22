package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNextToken(t *testing.T) {
	input := `program foo;
begin
	WriteLn('hi');
end;`
	l := New(NewFileFromString(input))
	assert.Equal(t, "program", l.NextToken().Literal)
	assert.Equal(t, "foo", l.NextToken().Literal)
	assert.Equal(t, ";", l.NextToken().Literal)
	assert.Equal(t, "begin", l.NextToken().Literal)
	assert.Equal(t, "WriteLn", l.NextToken().Literal)
	assert.Equal(t, "(", l.NextToken().Literal)
	assert.Equal(t, "hi", l.NextToken().Literal)
	assert.Equal(t, ")", l.NextToken().Literal)
	assert.Equal(t, ";", l.NextToken().Literal)
	assert.Equal(t, "end", l.NextToken().Literal)
	assert.Equal(t, ";", l.NextToken().Literal)
}

func TestTokenPos(t *testing.T) {
	input := `program foo;
begin
end;`
	want := []struct {
		literal string
		pos     string
	}{
		{"program", "<input>:1:1"},
		{"foo", "<input>:1:9"},
		{";", "<input>:1:12"},
		{"begin", "<input>:2:1"},
		{"end", "<input>:3:1"},
		{";", "<input>:3:4"},
	}
	l := New(NewFileFromString(input))
	for _, tt := range want {
		tok := l.NextToken()
		assert.Equal(t, tt.literal, tok.Literal)
		assert.Equal(t, tt.pos, tok.Pos.String())
	}
}

func TestTokenPosFromFile(t *testing.T) {
	want := []struct {
		literal string
		pos     string
	}{
		{"program", "emptyprogram.pinp:1:1"},
		{"foo", "emptyprogram.pinp:1:9"},
		{";", "emptyprogram.pinp:1:12"},
		{"begin", "emptyprogram.pinp:2:1"},
		{"end", "emptyprogram.pinp:3:1"},
		{";", "emptyprogram.pinp:3:4"},
	}
	f, err := NewFile("./testdata/emptyprogram.pinp")
	require.NoError(t, err)
	l := New(f)
	for _, tt := range want {
		tok := l.NextToken()
		assert.Equal(t, tt.literal, tok.Literal)
		assert.Equal(t, tt.pos, tok.Pos.String())
	}
}
