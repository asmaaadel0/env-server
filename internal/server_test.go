package internal

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHandleEnv(t *testing.T) {
	t.Run("Get request", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/env", nil)
		response := httptest.NewRecorder()

		handler := http.HandlerFunc(handleRequests)

		handler.ServeHTTP(response, request)

		if response.Code != http.StatusOK {
			t.Errorf("expected status %v, but got %v", http.StatusOK, response.Code)
		}
	})
}

func TestHandleEnvKey(t *testing.T) {
	t.Run("Get empty key", func(t *testing.T) {
		request, err := http.NewRequest("GET", "/env/", nil)
		if err != nil {
			t.Fatal(err)
		}

		respond := httptest.NewRecorder()
		handler := http.HandlerFunc(handleRequests)

		handler.ServeHTTP(respond, request)

		if respond.Code != http.StatusNotFound {
			t.Errorf("expected status %v, but got %v", http.StatusNotFound, respond.Code)
		}
	})

	t.Run("Get existing key", func(t *testing.T) {
		key := "SOME_VARIABLE"
		value := "value"

		os.Setenv(key, value)
		defer os.Unsetenv(key)

		request, _ := http.NewRequest("GET", "/env/"+key, nil)

		response := httptest.NewRecorder()
		handler := http.HandlerFunc(handleRequests)

		handler.ServeHTTP(response, request)

		if response.Code != http.StatusOK {
			t.Errorf("expected status %v, but got %v", http.StatusOK, response.Code)
		}

		want := value
		var got string
		json.NewDecoder(response.Body).Decode(&got)
		if got != want {
			t.Errorf("expected body %v, but got %v", want, got)
		}
	})

	t.Run("Get non existing key", func(t *testing.T) {
		key := "NON_EXISTENT_VARIABLE"

		request, err := http.NewRequest("GET", "/env/"+key, nil)
		if err != nil {
			t.Fatal(err)
		}

		respond := httptest.NewRecorder()
		handler := http.HandlerFunc(handleRequests)

		handler.ServeHTTP(respond, request)

		if respond.Code != http.StatusNotFound {
			t.Errorf("expected status %v, but got %v", http.StatusNotFound, respond.Code)
		}

		want := ""

		var got string
		json.NewDecoder(respond.Body).Decode(&got)

		if got != want {
			t.Errorf("expected body %v, but got %v", want, got)
		}
	})
}

func TestHandleRequests(t *testing.T) {
	t.Run("Post request", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodPost, "/env", nil)
		if err != nil {
			t.Fatal(err)
		}

		respond := httptest.NewRecorder()
		handler := http.HandlerFunc(handleRequests)

		handler.ServeHTTP(respond, request)
	})
}

func TestNotFound(t *testing.T) {
	t.Run("Run wrong path", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/notFound", nil)
		response := httptest.NewRecorder()

		handler := http.HandlerFunc(handleRequests)

		handler.ServeHTTP(response, request)

		if response.Code != http.StatusNotFound {
			t.Errorf("expected status %v, but got %v", http.StatusNotFound, response.Code)
		}
	})
}
