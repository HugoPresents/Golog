package main

import (
    "net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
    t := render("index")
    d := layout_data{Title:"Golog"}
    t.ExecuteTemplate(w, "layout", d)
    t.Execute(w, nil)
}