package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strconv"
)

type Message struct {
	Id         int    `json:"id"`
	Msg        string `json:"msg"`
	Importance bool   `json:"importance"`
}

var arr = make([]Message, 0)

func receive(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var msg Message
	if err := json.Unmarshal(httpRequestBody, &msg); err != nil {
		fmt.Println("err:", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	arr = append(arr, msg)
}
func send(w http.ResponseWriter, r *http.Request) {
	var msg = Message{len("Привет!") + 1, "Привет!", true}
	q, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("err", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	if _, err := w.Write(q); err != nil {
		fmt.Println("err", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
func message_id(w http.ResponseWriter, r *http.Request) {
	HttpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Произошла ошибка чтения!", err)
	}
	HttpRequestBodyString := string(HttpRequestBody)
	for _, msg := range arr {
		if msg.Msg == HttpRequestBodyString {
			w.Write([]byte(strconv.Itoa(msg.Id)))
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
			result += fmt.Sprintf("%d. ID: %d, Сообщение: %s\n", i, msg.Id, msg.Msg)
		}
		w.Write([]byte(result))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
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
			if msg.Id == idToDelete {
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
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
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
