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

	// Meta routes
	a.Get("/health", a.HealthCheck)
	a.Get("/metastats", a.GetMetaStats)

	// Profile & Users
	a.Get("/@{username}", a.GetProfileForUsername)
	a.Get("/users", a.GetAllUsers)

	// Blogs
	a.Get("/blog/{id}", a.GetBlogByID)

	// Account
	a.Post("/login", a.Login)
	a.Post("/signup", a.CreateNewUser)

	// Protected routes
	a.Post("/blog/new", handler.Protected(a.CreateNewBlog))
	a.Post("/jott/new", handler.Protected(a.CreateNewJott))
}

func (a *App) Get(route string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(route, f).Methods("GET")
}

func (a *App) Post(route string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(route, f).Methods("POST")
}

func (a *App) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	handler.GetAllUsers(a.DB, w, r)
}

func (a *App) GetProfileForUsername(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	handler.GetProfile(username, a.DB, w, r)
}

func (a *App) GetBlogByID(w http.ResponseWriter, r *http.Request) {
	handler.GetBlogByID(a.DB, w, r)
}

func (a *App) GetMetaStats(w http.ResponseWriter, r *http.Request) {
	handler.GetOverallStats(a.DB, w, r)
}

func (a *App) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	handler.CreateNewUser(a.DB, w, r)
}

func (a *App) CreateNewBlog(w http.ResponseWriter, r *http.Request) {
	handler.CreateNewBlog(a.DB, w, r)
}

func (a *App) CreateNewJott(w http.ResponseWriter, r *http.Request) {
	handler.CreateJott(a.DB, w, r)
}

func (a *App) Login(w http.ResponseWriter, r *http.Request) {
	handler.Login(a.DB, w, r)
}

func (a *App) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}
