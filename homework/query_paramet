package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	firstparam := r.URL.Query().Get("first")
	secondparam := r.URL.Query().Get("second")
	fmt.Println("firstparam: ", firstparam)
	fmt.Println("secondparam: ", secondparam)
}
func main() {
	http.HandleFunc("/receive_params", handler)
	http.ListenAndServe(":9091", nil)
}
