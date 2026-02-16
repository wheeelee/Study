package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Произошла ошибка чтения!", err)
	}
	httpRequestBodyString := string(httpRequestBody)
	msg, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		fmt.Println("Произошла ошибка перевода!", err)
	}
	if msg == 1 {
		http.Error(w, "Вы ввели код: 1!", http.StatusBadRequest)
	}
	if msg == 2 {
		http.Error(w, "Вы ввели код: 2!", http.StatusContinue)
	}
}
func main() {
	http.HandleFunc("/default", handler)
	http.ListenAndServe(":9091", nil)
}
