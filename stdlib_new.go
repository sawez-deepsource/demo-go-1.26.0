package go126demo

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// ============================================================
// Go 1.26 Feature: errors.AsType[T] - generic error unwrapping
// ============================================================

type NotFoundError struct {
	Resource string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s not found", e.Resource)
}

type PermissionError struct {
	User string
}

func (e *PermissionError) Error() string {
	return fmt.Sprintf("permission denied for user %s", e.User)
}

// GOOD: Using errors.AsType (new in Go 1.26) - type-safe and faster
func HandleErrorNew(err error) string {
	if nf, ok := errors.AsType[*NotFoundError](err); ok {
		return fmt.Sprintf("Not found: %s", nf.Resource)
	}
	if pe, ok := errors.AsType[*PermissionError](err); ok {
		return fmt.Sprintf("Permission denied: %s", pe.User)
	}
	return "unknown error"
}

// GOOD: Old-style errors.As still works (for comparison)
func HandleErrorOld(err error) string {
	var nf *NotFoundError
	if errors.As(err, &nf) {
		return fmt.Sprintf("Not found: %s", nf.Resource)
	}
	return "unknown"
}

// ============================================================
// Go 1.26 Feature: bytes.Buffer.Peek
// ============================================================

// GOOD: Using Buffer.Peek to look ahead without consuming
func PeekBuffer(data []byte) ([]byte, error) {
	buf := bytes.NewBuffer(data)
	peeked, err := buf.Peek(4)
	if err != nil {
		return nil, err
	}
	return peeked, nil
}

// ============================================================
// Go 1.26 Feature: Faster io.ReadAll
// ============================================================

// GOOD: io.ReadAll is now ~2x faster with less allocation
func ReadAllDemo(r io.Reader) ([]byte, error) {
	return io.ReadAll(r)
}

// BAD: Opening file without closing (resource leak, analyzer should catch)
func ReadFileBad(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	// BAD: missing f.Close() / defer f.Close()
	return io.ReadAll(f)
}

// GOOD: Proper file handling
func ReadFileGood(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return io.ReadAll(f)
}

// BAD: String comparison using == instead of strings.EqualFold for case-insensitive
func BadStringCompare(a, b string) bool {
	return strings.ToLower(a) == strings.ToLower(b)
}

// GOOD: Using strings.EqualFold
func GoodStringCompare(a, b string) bool {
	return strings.EqualFold(a, b)
}

// BAD: Empty error check branch
func BadErrorHandling(err error) {
	if err != nil {
		// empty branch - analyzer should flag this
	}
}

// BAD: Nil check after dereference (analyzer should catch)
func BadNilCheck(p *Person) string {
	name := p.Name
	if p == nil {
		return ""
	}
	return name
}
