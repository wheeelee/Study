package http

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	httpHandlers *HTTPHandlers
}

func NewHTTPServer(httpHandler *HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		httpHandlers: httpHandler,
	}
}

func (s *HTTPServer) StartServer() error {
	router := mux.NewRouter()

	router.Path("/books").Methods("POST").HandlerFunc(s.httpHandlers.HandleCreateBook)
	router.Path("/books/{title}").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetBook)

	router.Path("/books").Methods("GET").Queries("Сompleted", "false").HandlerFunc(s.httpHandlers.HandleGetAllUnreadBooks)

	router.Path("/books").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetAllBook)
	router.Path("/books/{title}").Methods("PATCH").HandlerFunc(s.httpHandlers.HandleCompleteBook)
	router.Path("/books/{title}").Methods("DELETE").HandlerFunc(s.httpHandlers.HandleDeleteTask)

	if err := http.ListenAndServe(":9091", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		return err
	}

	return nil
}
