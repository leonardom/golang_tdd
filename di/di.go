package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "Gopher")

	err := http.ListenAndServe(":3000", http.HandlerFunc(handleRequest))

	if err != nil {
		fmt.Println(err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Gopher from Web")
}
