package internal

import (
	"encoding/json"
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
	encoder := json.NewEncoder(w)

	envMap := make(map[string]string)
	for _, env := range os.Environ() {
		splited := strings.SplitN(env, "=", 2)
		envMap[splited[0]] = splited[1]
	}

	err := encoder.Encode(envMap)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func handleEnvKey(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	key := strings.TrimPrefix(r.URL.Path, "/env/")
	value := os.Getenv(key)

	if value == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err := encoder.Encode(value)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
