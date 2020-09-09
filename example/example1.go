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
	fmt.Printf("rs0: %s\n", rs0.Join(`, `))
	fmt.Printf("rs1: %s\n", rs1.Join(`, `))
	fmt.Printf("rs2: %s\n", rs2.Join(`, `))
	fmt.Printf("rs3: %s\n", rs3.Join(`, `))

	rs0 = rstack.New()
	rs1 = rstack.New()
	rs2 = rstack.New()
	fmt.Printf("rs0: %s\n", rs0.Join(`, `))
	fmt.Printf("rs1: %s\n", rs1.Join(`, `))
	fmt.Printf("rs2: %s\n", rs2.Join(`, `))
	fmt.Printf("rs3: %s\n", rs3.Join(`, `))

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
		fmt.Printf("rs0.Pop(): %s\n", err)
	}

	fmt.Printf("v0: %-8v rs0: %s\n", v0, rs0.Join(`, `))
	fmt.Printf("v1: %-8v rs1: %s\n", v1, rs1.Join(`, `))
	fmt.Printf("v2: %-8v rs2: %s\n", v2, rs2.Join(`, `))
	fmt.Printf("v3: %-8v rs3: %s\n", v3, rs3.Join(`, `))
}
