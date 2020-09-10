package main

import (
	"fmt"
	"log"

	"github.com/breise/rstack"
)

func main() {
	n := 20
	rs, err := fib(n)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	answer := rs.Join(`, `)
	fmt.Printf("The first %d Fibonacci numbers are: %v\n", n, answer)
}

func fib(n int) (*rstack.RStack, error) {
	if n < 1 {
		return nil, fmt.Errorf("fib(): value '%d' is out of range. Must be a positive integer", n)
	}
	if n == 1 {
		return rstack.New().Push(0), nil
	}
	if n == 2 {
		return rstack.NewFromSlice([]interface{}{0, 1}), nil
	}

	lastRStack, err := fib(n - 1)
	if err != nil {
		return nil, err
	}
	penultimateRStack, lastValue, err := lastRStack.Pop()
	if err != nil {
		return nil, err
	}
	_, penultimateValue, err := penultimateRStack.Pop()
	if err != nil {
		return nil, err
	}
	return lastRStack.Push(lastValue.(int) + penultimateValue.(int)), nil
}
