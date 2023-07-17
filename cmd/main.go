package main
import (
	"fmt"
	"net/http"

	"github.com/codescalersinternships/envserver-Asmaa/internal"
)
func main() {
	http.HandleFunc("/env", internal.HandleEnv)
	http.HandleFunc("/env/", internal.HandleEnvKey)

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}