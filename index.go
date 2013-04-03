package main

import (
    "net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
    render(w, "index")
}

func login(w http.ResponseWriter, r *http.Request) {
    render(w, "login")
}