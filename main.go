package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	_ "label/migrations"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

// This will be injected by the build process
var Version string

//go:embed frontend/build/*
var frontend embed.FS

func main() {
	fmt.Println("Running Version:", Version)
	app := pocketbase.New()

	// loosely check if it was executed using "go run"
	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		fs := echo.MustSubFS(frontend, "frontend/build")
		e.Router.StaticFS("/app", fs)

		e.Router.GET("/", func(c echo.Context) error {
			return c.Redirect(http.StatusPermanentRedirect, "/app")
		})

		return nil
	})

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		// in dev only
		Automigrate: isGoRun,
	})

	app.OnModelAfterCreate("images").Add(func(e *core.ModelEvent) error {
		id := e.Model.GetId()
		record, err := e.Dao.FindRecordById("images", id)
		if err != nil {
			return err
		}

		fileKey := record.BaseFilesPath() + "/" + record.GetString("file")

		fs, err := app.NewFilesystem()
		if err != nil {
			return err
		}
		defer fs.Close()

		br, err := fs.GetFile(fileKey)
		if err != nil {
			return err
		}
		defer br.Close()

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
