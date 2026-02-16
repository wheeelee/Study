package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Human struct {
	Name     string  `json:"name"`
	Address  string  `json:"address"`
	Age      int     `json:"age"`
	Married  bool    `json:"married"`
	H_height float64 `json:"h_height"`
}

func receive(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var sasha Human
	if err := json.Unmarshal(httpRequestBody, &sasha); err != nil {
		fmt.Println("err:", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Println(sasha.Name)
	fmt.Println(sasha.Address)
	fmt.Println(sasha.Age)
	fmt.Println(sasha.Married)
	fmt.Println(sasha.H_height)

}
func send(w http.ResponseWriter, r *http.Request) {
	sasha := Human{"Саша", "Москва", 21, false, 182.5}
	q, err := json.Marshal(sasha)
	if err != nil {
		fmt.Println("err", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	if _, err := w.Write(q); err != nil {
		fmt.Println("err", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
func main() {
	http.HandleFunc("/receive", receive)
	http.HandleFunc("/send", send)
	http.ListenAndServe(":9091", nil)
}
