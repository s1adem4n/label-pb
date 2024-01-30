package main

import (
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
	"github.com/spf13/cobra"

	"label/frontend"
	"label/pkg/version"
)

var Version string

func main() {
	if Version == "" {
		fmt.Println("Version was not set when this binary was built.")
		fmt.Println("This is fine for development, but please set a version when building for production.")
	} else {
		fmt.Println("Running Version:", Version)
	}
	latestRelease, err := version.GetLatestRelease()
	if err != nil {
		fmt.Println("Failed to get latest release:", err)
	} else {
		fmt.Println("Latest Version:", latestRelease.TagName)
	}

	app := pocketbase.New()

	app.RootCmd.AddCommand(&cobra.Command{
		Use: "install",
		Run: func(cmd *cobra.Command, args []string) {
			err := version.Install()
			if err != nil {
				fmt.Println("Failed to install:", err)
			}
		},
		Short: "Installs the binary to a new location (~/label)",
	})
	app.RootCmd.AddCommand(&cobra.Command{
		Use: "update",
		Run: func(cmd *cobra.Command, args []string) {
			err := version.Update()
			if err != nil {
				fmt.Println("Failed to update:", err)
			}
		},
		Short: "Updates the binary to the latest version",
	})

	// loosely check if it was executed using "go run"
	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		fs := echo.MustSubFS(frontend.Frontend, "frontend/build")
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
