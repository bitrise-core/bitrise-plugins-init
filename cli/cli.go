package cli

import (
	"fmt"
	"os"
	"path"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

// Run ...
func Run() {
	cli.VersionPrinter = func(c *cli.Context) { fmt.Println(c.App.Version) }

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		ForceColors:     true,
		TimestampFormat: "15:04:05",
	})

	app := cli.NewApp()

	app.Name = path.Base(os.Args[0])
	app.Usage = "Init bitrise config"
	app.Version = "0.9.0"

	app.Author = ""
	app.Email = ""

	app.Commands = []cli.Command{
		configCmd,
		manualConfigCmd,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
