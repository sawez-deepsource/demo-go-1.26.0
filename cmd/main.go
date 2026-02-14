package main

import (
	"fmt"
	"math/big"
	"strings"
	"time"

	demo "go126demo"
)

func main() {
	fmt.Println("=== Go 1.26 Demo ===")
	fmt.Println()

	// new(expression)
	fmt.Println("-- new(expression) --")
	p := demo.NewInt(42)
	fmt.Printf("NewInt(42) = %d\n", *p)
	s := demo.NewString("hello 1.26")
	fmt.Printf("NewString(\"hello 1.26\") = %s\n", *s)
	person := demo.NewPersonDirect()
	fmt.Printf("NewPersonDirect() = %+v\n", *person)

	born := time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)
	data, err := demo.PersonJSON("Alice", born)
	if err != nil {
		fmt.Printf("PersonJSON error: %v\n", err)
	} else {
		fmt.Printf("PersonJSON = %s\n", data)
	}
	fmt.Println()

	// Self-referential generics
	fmt.Println("-- Self-referential Generics --")
	fmt.Println(demo.DemoGenerics())
	fmt.Println()

	// errors.AsType[T]
	fmt.Println("-- errors.AsType[T] --")
	nfErr := &demo.NotFoundError{Resource: "user/123"}
	fmt.Printf("HandleErrorNew(NotFound) = %s\n", demo.HandleErrorNew(nfErr))
	fmt.Printf("HandleErrorOld(NotFound) = %s\n", demo.HandleErrorOld(nfErr))

	permErr := &demo.PermissionError{User: "bob"}
	fmt.Printf("HandleErrorNew(Permission) = %s\n", demo.HandleErrorNew(permErr))
	fmt.Println()

	// bytes.Buffer.Peek
	fmt.Println("-- bytes.Buffer.Peek --")
	peeked, err := demo.PeekBuffer([]byte("Go 1.26 is here!"))
	if err != nil {
		fmt.Printf("PeekBuffer error: %v\n", err)
	} else {
		fmt.Printf("PeekBuffer first 4 bytes = %q\n", string(peeked))
	}
	fmt.Println()

	// io.ReadAll
	fmt.Println("-- io.ReadAll --")
	data, err = demo.ReadAllDemo(strings.NewReader("fast reader in 1.26"))
	if err != nil {
		fmt.Printf("ReadAllDemo error: %v\n", err)
	} else {
		fmt.Printf("ReadAllDemo = %s\n", data)
	}
	fmt.Println()

	// crypto/rand
	fmt.Println("-- crypto/rand --")
	randBytes, err := demo.GenerateRandomBytes(16)
	if err != nil {
		fmt.Printf("GenerateRandomBytes error: %v\n", err)
	} else {
		fmt.Printf("GenerateRandomBytes(16) = %x\n", randBytes)
	}

	randInt, err := demo.GenerateRandomBigInt(big.NewInt(1000))
	if err != nil {
		fmt.Printf("GenerateRandomBigInt error: %v\n", err)
	} else {
		fmt.Printf("GenerateRandomBigInt(max=1000) = %d\n", randInt)
	}
	fmt.Println()

	// Patterns
	fmt.Println("-- Misc patterns --")
	fmt.Printf("GoodStringCompare(\"GO\", \"go\") = %v\n", demo.GoodStringCompare("GO", "go"))
	fmt.Printf("AcceptAnythingModern(3.14) = %s\n", demo.AcceptAnythingModern(3.14))
	result, err := demo.NakedReturn(5, 10)
	fmt.Printf("NakedReturn(5, 10) = %d, err=%v\n", result, err)

	fmt.Println()
	fmt.Println("=== Done ===")
}
