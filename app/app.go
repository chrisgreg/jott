package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chrisgreg/jott/app/handler"
	"github.com/chrisgreg/jott/app/models"
	"github.com/chrisgreg/jott/config"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type App struct {
	// DB     *sql.DB
	DB     *gorm.DB
	Router *mux.Router
	Port   int
}

func (a *App) Initialise(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=True",
		config.DB.User,
		config.DB.Pass,
		config.DB.Host,
		config.DB.DBName)

	fmt.Println(dbURI)

	db, err := gorm.Open("mysql", dbURI)

	if err != nil {
		panic(err)
	}

	a.DB = models.DBMigrate(db)
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

	a.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		handler.GetAllUsers(a.DB, w, r)
	})

	a.Get("/userblogs", a.GetBlogsForUser)
	a.Get("/blog/{id}", a.GetBlogByID)
	a.Get("/metastats", a.GetMetaStats)
}

func (a *App) Get(route string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(route, f).Methods("GET")
}

func (a *App) Post(route string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(route, f).Methods("POST")
}

// Handlers to manage Employee Data
func (a *App) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	handler.GetAllUsers(a.DB, w, r)
}

func (a *App) GetBlogsForUser(w http.ResponseWriter, r *http.Request) {
	// TODO: unhardcode 1 after implementing user login
	handler.GetAllBlogsForUser(1, a.DB, w, r)
}

func (a *App) GetBlogByID(w http.ResponseWriter, r *http.Request) {
	handler.GetBlogByID(a.DB, w, r)
}

func (a *App) GetMetaStats(w http.ResponseWriter, r *http.Request) {
	handler.GetOverallStats(a.DB, w, r)
}
