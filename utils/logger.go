package utils

import (
	"log"
	"net/http"
)

const (
	green = "\033[32m"
	reset = "\033[0m"
)

func LogRequest(r *http.Request) {
	log.Printf("Received %s%s%s (%s) from %s.", green, r.Method, reset, r.URL, r.RemoteAddr)
}
