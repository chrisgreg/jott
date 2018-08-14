package main

import (
	"fmt"
	"net/http"

	"github.com/chrisgreg/jott/app"
	"github.com/chrisgreg/jott/config"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	config := config.GetConfig()

	app := &app.App{}
	app.Initialise(config)

	defer app.DB.Close()

	// err := app.DB.Ping()
	// if err != nil {
	// 	panic(err.Error())
	// }

	app.Run()
}

func startServer(port string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my app")
	})

	// http.HandleFunc("/newBlogPost", func(w http.ResponseWriter, r *http.Request) {
	// 	// TODO grab user id from JWT
	// 	// TODO grab title + subtitle from JSON post
	// 	// TODO respond in JSON with success or failure
	// 	_, err := dbUtil.CreateNewBlogPost(db, 1, "title", "subtitle")
	// 	if err != nil {
	// 		fmt.Println("Error adding new blog post", err.Error())
	// 	} else {
	// 		w.Write([]byte("Successfully added new blog post"))
	// 	}
	// })

	http.ListenAndServe(port, nil)
	fmt.Println("Server started on port " + port)
}
