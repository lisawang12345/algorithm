package main

import (
    "database/sql"
    "fmt"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", "root:Zhangyq1234#@tcp(127.0.0.1:3306)/oa?charset=utf8")
    if err != nil {
        fmt.Printf("connect mysql fail ! [%s]", err)
    } else {
        fmt.Println("connect to mysql success")
    }
    rows, err := db.Query("select id, username from user_info")
    if err != nil {
        fmt.Printf("select fail [%s]", err)
    }
    var mapUser map[string]int
    mapUser = make(map[string]int)
    for rows.Next() {
        var id int
        var username string
        rows.Columns()
        err := rows.Scan(&id, &username)
        if err != nil {
            fmt.Printf("get user info error [%s]", err)
        }
        mapUser[username] = id
    }
    for k, v := range mapUser {
        fmt.Println(k, v)
    }
}
