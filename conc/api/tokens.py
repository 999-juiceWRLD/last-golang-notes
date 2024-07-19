from enum import Enum

class TokenType(Enum):
    VAR     = "VAR"
    LETEM   = "LETEM"
    
    # types
    INT = "INT"
    FLOAT = "FLOAT"
    STRING = "STRING"
    BOOL = "BOOL"
    
    # identifier
    IDENT = "IDENT"
    
    # delimiters
    COMMA       = ":"
    SEMICOLON   = ";"
    LPAREN      = "("
    RPAREN      = ")"
    LBRACE      = "{"
    RBRACE      = "}"

    # operators
    ASSIGN		= "="
    PLUS		= "+"
    MINUS		= "-"
    BANG		= "!"
    ASTERISK	= "*"	
    SLASH		= "/"
    LT			= "<"
    GT			= ">"

    # keywords
    FUNCTION	= "FUNCTION"
    LET			= "LET"
    TRUE 		= "TRUE"
    FALSE		= "FALSE"
    IF 			= "IF"
    ELSE 		= "ELSE"
    RETURN		= "RETURN"
    
    # else
    ILLEGAL     = "ILLEGAL"
    EOF         = "EOF"
 