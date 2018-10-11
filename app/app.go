package app

import (
	"os"
	"path/filepath"

	"github.com/skanehira/mockapi/app/config"
	"github.com/skanehira/mockapi/app/db"
	"github.com/skanehira/mockapi/app/server"
)

type App struct {
	AppDir string
	Config *config.Config
	Server *server.Server
}

func New() *App {
	app := &App{}

	return app.createAppDir().
		newConfig().
		newServer()
}

func (a *App) createAppDir() *App {
	appDir := filepath.Join(os.Getenv("HOME"), ".mockapi")
	if _, err := os.Stat(appDir); os.IsNotExist(err) {
		os.Mkdir(appDir, 0777)
	}

	a.AppDir = appDir
	return a
}

func (a *App) newConfig() *App {
	a.Config = config.New(filepath.Join(a.AppDir, "config.yaml"))
	return a
}

func (a *App) newDB() *db.DB {
	c := a.Config.DB
	return db.New(c.DBType, filepath.Join(a.AppDir, "mockapi.db"), c.LogMode)
}

func (a *App) newServer() *App {
	c := a.Config
	a.Server = server.New(c.Protocol, c.Address, c.Port, c.CertFile, c.CertKeyFile, a.newDB())
	return a
}

func (a *App) Run() {
	a.Server.Run()

	// TODO run panel
}
