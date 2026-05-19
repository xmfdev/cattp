package main

import (
	"log"
	"net/http"
)

var SERVER_ADDRESS = ""
var SERVER_PORT = ""
var SERVER_FILE_SYSTEM_ROOT = ""

func getServerAddress() string {
	return SERVER_ADDRESS + ":" + SERVER_PORT
}

func handle(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s (%s) from %s.", r.Method, r.URL, r.RemoteAddr)

	// FileServer returns a `Handler` with the contents of a file system.
	http.FileServer(http.Dir(SERVER_FILE_SYSTEM_ROOT)).ServeHTTP(w, r)
}

func isValidSettings() bool {
	return SERVER_ADDRESS != "" && SERVER_PORT != "" && SERVER_FILE_SYSTEM_ROOT != ""
}

func startServer() {
	if !isValidSettings() {
		log.Fatal("Invalid settings, quitting.")
	}

	http.HandleFunc("/", handle)

	serverAddress := getServerAddress()
	log.Printf("Starting server on %s (root = %s).", serverAddress, SERVER_FILE_SYSTEM_ROOT)

	log.Fatal(http.ListenAndServe(getServerAddress(), nil))
}
