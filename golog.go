package main

import (
    "fmt"
    "log"
    "net/http"
    "text/template"
)

func test() {
    fmt.Printf("hello Golog!\n")
}

func main() {
    for rule, funcname := range Urls {
        http.HandleFunc(rule, funcname)
    }
    err := http.ListenAndServe(":8888", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

func render(tplName string, w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("templates/layout.html", "templates/"+tplName+".html")
    t.ExecuteTemplate(w, "layout", nil)
    t.Execute(w, nil)
}