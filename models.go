package main

type Category struct {
    Id int
    ParentId int
    OrderId int
    NavDisplay int
    Name string
    DisplayName string
    Status int
}

type Comment struct {
    Id int
    PostId int
    Author string
    AuthorEmail string
    AuthorUrl string
    Content string
    CreateTime int
    Status int
}

type Page struct {
    Id int
    Name string
    DisplayName string
    NavDisplay int
    Content string
    CreateTime int
    UpdateTime int
    OrderId int
    Status int
}

type Post struct {
    Id int
    CategoryId int
    Title string
    Content string
    CreateTime string
    UpdateTime string
    Status int
}