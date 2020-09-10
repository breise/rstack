# rstack

RStack is a recursive stack, meaning that every node in the stack is itself an
RStack.  There is no special "stack" struct with a pointer to the first (or
last) "node" struct in the stack.  The RStack you are pointing to _is_ the last
node in the list, and points to another RStack node, which is the last node in
_its_ list.

This is convenient for passing a list (stack) of results into a function, which
pushes a result and passes the cumulative results to another function (maybe
recursively).  Each function sees its caller's result list (stack), without
having to manage a global stack at each invocation.

### Example

Basic RStack functions.

This example features...
- `New()`
- `NewFromSlice()`
- `Push()`
- `Pop()`
- `Join()`
- required type assertion when using a `Pop()`ed value

```
package main

import (
	"fmt"
	"log"

	"github.com/breise/rstack"
)

func main() {
	rs0 := rstack.New()
	rs1 := rs0.Push("one")
	rs2 := rs1.Push("two")
	rs3 := rs2.Push("three")
	fmt.Printf("rs0: %s\n", rs0.Join(`, `)) // rs0: 
	fmt.Printf("rs1: %s\n", rs1.Join(`, `)) // rs1: one
	fmt.Printf("rs2: %s\n", rs2.Join(`, `)) // rs2: one, two
	fmt.Printf("rs3: %s\n", rs3.Join(`, `)) // rs3: one, two, three

	rs0 = rstack.New()
	rs1 = rstack.New()
	rs2 = rstack.New()
	// but don't touch rs3
	fmt.Printf("rs0: %s\n", rs0.Join(`, `)) // rs0: 
	fmt.Printf("rs1: %s\n", rs1.Join(`, `)) // rs1: 
	fmt.Printf("rs2: %s\n", rs2.Join(`, `)) // rs2: 
	fmt.Printf("rs3: %s\n", rs3.Join(`, `)) // rs3: one, two, three

	rs2, v3, err := rs3.Pop()
	if err != nil {
		log.Fatal(err)
	}

	rs1, v2, err := rs2.Pop()
	if err != nil {
		log.Fatal(err)
	}

	rs0, v1, err := rs1.Pop()
	if err != nil {
		log.Fatal(err)
	}

	_, v0, err := rs0.Pop()
	if err != nil {
		fmt.Printf("rs0.Pop(): %s\n", err)                 // rs0.Pop(): Pop(): Cannot pop any empty RStack
	}

	fmt.Printf("v0: %-8v rs0: %s\n", v0, rs0.Join(`, `)) // v0: <nil>    rs0: 
	fmt.Printf("v1: %-8v rs1: %s\n", v1, rs1.Join(`, `)) // v1: one      rs1: one
	fmt.Printf("v2: %-8v rs2: %s\n", v2, rs2.Join(`, `)) // v2: two      rs2: one, two
	fmt.Printf("v3: %-8v rs3: %s\n", v3, rs3.Join(`, `)) // v3: three    rs3: one, two, three

	rs4 := rstack.NewFromSlice([]interface{}{"one", "two", "three", "four"})
	fmt.Printf("rs4: %s\n", rs4.Join(`, `))              // rs4: one, two, three, four
}
```

### A Whimsical Fibonacci Example

Demonstrates using `RStack` in a recursive function, in this case, as the result object.

This example features...
- `New()`
- `NewFromSlice()`
- `Push()`
- `Pop()`
- `Join()`
- required type assertion when using a `Pop()`ed value

```
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
```
