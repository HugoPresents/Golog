package main

import (
    "fmt"
    "log"
    "os"
    "net/http"
    "text/template"
    "encoding/json"
)

type layout_data struct {
    Title string
    SubTpl string
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
    http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir(settings["root"]+settings["static"]))))
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

func render(w http.ResponseWriter, tplName string) {
    templateDir := settings["root"]+settings["template"]
    t, err := template.ParseFiles(templateDir+"header.html", templateDir+tplName+".html", templateDir+"footer.html")
    if err != nil {
        fmt.Printf("\n#Template Dir: \n%s\n", templateDir)
        fmt.Printf("Error : %v\n", err)
    }
    t.ExecuteTemplate(w, "header", nil)
    t.ExecuteTemplate(w, tplName, nil)
    t.ExecuteTemplate(w, "footer", nil)
    t.Execute(w, nil)
}