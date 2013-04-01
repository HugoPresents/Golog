package main

import (
    "net/http"
    //"text/template"
    //"fmt"
)

func index(w http.ResponseWriter, r *http.Request) {
    //fmt.Fprintf(w, "Hello Golog<br>")
    /*
    t, _ := template.ParseFiles("templates/layout.html", "templates/index.html")
    t.ExecuteTemplate(w, "layout", nil)
    t.Execute(w, nil)
    */
    //render("index", w, r)
    t := render("index")
    t.ExecuteTemplate(w, "layout", nil)
    t.Execute(w, nil)
}