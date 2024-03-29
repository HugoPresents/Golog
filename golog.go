package main

import (
    "fmt"
    "log"
    "os"
    "net/http"
    "text/template"
    "encoding/json"
    "github.com/astaxie/beedb"
    _ "github.com/ziutek/mymysql/godrv"
    "database/sql"
)

type layoutData struct {
    Title string
    Css []string
    Script []string
}
var orm beedb.Model
var settings map[string]string

func main() {
    // load settings
    settings = loadSettings()
    // func routers
    for rule, funcname := range Urls {
        http.HandleFunc(rule, funcname)
    }
    // static files
    http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir(settings["root"]+settings["static"]))))
    link := fmt.Sprintf("%s/%s/%s", settings["db_name"], settings["db_user"], settings["db_pass"])
    db, err := sql.Open("mymysql", link)
    if err != nil {
        panic(err)
    }
    orm = beedb.New(db)
    err = http.ListenAndServe(":8888", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

func loadSettings() map[string]string {
    var j map[string]string
    f, _ := os.Open("./settings.json")
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
    t.ExecuteTemplate(w, layoutName, layout_data)
    t.ExecuteTemplate(w, tplName, content_data)
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}