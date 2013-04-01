package main

import (
    "fmt"
    "log"
    "net/http"
    "text/template"
)

type data struct {
    Title string
    Css []string
    Script []string
}

func test() {
    fmt.Printf("hello Golog!\n")
}

func main() {
    // func routers
    for rule, funcname := range Urls {
        http.HandleFunc(rule, funcname)
    }
    // static files
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(settings["root"]+"/static"))))
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