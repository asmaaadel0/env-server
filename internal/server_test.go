package internal

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHandleEnv(t *testing.T) {
	t.Run("Get request", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/env", nil)
		response := httptest.NewRecorder()

		handler := http.HandlerFunc(handleEnv)

		handler.ServeHTTP(response, request)

		if response.Code != http.StatusOK {
			t.Errorf("expected status %v, but got %v", http.StatusOK, response.Code)
		}
	})
}

func TestHandleEnvKey(t *testing.T) {
	t.Run("Get existing key", func(t *testing.T) {
		key := "SOME_VARIABLE"
		value := "some value"
		os.Setenv(key, value)
		defer os.Unsetenv(key)

		request, _ := http.NewRequest("GET", "/env/"+key, nil)

		response := httptest.NewRecorder()
		handler := http.HandlerFunc(handleEnvKey)

		handler.ServeHTTP(response, request)

		if response.Code != http.StatusOK {
			t.Errorf("expected status %v, but got %v", http.StatusOK, response.Code)
		}

		want := key + "=" + value
		if response.Body.String() != want {
			t.Errorf("expected body %v, but got %v", want, response.Body.String())
		}
	})

	t.Run("Get non existing key", func(t *testing.T) {
		key := "NON_EXISTENT_VARIABLE"

		request, err := http.NewRequest("GET", "/env/"+key, nil)
		if err != nil {
			t.Fatal(err)
		}

		respond := httptest.NewRecorder()
		handler := http.HandlerFunc(handleEnvKey)

		handler.ServeHTTP(respond, request)

		if respond.Code != http.StatusOK {
			t.Errorf("expected status %v, but got %v", http.StatusOK, respond.Code)
		}

		want := "Environment variable '" + key + "' not found"
		if respond.Body.String() != want {
			t.Errorf("expected body %v, but got %v", want, respond.Body.String())
		}
	})
}
