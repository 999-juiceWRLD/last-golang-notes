from typing import Union
from tokens import TokenType
import sys

token_array: list = ([])
print(TokenType.COMMA)

class Token:
    
    def __init__(self, type):
        self.type = type
    
    def __repr__(self) -> str:
        return f"{self.type}" 
    
    current_char_idx: int   =   0
    next_char_idx: int      =   -1
    
    current_char: str       =   ""
    next_char: str          =   ""

class Lexer:
    
    current_char_idx: int   =   0
    next_char_idx: int      =   1
    
    should_stop: bool       =   False
    
    current_char: str       =   ""
    next_char: str          =   ""
    
    def __init__(self, string: str, length: Union[int, None] = None) -> None:
        self.string = string
        if length is not None:
            assert type(length) == int
            assert len(string) == self._get_length
            self.length = self._get_length
        else:
            self.length = self._get_length
        
        print("new Lexer class initialised.")
    
    def roll_loop(self) -> None:
        for token in self.string:
            match token:
                case ':':
                    Lexer.create_new_token(TokenType.COMMA)
                case ';':
                    Lexer.create_new_token(TokenType.SEMICOLON)
                case '(':
                    Lexer.create_new_token(TokenType.LPAREN)
                case ')':
                    Lexer.create_new_token(TokenType.RPAREN)
                case '{':
                    Lexer.create_new_token(TokenType.LBRACE)
                case '}':
                    Lexer.create_new_token(TokenType.RBRACE)
                case '=':
                    Lexer.create_new_token(TokenType.ASSIGN)
                case '+':
                    Lexer.create_new_token(TokenType.PLUS)
                case '-':
                    Lexer.create_new_token(TokenType.MINUS)
                case '!':
                    Lexer.create_new_token(TokenType.BANG)
                case '#':
                    Lexer.create_new_token(TokenType.ASTERISK)
                case '/':
                    Lexer.create_new_token(TokenType.SLASH)
                case '<':
                    Lexer.create_new_token(TokenType.LT)
                case '>':
                    Lexer.create_new_token(TokenType.GT)
                case _:
                    Lexer.is_string(token)
        print("loop is done")
            
    @property
    def _get_length(self) -> int:
        return len(self.string)
    
    def _get_char(self):
        if Lexer.current_char_idx >= self.length:
            print("we're done here.")
            sys.exit(0)
        else:
            Lexer.current_char = self.string[Lexer.current_char_idx]
        Lexer.current_char_idx  += 1
        Lexer.next_char_idx     += 1
    
    @staticmethod
    def _skip_whitespace() -> None: 
        lcc = Lexer.current_char
        if lcc == " " or lcc == "\n" or lcc == "\t":
            None
            
    @staticmethod
    def initialise_lexer(*args) -> None:
        return Lexer(*args)
    
    @classmethod
    def create_new_token(type: TokenType):
        token_array.append(
            Token(type)
        )

    @staticmethod
    def is_string(token: str) -> bool:
        return True if token.isalpha() and ord(token) < 128 else False



string: str = """let num: int = 12"""

text: str = Lexer.initialise_lexer(string)
text.roll_loop()
for i in token_array:
    print(i)
