package go126demo

import (
	"context"
	"fmt"
	"sync"
)

// ============================================================
// Common patterns that the analyzer should check on Go 1.26 code
// ============================================================

// BAD: Goroutine leak pattern (Go 1.26 has experimental goroutine leak profiler)
func LeakyGoroutine() error {
	ch := make(chan int)
	go func() {
		ch <- 42 // This goroutine leaks if nobody reads from ch
	}()
	return fmt.Errorf("returning without reading from channel")
}

// BAD: Context passed as non-first parameter (convention violation)
func BadContextParam(name string, ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		fmt.Println(name)
		return nil
	}
}

// GOOD: Context as first parameter
func GoodContextParam(ctx context.Context, name string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		fmt.Println(name)
		return nil
	}
}

// BAD: sync.Mutex copied (analyzer should catch)
type BadMutexCopy struct {
	mu sync.Mutex
}

func CopyMutex() {
	original := BadMutexCopy{}
	original.mu.Lock()
	copied := original // BAD: copies the mutex
	copied.mu.Unlock()
}

// BAD: Range over integer without using index (Go 1.22+ pattern, should work in 1.26)
func RangeOverInt() {
	for range 10 {
		fmt.Println("hello")
	}
}

// BAD: Defer in loop
func DeferInLoop(items []string) {
	for _, item := range items {
		defer fmt.Println(item) // Defers accumulate, don't execute until function returns
	}
}

// BAD: Boolean parameter (style check)
func ProcessItem(item string, verbose bool) {
	if verbose {
		fmt.Println("Processing:", item)
	}
	fmt.Println(item)
}

// BAD: Deeply nested if statements (cyclomatic complexity)
func DeeplyNested(a, b, c, d int) string {
	if a > 0 {
		if b > 0 {
			if c > 0 {
				if d > 0 {
					return "all positive"
				}
				return "d not positive"
			}
			return "c not positive"
		}
		return "b not positive"
	}
	return "a not positive"
}

// BAD: Empty interface{} instead of any (Go 1.18+ style)
func AcceptAnything(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

// GOOD: Using 'any' (modern Go style)
func AcceptAnythingModern(v any) string {
	return fmt.Sprintf("%v", v)
}

// BAD: Naked return in a complex function
func NakedReturn(x, y int) (result int, err error) {
	if x < 0 {
		err = fmt.Errorf("x must be non-negative")
		return // naked return
	}
	result = x + y
	return // naked return
}

// BAD: Using new(Type) where a literal would be clearer (pre-1.26 style)
// Note: In Go 1.26, new(expression) is the NEW feature, but new(Type) is old
func OldStyleNew() *int {
	p := new(int)
	*p = 42
	return p
}

// GOOD: Go 1.26 style - new(expression) is cleaner
func NewStyleNew() *int {
	return new(42)
}
