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
