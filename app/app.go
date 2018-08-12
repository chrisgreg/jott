package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/chrisgreg/jott/config"
	"github.com/gorilla/mux"
)

type App struct {
	DB     *sql.DB
	Port   int
	Router *mux.Router
}

func (a *App) Initialise(config *config.Config) {
	//dsn := dbUser + ":" + dbPass + "@" + dbHost + "/" + dbName + "?charset=utf8"
	dbURI := fmt.Sprintf("%s:%s@%s/%s?charset=utf8",
		config.DB.User,
		config.DB.Pass,
		config.DB.Host,
		config.DB.DBName)

	db, err := sql.Open("mysql", dbURI)

	if err != nil {
		panic(err)
	}

	a.DB = db
	a.Port = config.Port

	a.Router = mux.NewRouter()
	a.setRoutes()
}

func (a *App) Run() {
	host := fmt.Sprintf(":%d", a.Port)
	log.Fatal(http.ListenAndServe(host, a.Router))
}

func (a *App) setRoutes() {
	a.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I'm alive")
	})
}

func (a *App) Get(route string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(route, f).Methods("GET")
}
