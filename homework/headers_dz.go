package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	for idd, heads := range r.Header {
		fmt.Println(idd, heads)
	}
}
func main() {
	http.HandleFunc("/default", handler)
	http.ListenAndServe(":9091", nil)
}
