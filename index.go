package main

import (
    "net/http"
    "fmt"
)

func index(w http.ResponseWriter, r *http.Request) {
    t := render("index")
    d := data{Title:"Golog"}
    t.ExecuteTemplate(w, "layout", d)
    t.Execute(w, nil)
}

func static(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("%v\n", r.URL.Path)
    file := settings["root"]+settings["static"] + r.URL.Path[len(settings["static"]):]
    http.ServeFile(w, r, file)
    return
}