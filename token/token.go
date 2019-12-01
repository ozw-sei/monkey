package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"
	
	// 識別子
	IDENT = "IDENT"
	INT = "INT"
	
	// Token
	ASSIGN = "="
	PLUS = "+"
	
	// デリミタ
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"


	FUNCTION = "FUNCTION"
	LET = "LET"
)