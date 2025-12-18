package handler

import (
	"fmt"
	"net/http"
	"os"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	env := os.Getenv("TEST_ENV_VAR")
	fmt.Fprintf(w, "<h1>Hello from Go! Environment: %s</h1>", env)
}
