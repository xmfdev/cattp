package main

import (
	"cattp/utils"
	"log"
	"net/http"
)

var SERVER_ADDRESS = ""
var SERVER_PORT = ""
var SERVER_FILE_SYSTEM_ROOT = ""
var SERVER_ENABLE_EXTRA_LOGS = false

func getServerAddress() string {
	return SERVER_ADDRESS + ":" + SERVER_PORT
}

func handle(w http.ResponseWriter, r *http.Request) {
	if SERVER_ENABLE_EXTRA_LOGS {
		utils.LogRequest(r)
	}

	// FileServer returns a `Handler` with the contents of a file system.
	http.FileServer(http.Dir(SERVER_FILE_SYSTEM_ROOT)).ServeHTTP(w, r)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")

	if SERVER_ENABLE_EXTRA_LOGS {
		utils.LogRequest(r)
	}

	log.Printf("Received form submission name=%s, e-mail=%s.", name, email)

	w.Write([]byte("Form successfully received!"))
}

func isValidSettings() bool {
	return SERVER_ADDRESS != "" && SERVER_PORT != "" && SERVER_FILE_SYSTEM_ROOT != ""
}

func startServer() {
	if !isValidSettings() {
		log.Fatal("Invalid settings, quitting.")
	}

	http.HandleFunc("/", handle)
	http.HandleFunc("/submit", postHandler)

	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/form.html", http.StatusMovedPermanently)
	})

	serverAddress := getServerAddress()
	log.Printf("Starting server on %s (root = %s).", serverAddress, SERVER_FILE_SYSTEM_ROOT)

	log.Fatal(http.ListenAndServe(getServerAddress(), nil))
}
