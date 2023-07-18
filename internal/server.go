package internal

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
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

func handleRequests(w http.ResponseWriter, r *http.Request) {

	key := strings.TrimPrefix(r.URL.Path, "/env")

	if key == "" {
		handleEnv(w, r)
	} else {
		handleEnvKey(w, r)

	}
}

func handleEnv(w http.ResponseWriter, r *http.Request) {
	for _, env := range os.Environ() {
		fmt.Fprintln(w, env)
	}
}

func handleEnvKey(w http.ResponseWriter, r *http.Request) {
	key := strings.TrimPrefix(r.URL.Path, "/env/")
	value := os.Getenv(key)
	if value != "" {
		fmt.Fprintf(w, "%s=%s", key, value)
	} else {
		fmt.Fprintf(w, "Environment variable '%s' not found", key)
	}
}
