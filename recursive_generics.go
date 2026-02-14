package go126demo

import "fmt"

// ============================================================
// Go 1.26 Feature: Self-Referential Generic Type Constraints
// ============================================================

// GOOD: Self-referential generic constraint - A type that references itself
// in its own type parameter list. This was not allowed before Go 1.26.
type Adder[A Adder[A]] interface {
	Add(A) A
}

// GOOD: Concrete type implementing the self-referential Adder interface
type Vec2D struct {
	X, Y float64
}

func (v Vec2D) Add(other Vec2D) Vec2D {
	return Vec2D{X: v.X + other.X, Y: v.Y + other.Y}
}

// GOOD: Generic function using self-referential constraint
func Sum[A Adder[A]](items ...A) A {
	var result A
	for _, item := range items {
		result = result.Add(item)
	}
	return result
}

// GOOD: Another self-referential constraint - Comparable type
type Comparable[T Comparable[T]] interface {
	CompareTo(T) int
}

type SortedInt int

func (s SortedInt) CompareTo(other SortedInt) int {
	if s < other {
		return -1
	}
	if s > other {
		return 1
	}
	return 0
}

// GOOD: Self-referential constraint for a builder pattern
type Builder[B Builder[B]] interface {
	WithName(string) B
	Build() string
}

type HTMLBuilder struct {
	name string
}

func (h HTMLBuilder) WithName(name string) HTMLBuilder {
	h.name = name
	return h
}

func (h HTMLBuilder) Build() string {
	return fmt.Sprintf("<div>%s</div>", h.name)
}

// GOOD: Using the generic function with concrete types
func DemoGenerics() string {
	v1 := Vec2D{1, 2}
	v2 := Vec2D{3, 4}
	v3 := Vec2D{5, 6}
	result := Sum(v1, v2, v3)
	return fmt.Sprintf("Sum: (%f, %f)", result.X, result.Y)
}

// BAD: Unreachable code after return (analyzer should catch)
func BadGenericUsage() string {
	v := Vec2D{1, 2}
	result := v.Add(Vec2D{3, 4})
	return fmt.Sprintf("(%f, %f)", result.X, result.Y)
	fmt.Println("unreachable") // unreachable code
	return ""
}
