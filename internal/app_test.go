package internal

import (
	"errors"
	"testing"
)

func TestNewApp(t *testing.T) {
	t.Run("fail to create app", func(t *testing.T) {
		port := 1
		_, err := NewApp(port)
		if errors.Is(err, ErrorOutOfRange) {
			t.Errorf("expected error in creating app")
		}

	})
	t.Run("create app sucessfully", func(t *testing.T) {
		port := 8080
		_, err := NewApp(port)
		if err != nil {
			t.Errorf("expected creating app")
		}

	})
}
