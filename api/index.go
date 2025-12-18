package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	initialized bool
	env         string
)

func init() {
	env = os.Getenv("TEST_ENV_VAR")
	if env == "" {
		log.Fatalf("Failed to load environment variable: TEST_ENV_VAR is not set")
	}
	initialized = true
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if !initialized {
		http.Error(w, "Not initialized", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "<h1>Hello from Go! Environment: %s</h1>", env)
}
