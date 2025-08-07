# Interpreter

> The interpreter pattern is a design pattern that specifies how to evaluate sentences in a language. The basic idea is to have a class for each symbol (terminal or nonterminal) in a specialized computer language. The syntax tree of a sentence in the language is an instance of the composite pattern and is used to evaluate (interpret) the sentence for a client. (Wikipedia)

> The Interpreter Design Pattern is a behavioral design pattern used to define a language's grammar and provide an interpreter to process statements in that language. It is useful for parsing and executing expressions or commands in a system. By breaking down complex problems into smaller, understandable pieces, this pattern simplifies the interpretation of structured data. (GeeksForGeeks)

## Example

### Polish

```go
package main

import (
 "fmt"
 "strconv"
 "strings"
)

type Interpreter interface {
 Read() int
}

type value int

func (v *value) Read() int {
 return int(*v)
}

type operationSum struct {
 Left  Interpreter
 Right Interpreter
}

func (a *operationSum) Read() int {
 return a.Left.Read() + a.Right.Read()
}

type operationSub struct {
 Left  Interpreter
 Right Interpreter
}

func (s *operationSub) Read() int {
 return s.Left.Read() - s.Right.Read()
}

func operatorFactory(o string, left, right Interpreter) Interpreter {
 switch o {
 case SUM:
  return &operationSum{
   Left:  left,
   Right: right,
  }
 case SUB:
  return &operationSub{
   Left:  left,
   Right: right,
  }
 }
 return nil
}

const (
 SUM = "sum"
 SUB = "sub"
 MUL = "mul"
 DIV = "div"
)

type polishNotationStack []Interpreter

func (p *polishNotationStack) Push(s Interpreter) {
 *p = append(*p, s)
}

func (p *polishNotationStack) Pop() Interpreter {
 length := len(*p)

 if length > 0 {
  temp := (*p)[length-1]
  *p = (*p)[:length-1]
  return temp
 }

 return nil
}

func Calculate(o string) (int, error) {
 stack := polishNotationStack{}
 operators := strings.Split(o, " ")

 for _, operatorString := range operators {
  if operatorString == SUM || operatorString == SUB {
   right := stack.Pop()
   left := stack.Pop()
   mathFunc := operatorFactory(operatorString, left, right)
   res := value(mathFunc.Read())
   stack.Push(&res)
  } else {
   val, err := strconv.Atoi(operatorString)
   if err != nil {
    panic(err)
   }
   temp := value(val)
   stack.Push(&temp)
  }
 }
 return int(stack.Pop().Read()), nil
}

func main() {
 operation := "5 3 sum 2 sub"
 res, err := Calculate(operation)
 if err != nil {
  panic(err)
 }
 fmt.Println(res)
}
```

### Lexing and Parse

```go
package main

import (
 "fmt"
 "strconv"
 "strings"
 "unicode"
)

type Element interface {
 Value() int
}

type Integer struct {
 value int
}

func NewInteger(value int) *Integer {
 return &Integer{value: value}
}

func (i *Integer) Value() int {
 return i.value
}

type Operation int

const (
 Addition Operation = iota
 Subtraction
)

type BinaryOperation struct {
 Type        Operation
 Left, Right Element
}

func (b *BinaryOperation) Value() int {
 switch b.Type {
 case Addition:
  return b.Left.Value() + b.Right.Value()
 case Subtraction:
  return b.Left.Value() - b.Right.Value()
 default:
  panic("Unsuported operation")
 }
}

type TokenType int

const (
 Int TokenType = iota
 Plus
 Minus
 Lparen
 Rparen
)

type Token struct {
 Type TokenType
 Text string
}

func (t *Token) String() string {
 return fmt.Sprintf("`%s`", t.Text)
}

func Lex(input string) []Token {
 var result []Token

 for i := 0; i < len(input); i++ {
  switch input[i] {
  case '+':
   result = append(result, Token{Plus, "+"})
  case '-':
   result = append(result, Token{Minus, "-"})
  case '(':
   result = append(result, Token{Lparen, "("})
  case ')':
   result = append(result, Token{Rparen, ")"})
  default:
   sb := strings.Builder{}
   for j := i; j < len(input); j++ {
    if unicode.IsDigit(rune(input[j])) {
     sb.WriteRune(rune(input[j]))
     i++
    } else {
     result = append(result, Token{
      Int, sb.String()})
     i--
     break
    }
   }
  }
 }
 return result
}

func Parse(tokens []Token) Element {
 result := BinaryOperation{}
 haveLhs := false
 for i := 0; i < len(tokens); i++ {
  token := &tokens[i]
  switch token.Type {
  case Int:
   n, _ := strconv.Atoi(token.Text)
   integer := Integer{n}
   if !haveLhs {
    result.Left = &integer
    haveLhs = true
   } else {
    result.Right = &integer
   }
  case Plus:
   result.Type = Addition
  case Minus:
   result.Type = Subtraction
  case Lparen:
   j := i
   for ; j < len(tokens); j++ {
    if tokens[j].Type == Rparen {
     break
    }
   }
   var subexp []Token
   for k := i + 1; k < j; k++ {
    subexp = append(subexp, tokens[k])
   }
   element := Parse(subexp)
   if !haveLhs {
    result.Left = element
    haveLhs = true
   } else {
    result.Right = element
   }
   i = j
  }
 }
 return &result
}

func main() {
 input := "(13+4)-(12+1)"
 tokens := Lex(input)
 fmt.Println(tokens)

 parsed := Parse(tokens)
 fmt.Printf("%s = %d\n", input, parsed.Value())
}
```
