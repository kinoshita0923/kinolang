package token

type TokenType string

type Token struct {
	Type	TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF		= "EOF"

	// 識別子 + リテラル
	IDENT	= "IDENT"
	INT		= "INT"
	STRING	= "STRING"
	DOUBLE  = "DOUBLE"

	// 演算子
	ASSIGN		= "="
	PLUS		= "+"
	MINUS		= "-"
	BANG		= "!"
	ASTERISK	= "*"
	SLASH		= "/"

	LT = "<"
	GT = ">"

	EQ		= "=="
	NOT_EQ	= "!="
	MORE    = "<="
	LESS    = ">="

	INCREMENT = "++"
	DECREMENT = "--"

	// デリミタ
	COMMA		= ","
	SEMICOLON	= ";"
	COLON       = ":"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// キーワード
	FUNCTION = "FUNCTION"
	LET		 = "LET"
	TRUE	 = "TRUE"
	FALSE	 = "FALSE"
	IF		 = "IF"
	ELSE	 = "ELSE"
	ELIF     = "ELIF"
	RETURN	 = "RETURN"
	MACRO    = "MACRO"
	FOR      = "FOR"
	WHILE    = "WHILE"
	CLASS    = "CLASS"
	NEW      = "NEW"
)

var keywords = map[string]TokenType {
	"fn": FUNCTION,
	"let": LET,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,
	"elif": ELIF,
	"return": RETURN,
	"macro": MACRO,
	"for": FOR,
	"while": WHILE,
	"class": CLASS,
	"new": NEW,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}