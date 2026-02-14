package go126demo

import (
	"encoding/json"
	"time"
)

// ============================================================
// Go 1.26 Feature: new(expression) - creates pointer to value
// ============================================================

type Person struct {
	Name string `json:"name"`
	Age  *int   `json:"age"`
}

// GOOD: Using new(expression) to create a pointer to a computed value.
// The analyzer should NOT flag this as an issue.
func PersonJSON(name string, born time.Time) ([]byte, error) {
	return json.Marshal(Person{
		Name: name,
		Age:  new(yearsSince(born)),
	})
}

func yearsSince(t time.Time) int {
	return int(time.Since(t).Hours() / 8760)
}

// GOOD: new(expression) with a literal value
func NewInt(v int) *int {
	return new(v)
}

// GOOD: new(expression) with a string
func NewString(s string) *string {
	return new(s)
}

// GOOD: new(expression) with a struct literal
func NewPersonDirect() *Person {
	return new(Person{Name: "Alice", Age: new(30)})
}

// BAD: Unused parameter (analyzer should catch GO-W1029 or similar)
func NewIntIgnored(v int, unused string) *int {
	return new(v)
}

// BAD: Error return value ignored (analyzer should catch)
func PersonJSONBad(name string, born time.Time) []byte {
	data, _ := json.Marshal(Person{
		Name: name,
		Age:  new(yearsSince(born)),
	})
	return data
}
