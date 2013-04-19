/*
    data provider
*/
package main

import (
    //"fmt"
    "strings"
    "github.com/mikespook/mymysql/mysql"
    _ "github.com/mikespook/mymysql/native"
)

/* Model Struct*/
type Model struct {
    TableName string
    Db mysql.Conn
}

func (model *Model) fetchOne(data map[string]interface{}) {
    sql, params := mapSql(data)
    sql = "SELECT * FROM "+model.TableName+ " WHERE "
}

func (model *Model) fetchAll()  {
    
}

func (model *Model) deleteOne() {

}

func (model *Model) deleteAll() {

}

func (model *Model) update() {

}

func (model *Model) insert(data map[string]interface{}) uint64 {
    sql, params := mapSql(data)
    sql = "INSERT "+model.TableName+" SET " + sql
    stmt, err := model.Db.Prepare(sql)
    res, err := stmt.Run(params...)
    checkErr(err)
    return res.InsertId()
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
    db.Query("set names "+ settings["db_charset"])
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

/*
 name=?, time=?
*/
func mapInsert(data map[string]interface{}) (string, []interface{}) {
    var sql string
    params := make([]interface{}, len(data))
    i := 0
    for column, value := range(data) {
        params[i] = value
        sql += column + "=?, "
        i ++
    }
    sql = strings.Trim(sql, ", ")
    return sql, params
}

/*
 name =? AND time =?
*/

func mapWhere(data map[string]interface{}) (string, []interface{}) {
    sql := "true"
    params := make([]interface{}, len(data))
    i := 0
    for column, value := range(data) {
        params[i] = value
        sql += " AND "column + "=?"
        i ++
    }
    return sql, params
}