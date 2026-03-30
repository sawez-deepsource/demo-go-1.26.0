package go126demo

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
)

// Hardcoded credentials
const dbPassword = "admin123"
const apiKey = "sk-secret-key-do-not-share"

func connectDB() *sql.DB {
	db, _ := sql.Open("postgres", "postgres://admin:"+dbPassword+"@localhost/mydb")
	return db
}

// SQL injection
func getUser(db *sql.DB, userID string) string {
	query := "SELECT name FROM users WHERE id = '" + userID + "'"
	row := db.QueryRow(query)
	var name string
	row.Scan(&name)
	return name
}

// Resource leak - file never closed
func readConfig(path string) string {
	f, err := os.Open(path)
	if err != nil {
		return ""
	}
	buf := make([]byte, 1024)
	n, _ := f.Read(buf)
	return string(buf[:n])
}

// Mutex copied by value
func processWith(mu sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Println("processing")
}

// Inefficient string concatenation in loop
func buildString(items []string) string {
	result := ""
	for _, item := range items {
		result += item + ","
	}
	return result
}

// Error ignored
func writeFile(path string, data []byte) {
	f, _ := os.Create(path)
	f.Write(data)
	f.Close()
}

// Naked return with named results
func divide(a, b float64) (result float64, err error) {
	if b == 0 {
		err = fmt.Errorf("division by zero")
		return
	}
	result = a / b
	return
}

// Boolean parameter - unclear at call site
func fetchData(url string, useCache bool, retryOnError bool) ([]byte, error) {
	if useCache {
		return nil, nil
	}
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	return nil, nil
}

// Deeply nested code
func processRequest(r *http.Request) string {
	if r != nil {
		if r.Method == "POST" {
			if r.Body != nil {
				if r.ContentLength > 0 {
					if r.ContentLength < 1000000 {
						return "ok"
					}
				}
			}
		}
	}
	return "fail"
}

// Goroutine leak - channel never read
func leakyGoroutine() {
	ch := make(chan int)
	go func() {
		ch <- 42
	}()
}

// Unused parameter
func calculate(x int, y int, z int) int {
	return x + y
}

// Defer in loop
func closeFiles(paths []string) {
	for _, p := range paths {
		f, err := os.Open(p)
		if err != nil {
			continue
		}
		defer f.Close()
		buf := make([]byte, 100)
		f.Read(buf)
	}
}

// Type assertion without ok check
func getString(v interface{}) string {
	return v.(string)
}

// Unreachable code
func earlyReturn(x int) int {
	return x * 2
	fmt.Println("this never runs")
	return x
}

// Inefficient use of fmt.Sprintf for simple conversion
func intToString(n int) string {
	return fmt.Sprintf("%d", n)
}

// Should use strconv.Itoa instead
var _ = strconv.Itoa
