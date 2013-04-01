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

func render(tplName string) *template.Template{
    templateDir := settings["root"]+settings["template"]
    t, err := template.ParseFiles(templateDir+"layout.html", templateDir+tplName+".html")
    if err != nil {
        fmt.Printf("\n#Template Dir: \n%s\n", templateDir)
        fmt.Printf("Error : %v\n", err)
    }
    return t
}