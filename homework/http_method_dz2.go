package main

import (
	"fmt"
	"io"
	"net/http"
	"slices"
	"strconv"
)

type Message struct {
	id  int
	msg string
}

var arr = make([]Message, 0)

func receive(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		HttpRequestBody, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Произошла ошибка чтения!", err)
		}
		HttpRequestBodyString := string(HttpRequestBody)
		if err != nil {
			fmt.Println("Не удалось расшифровать!", err)
		}
		arr = append(arr, Message{len(HttpRequestBodyString) + 1, HttpRequestBodyString})
	}
}
func send(w http.ResponseWriter, r *http.Request) {
	HttpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Произошла ошибка чтения!", err)
	}
	if _, err := w.Write([]byte(HttpRequestBody)); err != nil {
		fmt.Println("Произошла ошибка!", err)
	}
	HttpRequestBodyString := string(HttpRequestBody)
	arr = append(arr, Message{len(HttpRequestBodyString) + 1, HttpRequestBodyString})
}
func message_id(w http.ResponseWriter, r *http.Request) {
	HttpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Произошла ошибка чтения!", err)
	}
	HttpRequestBodyString := string(HttpRequestBody)
	for _, msg := range arr {
		if msg.msg == HttpRequestBodyString {
			w.Write([]byte(strconv.Itoa(msg.id)))
		}
	}
}
func message_list(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if len(arr) == 0 {
			w.Write([]byte("Нет сообщений"))
			return
		}
		var result string
		for i, msg := range arr {
			result += fmt.Sprintf("%d. ID: %d, Сообщение: %s\n", i, msg.id, msg.msg)
		}
		w.Write([]byte(result))
	}
}
func message_del(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Ошибка чтения:", err)
			http.Error(w, "Ошибка чтения запроса", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		idToDelete, err := strconv.Atoi(string(body))
		if err != nil {
			fmt.Println("Ошибка преобразования:", err)
			http.Error(w, "ID должен быть числом", http.StatusBadRequest)
			return
		}

		indexToDelete := -1
		for i, msg := range arr {
			if msg.id == idToDelete {
				indexToDelete = i
				break
			}
		}

		if indexToDelete != -1 {
			arr = slices.Delete(arr, indexToDelete, indexToDelete+1)
			fmt.Fprintf(w, "Сообщение с ID %d удалено", idToDelete)
		} else {
			http.Error(w, fmt.Sprintf("Сообщение с ID %d не найдено", idToDelete), http.StatusNotFound)
		}
	}
}
func main() {
	http.HandleFunc("/receive", receive)
	http.HandleFunc("/send", send)
	http.HandleFunc("/message_list", message_list)
	http.HandleFunc("/message_id", message_id)
	http.HandleFunc("/message_del", message_del)
	http.ListenAndServe(":9091", nil)
}
