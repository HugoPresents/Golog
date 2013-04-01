package main

import (
    "net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
    t := render("index")
    d := data{Title:"Golog"}
    t.ExecuteTemplate(w, "layout", d)
    t.Execute(w, nil)
}