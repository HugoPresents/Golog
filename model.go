package main

/* Model Struct*/
type Model struct {
    TableName string
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
    sql := "INSERT category SET "
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
    stmt, err := db.Prepare(sql)
    res, err := stmt.Exec(params...)
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
func newPostModel() *PostModel {
    return &PostModel{Model{"post"}}
}

func newCatModel() *CatModel {
    return &CatModel{Model{"category"}}
}

func newCommentModel() *CommentModel {
    return &CommentModel{Model{"comment"}}
}

func newPageModel() *PageModel {
    return &PageModel{Model{"page"}}
}