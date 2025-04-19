# Information necessary in order to write a compiler

## Pieces of a compiler
- Lexer
- Parser
- AST
- Evaluator


## Language example

```
int a = 3;
if (a <= 5) {
    print("a is small");
} else {
    print(a);
}
for (a <= 5) {
    a = a + 1;
}

func add(int x, int y) int {
    return x + y;
}
```

## Requirements of the language to build
### Data types
- integer
- float
- string
- boolean

### Structures
- arrays

### Flow control
- functions
- if/else statements
- loops

## How an interpreter works
Lexing, Parsing and building of the AST all happen at static time, meaning the programm is not running so far. It's just a translation from one representation to another. Then the evaluator takes over and produces the output.

### Lexer 
- Aka **tokenizer**. Turns an input (the code) into Tokens (meaningful bits of the langauge).
- Tokens are the names of things that appear in the programm and their corresponding values. (IDENTIFIER, EQUAL, SEMILCOLON, IF, FOR, PRINT, LBRACE, RBRACE...). Basically, *Type:Value* pairs.
- Everything becomes a token except for whitespace
- Tokenizer reads through the input character by character and groups them together into the tokens.
- The Lexer does **not** check for syntactically correct programms!

### Parser
- Takes tokens as an input and tries to figure out the structure of the original program. Conversion into an **Abstract Syntax Tree** (AST)
- Check for syntactically correct programm.

#### Grammar
- A grammar is a set of rules that is capable of defining an infinitely large language with an finite set of rules.
- The check whether a programm is syntactically correct boild down to answering wheter or not it can be derived from the languages grammar.

### AST
- Representation of the input programm that is suitable for a machine.

## Designing JK
### Grammar
#### Expressions
- A programm consists of several expressions. All possible (meaning syntactically correct expressions) are described by the following grammar:

```ebnf
expression ::= literal | unary | binary | grouping ;
literal    ::= NUMBER | STRING | "true" | "false" | "nil" ;
unary      ::= ( "-" | "!" ) expression ;
binary     ::= expression operator expression ;
grouping   ::= "(" expression ")" ;
operator   ::= "==" | "<=" | ">=" | "!=" | ">" | "<" | "+" | "-" | "*" | "/" ;
```

## Credits
- I'm using the book "Crafting Interpreters" as a guide: [Book link]{https://craftinginterpreters.com/}
