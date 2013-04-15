package main

import (
    "fmt"
    "log"
    "os"
    "net/http"
    "text/template"
    "encoding/json"
)

type layoutData struct {
    Title string
    Css []string
    Script []string
}

var settings map[string]string

func main() {
    // load settings
    settings = loadSettings()
    // db driver
    //fmt.Printf("%s\n", settings["db_user"]+":"+settings["db_pass"]+"@tcp("+settings["db_host"]+")/"+settings["db_name"])
    /*
    var db_err error
    db, db_err = sql.Open("mysql", settings["db_user"]+":"+settings["db_pass"]+"@tcp("+settings["db_host"]+":"+settings["db_port"]+")/"+settings["db_name"])
    */
    // func routers
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

func renderSingle(w http.ResponseWriter, tplName string) {
    templateDir := settings["root"]+settings["template"]
    t, err := template.ParseFiles(templateDir+tplName+".html")
    if err != nil {
        fmt.Printf("\n#Template Dir: \n%s\n", templateDir)
        fmt.Printf("Error : %v\n", err)
    }
    t.ExecuteTemplate(w, tplName, nil)
}

func renderLayout(w http.ResponseWriter, tplName string, layout_data interface{}, content_data interface{}, layout ...string) {
    templateDir := settings["root"]+settings["template"]
    layoutName := "layout"
    if len(layout) > 0 {
        layoutName = layout[0]
    }
    t, err := template.ParseFiles(templateDir+layoutName+".html", templateDir+tplName+".html")
    if err != nil {
        fmt.Printf("\n#Template Dir: \n%s\n", templateDir)
        fmt.Printf("Error : %v\n", err)
    }
    //fmt.Printf("Layout name :%s\n", layoutName)
    t.ExecuteTemplate(w, layoutName, layout_data)
    t.ExecuteTemplate(w, tplName, content_data)
}