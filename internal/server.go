package internal

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func handleRequests(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		return
	}
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
