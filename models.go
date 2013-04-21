package main

type Category struct {
    Id int
    Parent_id int
    Order_id int
    Nav_display int
    Name string
    Display_name string
    Status int
}

type Comment struct {
    Id int
    Post_id int
    Author string
    Author_email string
    Author_url string
    Content string
    Create_time int
    Status int
}

type Page struct {
    Id int
    Name string
    Display_name string
    Nav_display int
    Content string
    Create_time int
    Update_time int
    Order_id int
    Status int
}

type Post struct {
    Id int
    Category_id int
    Title string
    Content string
    Create_time string
    Update_time string
    Status int
}