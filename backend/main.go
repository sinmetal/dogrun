package backend

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/helloworld", handle)
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}