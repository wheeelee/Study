package main

import (
	"fmt"
	"study2/Library"
	"study2/http"
)

func main() {
	ListofBooks := Library.NewList()
	httpHandlers := http.NewHTTPHandlers(ListofBooks)
	httpServer := http.NewHTTPServer(httpHandlers)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed to start http server:", err)
	}
}
