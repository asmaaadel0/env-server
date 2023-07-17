package internal

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func HandleEnv(w http.ResponseWriter, r *http.Request) {
	for _, env := range os.Environ() {
		fmt.Fprintln(w, env)
	}
}

func HandleEnvKey(w http.ResponseWriter, r *http.Request) {
	key := strings.TrimPrefix(r.URL.Path, "/env/")
	value := os.Getenv(key)
	if value != "" {
		fmt.Fprintf(w, "%s=%s", key, value)
	} else {
		fmt.Fprintf(w, "Environment variable '%s' not found", key)
	}
}
