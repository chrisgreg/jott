package main

import (
	"database/sql"
	"fmt"
	"net/http"

	dbUtil "github.com/chrisgreg/jott/db"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

const (
	dbHost = "tcp(db:3306)"
	dbName = "jott"
	dbUser = "root"
	dbPass = "root"
)

func main() {

	dsn := dbUser + ":" + dbPass + "@" + dbHost + "/" + dbName + "?charset=utf8"
	var err error

	db, err = sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(db)

	startServer(":3001")
}

func startServer(port string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my app")
	})

	http.HandleFunc("/newBlogPost", func(w http.ResponseWriter, r *http.Request) {
		// TODO grab user id from JWT
		// TODO grab title + subtitle from JSON post
		// TODO respond in JSON with success or failure
		_, err := dbUtil.CreateNewBlogPost(db, 1, "title", "subtitle")
		if err != nil {
			fmt.Println("Error adding new blog post", err.Error())
		} else {
			w.Write([]byte("Successfully added new blog post"))
		}
	})

	http.ListenAndServe(port, nil)
	fmt.Println("Server started on port " + port)
}
