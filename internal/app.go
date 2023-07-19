package internal

import (
	"errors"
	"fmt"
	"net/http"
)

// ErrorOutOfRange if user enter invalid port
var ErrorOutOfRange = errors.New("port number out of range, range should be between [1, 65535]")

type App struct {
	port int
}

func NewApp(port int) (*App, error) {
	if port < 1 || port > 65535 {
		return nil, ErrorOutOfRange
	}
	return &App{port: port}, nil
}

func (app *App) Run() error {

	http.HandleFunc("/env", handleRequests)
	http.HandleFunc("/env/", handleRequests)

	portListner := fmt.Sprintf(":%d", app.port)
	fmt.Println("Server started on port", portListner)
	err := http.ListenAndServe(portListner, nil)
	return err
}
