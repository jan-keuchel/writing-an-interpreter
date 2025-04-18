package token

type TokenType  int

const (

	EOF TokenType = iota
	INVALID

	INT
	FLOAT
	STRING
	BOOL
	CHAR

	IDENT

	IF
	ELSE
	FALSE
	TRUE
	FOR
	NIL
	PRINT
	RETURN

	DOT
	COMMA
	SEMICOLON
	LPAREN //(
	RPAREN //)
	LBRACE //{
	RBRACE //}
	LBRACK //[
	RBRACK //]

	MUL
	PLUS
	MINUS
	DIV

	EQUAL
	EQUAL_EQUAL
	NOT
	NOT_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL
	PLUS_EQUAL
	MINUS_EQUAL
	MUL_EQUAL
	DIV_EQUAL

)

type Token struct {
	Type 	TokenType
	Value	string
}

func NewToken(tokenType TokenType, value string) *Token {
	return &Token{
		Type:  tokenType,
		Value: value,
	}
}



func (t TokenType) String() string {
	switch t {
	case EOF: return "EOF"
	case INVALID: return "INVALID"
	case INT: return "INT"
	case FLOAT: return "FLOAT"
	case STRING: return "STRING"
	case BOOL: return "BOOL"
	case CHAR: return "CHAR"
	case IDENT: return "IDENT"
	case DOT: return "DOT"
	case COMMA: return "COMMA"
	case SEMICOLON: return "SEMICOLON"
	case LPAREN: return "LPAREN"
	case RPAREN: return "RPAREN"
	case LBRACE: return "LBRACE"
	case RBRACE: return "RBRACE"
	case LBRACK: return "LBRACK"
	case RBRACK: return "RBRACK"
	case MUL: return "MUL"
	case PLUS: return "PLUS"
	case MINUS: return "MINUS"
	case DIV: return "DIV"
	case EQUAL: return "EQUAL"
	case EQUAL_EQUAL: return "EQUAL_EQUAL"
	case NOT: return "NOT"
	case NOT_EQUAL: return "NOT_EQUAL"
	case GREATER: return "GREATER"
	case GREATER_EQUAL: return "GREATER_EQUAL"
	case LESS: return "SMALLER"
	case LESS_EQUAL: return "SMALLER_EQUAL"
	case PLUS_EQUAL: return "PLUS_EQUAL"
	case MINUS_EQUAL: return "MINUS_EQUAL"
	case MUL_EQUAL: return "MUL_EQUAL"
	case DIV_EQUAL: return "DIV_EQUAL"
	case IF: return "IF"
	case ELSE: return "ELSE"
	case FALSE: return "FALSE"
	case TRUE: return "TRUE"
	case FOR: return "FOR"
	case NIL: return "NIL"
	case PRINT: return "PRINT"
	case RETURN: return "RETURN"
	default: return "UNKNOWN"
	}
}
