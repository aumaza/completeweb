package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
)

var db *sql.DB;

func connectionDB() * sql.DB {

    db, err := sql.Open("mysql", "root:slack142@tcp(slackzone.ddns.net:3306)/my_bills");

    if err != nil {
        fmt.Println("Connection to Database Failure...");
        panic(err);
    }

    return db;

}
