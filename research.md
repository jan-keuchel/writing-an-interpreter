# Information necessary in order to write an compiler

## Pieces of a compiler
- Lexer
- Parser
- AST
- Bytecode-compiler
- Bytecode interpreter
- Runtime


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

### Requirements of language
#### Data types
- integer
- float
- string
- boolean

#### Structures
- arrays

#### Flow control
- functions
- if/else statements
- loops

## How the compiler works
### Lexer 
- Aka tokenizer. Turns an input into Tokens. Tokens are the names of things that appear in the programm and their corresponding values. (NAME, EQUAL, SEMILCOLON, IF, FOR, PRINT, LBRACE, RBRACE...)
- Everything becomes a token except for whitespace
- Tokenizer has a *next* function that spits out the next Token it read.

### Parser
- Takes tokens as a input and tries to figure out the structure of the original program. Conversion into an **Abstract Syntax Tree** (AST)
- Has building *blocks* (If, Assignment, Comparison, Equals, Number) and chains them together, thus creating the structure of the programm.

