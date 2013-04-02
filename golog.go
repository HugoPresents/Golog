package main

import (
    "fmt"
    "log"
    "os"
    "net/http"
    "text/template"
    "encoding/json"
)

type data struct {
    Title string
    Css []string
    Script []string
}

var settings map[string]string

func main() {
    // func routers
    settings = loadSettings()
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

func loadSettings() map[string]string {
    var j map[string]string
    f, _ := os.Open("./settings.json")
    //settings := string("")
    buf := make([]byte, 1024)
    n, _ := f.Read(buf)
    json.Unmarshal(buf[:n], &j)
    return j
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