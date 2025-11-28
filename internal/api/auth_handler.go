package api

import "net/http"

func Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Login"))
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from register"))
}
