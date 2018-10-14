package app

import (
	"os"
	"path/filepath"

	"github.com/skanehira/mockapi/app/config"
	"github.com/skanehira/mockapi/app/db"
	"github.com/skanehira/mockapi/app/server"
	"github.com/skanehira/mockapi/app/view"
)

type App struct {
	dir    string
	server *server.Server
	view   *view.View
}

func New() *App {
	return &App{}
}

func (app *App) Setup() {
	dir := app.createAppDir()

	config := config.New(filepath.Join(dir, "config.yaml"))
	db := db.New(config.DB.DBType, filepath.Join(dir, "mockapi.db"), config.DB.LogMode)

	// TODO 引数でmigrateする
	db.Migration()

	//app.server = server.New(db, config)
	app.view = view.New(db, config)
	app.view.Setup()
}

func (a *App) createAppDir() string {
	homeDir := os.Getenv("HOME")
	if homeDir == "" {
		panic("Undefined Enviroment $HOME")
	}

	dir := filepath.Join(homeDir, ".mockapi")

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0777)
	}
	return dir
}

func (a *App) Run() {
	//go a.server.Run()

	if err := a.view.Run(); err != nil {
		panic(err)
	}
}
