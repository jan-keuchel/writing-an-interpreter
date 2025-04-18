package lexer

import (
	"fmt"

	"github.com/jan-keuchel/writing-an-interpreter/src/token"
	"github.com/jan-keuchel/writing-an-interpreter/src/utils"
)

type Lexer struct {
	code  		string
	tokens  	[]*token.Token
	start  		int
	current  	int
	line   		int
	keywords 	map[string]token.TokenType
}

func NewLexer(code string) *Lexer {
	return &Lexer{
		code:  		code,
		tokens: 	make([]*token.Token, 0),
		start: 		0,
		current: 	0,
		line: 		1,
		keywords:  	map[string]token.TokenType {
			"if": 		token.IF,
			"else": 	token.ELSE,
			"false": 	token.FALSE,
			"true": 	token.TRUE,
			"for": 		token.FOR,
			"nil": 		token.NIL,
			"print": 	token.PRINT,
			"return": 	token.RETURN,
			"int": 		token.INT,
			"float": 	token.FLOAT,
			"string": 	token.STRING,
			"bool": 	token.BOOL,
			"char": 	token.CHAR,
		},
	}
}

func (l *Lexer) isAtEnd() bool {
	return l.current >= len(l.code)
}

func (l *Lexer) advance() rune {
	l.current++
	return rune(l.code[l.current-1])
}

func (l *Lexer) peek() rune {
	return rune(l.code[l.current])
}

func (l *Lexer) peekNext() rune {
	if l.current + 1 >= len(l.code) {
		return rune(0)
	}
	return rune(l.code[l.current+1])
}

func (l *Lexer) match(expected rune) bool {
	if l.isAtEnd() {return false}
	if expected != rune(l.code[l.current]) {return false}

	l.current++
	return true
}

func (l *Lexer) addToken(tokenType token.TokenType, 
	tokenValue string) {
	l.tokens = append(l.tokens, token.NewToken(tokenType, tokenValue))
}

func (l *Lexer) prepToken(tokenType token.TokenType) {
	l.addToken(tokenType, "")
}

func (l *Lexer) lexComment() {
	for !l.isAtEnd() && l.peek() != '\n' { l.advance() }
}
func (l *Lexer) lexString() {

	for !l.isAtEnd() && l.peek() != '"' { 
		if l.peek() == '\n' {
			l.line++
		}
		l.advance() 
	}
	if l.isAtEnd() {
		utils.Error(l.line, "Unterminated string.")
		return 
	}
	l.advance()
	l.addToken(token.STRING, l.code[l.start+1:l.current-1])

}
func (l *Lexer) lexNumber() {

	isFloatingPoint := false
	for !l.isAtEnd() && l.isDigit(l.peek()) { l.advance() }
	if !l.isAtEnd() && l.peek() == '.' && l.isDigit(l.peekNext()) {
		l.advance()
		isFloatingPoint = true
	}
	for !l.isAtEnd() && l.isDigit(l.peek()) { l.advance() }

	if isFloatingPoint {
		l.addToken(token.FLOAT, l.code[l.start:l.current])
	} else {
		l.addToken(token.INT, l.code[l.start:l.current])	
	}

}
func (l *Lexer) lexIdentifier() {

	for !l.isAtEnd() && l.isAlphaNumeric(l.peek()) { l.advance() }
	lexeme := l.code[l.start:l.current]

	tokenType, exists := l.keywords[lexeme]
	if exists {
		l.prepToken(tokenType)
	} else {
		l.addToken(token.IDENT, lexeme)
	}

}

func (l *Lexer) isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
func (l *Lexer) isAlpha(r rune) bool {
	return r == '_' || 
	(r >= 'a' && r <= 'z') || 
	(r >= 'A' && r <= 'Z')
}
func (l *Lexer) isAlphaNumeric(r rune) bool {
	return l.isAlpha(r) || l.isDigit(r)
}

func (l *Lexer) scanToken() {

	c := l.advance()

	switch c {
	// Single character tokens
	case '(': l.prepToken(token.LPAREN)
	case ')': l.prepToken(token.RPAREN)
	case '{': l.prepToken(token.LBRACE)
	case '}': l.prepToken(token.RBRACE)
	case '[': l.prepToken(token.LBRACK)
	case ']': l.prepToken(token.RBRACK)
	case ';': l.prepToken(token.SEMICOLON)
	case ',': l.prepToken(token.COMMA)
	case '.': l.prepToken(token.DOT)

	// Dual character tokens
	case '=': {  	if l.match('=')  	{l.prepToken(token.EQUAL_EQUAL)}  	else {l.prepToken(token.EQUAL)}  	}
	case '!': {  	if l.match('=')  	{l.prepToken(token.NOT_EQUAL)}  	else {l.prepToken(token.NOT)}  		}
	case '<': {  	if l.match('=')  	{l.prepToken(token.LESS_EQUAL)}  	else {l.prepToken(token.LESS)}  	}
	case '>': {  	if l.match('=')  	{l.prepToken(token.GREATER_EQUAL)} 	else {l.prepToken(token.GREATER)}  	}
	case '+': {  	if l.match('=')  	{l.prepToken(token.PLUS_EQUAL)}  	else {l.prepToken(token.PLUS)}  	}
	case '-': {  	if l.match('=')  	{l.prepToken(token.MINUS_EQUAL)}  	else {l.prepToken(token.MINUS)}  	}
	case '*': {  	if l.match('=')  	{l.prepToken(token.MUL_EQUAL)}  	else {l.prepToken(token.MUL)}  		}

	// Whitespaces
	case '\t', '\r', ' ':
	case '\n': l.line++

	// Multi character tokens
	// Comments
	case '/': {
		if l.match('=') {
			l.prepToken(token.DIV_EQUAL)
		} else if l.match('/') {
			l.lexComment()
		} else {
			l.prepToken(token.DIV)
		}
	}

	// String
	case '"': l.lexString()

	// Numbers
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9': l.lexNumber()

	default: {
		if l.isAlpha(c) {
			l.lexIdentifier()
		} else {utils.Error(l.line, "Unidentified symbol.")}
	}

	}

}

func (l *Lexer) LexCode() []*token.Token {

	for !l.isAtEnd() {
		l.start = l.current
		l.scanToken()
	}
	l.addToken(token.EOF, "")

	for _, token := range l.tokens {
		fmt.Printf("Type: %-15s Value: %s\n", token.Type, token.Value)
	}

	return l.tokens

}


