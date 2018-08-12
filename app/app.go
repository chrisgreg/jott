package app

import (
	"database/sql"
	"fmt"

	"github.com/chrisgreg/jott/config"
)

type App struct {
	DB   *sql.DB
	Port int
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
}

// func (a *App) Run(config *config.Config) {

// }
