package main

import (
	"fmt"
	"net/http"
)

func main() {
	msg := []int{1, 3, 4, 5, 6, 9374983745, 4358598, 456578}
	msg = append(msg, 45)
	fmt.Println(len(msg))

	// functions
	add(5, 4)

	// Register route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Go!")
	})

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}

func add(a int, b int) int {

	sum := a + b
	return sum

}
