package lexer

import (
	"testing"

	"github.com/jan-keuchel/writing-an-interpreter/src/token"
)

func TestLexerSingleCharacterTokens(t *testing.T) {
	input := `(){}[],.;`
	lexer := NewLexer(input)
	tokens := lexer.LexCode()

	expected := []struct {
		tokenType token.TokenType
		literal   string
	}{
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.LBRACK, "["},
		{token.RBRACK, "]"},
		{token.COMMA, ","},
		{token.DOT, "."},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	if len(tokens) != len(expected) {
		t.Fatalf("expected %d tokens, got %d", len(expected), len(tokens))
	}

	for i, exp := range expected {
		if tokens[i].Type != exp.tokenType {
			t.Errorf("token %d: expected type %v, got %v", i, exp.tokenType, tokens[i].Type)
		}
		if tokens[i].Literal != exp.literal {
			t.Errorf("token %d: expected literal %q, got %q", i, exp.literal, tokens[i].Literal)
		}
	}
}

func TestLexerDualCharacterTokens(t *testing.T) {
	input := `== != <= >= += -= *= /=`
	lexer := NewLexer(input)
	tokens := lexer.LexCode()

	expected := []struct {
		tokenType token.TokenType
		literal   string
	}{
		{token.EQUAL_EQUAL, "=="},
		{token.NOT_EQUAL, "!="},
		{token.LESS_EQUAL, "<="},
		{token.GREATER_EQUAL, ">="},
		{token.PLUS_EQUAL, "+="},
		{token.MINUS_EQUAL, "-="},
		{token.MUL_EQUAL, "*="},
		{token.DIV_EQUAL, "/="},
		{token.EOF, ""},
	}

	if len(tokens) != len(expected) {
		t.Fatalf("expected %d tokens, got %d", len(expected), len(tokens))
	}

	for i, exp := range expected {
		if tokens[i].Type != exp.tokenType {
			t.Errorf("token %d: expected type %v, got %v", i, exp.tokenType, tokens[i].Type)
		}
		if tokens[i].Literal != exp.literal {
			t.Errorf("token %d: expected literal %q, got %q", i, exp.literal, tokens[i].Literal)
		}
	}
}

func TestLexerNumbers(t *testing.T) {
	input := `123 45.67`
	lexer := NewLexer(input)
	tokens := lexer.LexCode()

	expected := []struct {
		tokenType token.TokenType
		literal   string
		value     any
	}{
		{token.INT, "123", 123},
		{token.FLOAT, "45.67", 45.67},
		{token.EOF, "", nil},
	}

	if len(tokens) != len(expected) {
		t.Fatalf("expected %d tokens, got %d", len(expected), len(tokens))
	}

	for i, exp := range expected {
		if tokens[i].Type != exp.tokenType {
			t.Errorf("token %d: expected type %v, got %v", i, exp.tokenType, tokens[i].Type)
		}
		if tokens[i].Literal != exp.literal {
			t.Errorf("token %d: expected literal %q, got %q", i, exp.literal, tokens[i].Literal)
		}
		if tokens[i].Value != exp.value {
			t.Errorf("token %d: expected value %v, got %v", i, exp.value, tokens[i].Value)
		}
	}
}

func TestLexerStrings(t *testing.T) {
	input := `"hello" "world\nline"`
	lexer := NewLexer(input)
	tokens := lexer.LexCode()

	expected := []struct {
		tokenType token.TokenType
		literal   string
		value     any
		line      int
	}{
		{token.STRING, `"hello"`, "hello", 1},
		{token.STRING, `"world\nline"`, "world\nline", 1},
		{token.EOF, "", nil, 1},
	}

	if len(tokens) != len(expected) {
		t.Fatalf("expected %d tokens, got %d", len(expected), len(tokens))
	}

	for i, exp := range expected {
		if tokens[i].Type != exp.tokenType {
			t.Errorf("token %d: expected type %v, got %v", i, exp.tokenType, tokens[i].Type)
		}
		if tokens[i].Literal != exp.literal {
			t.Errorf("token %d: expected literal %q, got %q", i, exp.literal, tokens[i].Literal)
		}
		if tokens[i].Value != exp.value {
			t.Errorf("token %d: expected value %v, got %v", i, exp.value, tokens[i].Value)
		}
		if tokens[i].Line != exp.line {
			t.Errorf("token %d: expected line %d, got %d", i, exp.line, tokens[i].Line)
		}
	}
}

func TestLexerKeywordsAndIdentifiers(t *testing.T) {
	input := `if else true false print variable _underscore`
	lexer := NewLexer(input)
	tokens := lexer.LexCode()

	expected := []struct {
		tokenType token.TokenType
		literal   string
		value     any
	}{
		{token.IF, "if", nil},
		{token.ELSE, "else", nil},
		{token.TRUE, "true", true},
		{token.FALSE, "false", false},
		{token.PRINT, "print", nil},
		{token.IDENT, "variable", nil},
		{token.IDENT, "_underscore", nil},
		{token.EOF, "", nil},
	}

	if len(tokens) != len(expected) {
		t.Fatalf("expected %d tokens, got %d", len(expected), len(tokens))
	}

	for i, exp := range expected {
		if tokens[i].Type != exp.tokenType {
			t.Errorf("token %d: expected type %v, got %v", i, exp.tokenType, tokens[i].Type)
		}
		if tokens[i].Literal != exp.literal {
			t.Errorf("token %d: expected literal %q, got %q", i, exp.literal, tokens[i].Literal)
		}
		if tokens[i].Value != exp.value {
			t.Errorf("token %d: expected value %v, got %v", i, exp.value, tokens[i].Value)
		}
	}
}

func TestLexerComments(t *testing.T) {
	input := `// This is a comment
print`
	lexer := NewLexer(input)
	tokens := lexer.LexCode()

	expected := []struct {
		tokenType token.TokenType
		literal   string
		line      int
	}{
		{token.PRINT, "print", 2},
		{token.EOF, "", 2},
	}

	if len(tokens) != len(expected) {
		t.Fatalf("expected %d tokens, got %d", len(expected), len(tokens))
	}

	for i, exp := range expected {
		if tokens[i].Type != exp.tokenType {
			t.Errorf("token %d: expected type %v, got %v", i, exp.tokenType, tokens[i].Type)
		}
		if tokens[i].Literal != exp.literal {
			t.Errorf("token %d: expected literal %q, got %q", i, exp.literal, tokens[i].Literal)
		}
		if tokens[i].Line != exp.line {
			t.Errorf("token %d: expected line %d, got %d", i, exp.line, tokens[i].Line)
		}
	}
}

func TestLexerUnterminatedString(t *testing.T) {
	input := `"unterminated`
	lexer := NewLexer(input)
	tokens := lexer.LexCode()

	if len(tokens) != 1 || tokens[0].Type != token.EOF {
		t.Errorf("expected only EOF token for unterminated string, got %v", tokens)
	}
}

func TestLexerInvalidSymbol(t *testing.T) {
	input := `@`
	lexer := NewLexer(input)
	tokens := lexer.LexCode()

	if len(tokens) != 1 || tokens[0].Type != token.EOF {
		t.Errorf("expected only EOF token for invalid symbol, got %v", tokens)
	}
}
