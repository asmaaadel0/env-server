package internal

import (
	"errors"
	"fmt"
	"net/http"
)

var ErrorOutOfRange = errors.New("port number out of range")

type App struct {
	port int
}

func NewApp(port int) (*App, error) {
	if port < 1 || port > 65535 {
		return nil, ErrorOutOfRange
	}
	return &App{port: port}, nil
}

func (app *App) Run() {

	http.HandleFunc("/env", handleRequests)
	http.HandleFunc("/env/", handleRequests)

	portListner := fmt.Sprintf(":%d", app.port)
	fmt.Println("Server started on port", portListner)
	http.ListenAndServe(portListner, nil)
}
