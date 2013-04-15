/*
    data provider
*/
package main

import (
    "fmt"
    "strings"
    "github.com/mikespook/mymysql/mysql"
    _ "github.com/mikespook/mymysql/native"
)

/* Model Struct*/
type Model struct {
    TableName string
    Db mysql.Conn
}

func (model *Model) fetchOne() {

}

func (model *Model) fetchAll()  {

}

func (model *Model) deleteOne() {

}

func (model *Model) deleteAll() {

}

func (model *Model) update() {

}

func (model *Model) insert(data map[string]string) {
    
    sql := "INSERT "+model.TableName+" SET "
    if len(data) < 1 {
        fmt.Print("no data!")
        return
    }
    params := make([]interface{}, len(data))
    i := 0
    for column, value := range(data) {
        params[i] = value
        sql += column + "=?, "
        i ++
    }
    sql = strings.Trim(sql, ", ")
    stmt, err := model.Db.Prepare(sql)
    res, err := stmt.Run(params...)
    checkErr(err)
    fmt.Println("%v", res)
}
/* Model struct end */

/*  Category Model */
type CatModel struct {
    Model
}

/* Category Model end */

/* Post Model */
type PostModel struct {
    Model
}

/* Post Model end */

/* Comment Model */
type CommentModel struct {
    Model
}

/* Comment Model end */

/* Page Model */
type PageModel struct {
    Model
}

/* Page Model end */

/* model generator */

func newModel() Model {
    db := mysql.New("tcp", "", settings["db_host"]+":"+settings["db_port"], settings["db_user"], settings["db_pass"], settings["db_name"])
    err := db.Connect()
    if err != nil {
        panic(err)
    }
    return Model{Db:db}
}

func newPostModel() *PostModel {
    model := newModel()
    model.TableName = "post"
    return &PostModel{model}
}

func newCatModel() *CatModel {
    model := newModel()
    model.TableName = "category"
    return &CatModel{model}
}

func newCommentModel() *CommentModel {
    model := newModel()
    model.TableName = "comment"
    return &CommentModel{model}
}

func newPageModel() *PageModel {
    model := newModel()
    model.TableName = "page"
    return &PageModel{model}
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}