package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Test")
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":3000", nil)
}
