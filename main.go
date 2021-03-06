package main

import (
	"errors"
	"os"

	"github.com/Jumpscale/go-raml/commands"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

//Version define software version
var Version = "0.1-Dev"

//ApplicationName is the name of the application
var ApplicationName = "RAML code generation toolset"

var (
	serverCommand = &commands.ServerCommand{}
	clientCommand = &commands.ClientCommand{}
	specCommand   = &commands.SpecCommand{}
)

func main() {
	app := cli.NewApp()
	app.Name = ApplicationName
	app.Version = Version
	app.Usage = "Using a RAML specification, generate server and client code or a RAML specification from go code."

	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	var debugLogging bool
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug, d",
			Usage:       "Enable debug logging",
			Destination: &debugLogging,
		},
	}
	app.Before = func(c *cli.Context) error {
		if debugLogging {
			log.SetLevel(log.DebugLevel)
			log.Debug("Debug logging enabled")
			log.Debug(ApplicationName, "-", Version)
		}
		return nil
	}
	app.Commands = []cli.Command{
		{
			Name:  "server",
			Usage: "Generate a go server according to a RAML specification",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "language, l",
					Value:       "go",
					Usage:       "Language to construct a server for",
					Destination: &serverCommand.Language,
				},
				cli.StringFlag{
					Name:        "dir",
					Value:       ".",
					Usage:       "target directory",
					Destination: &serverCommand.Dir,
				},
				cli.StringFlag{
					Name:        "ramlfile",
					Value:       ".",
					Usage:       "Source raml file",
					Destination: &serverCommand.RamlFile,
				},
				cli.StringFlag{
					Name:        "package",
					Value:       "main",
					Usage:       "package name",
					Destination: &serverCommand.PackageName,
				},
				cli.BoolFlag{
					Name:        "no-main",
					Usage:       "Do not generate a main.go file",
					Destination: &serverCommand.NoMainGeneration,
				},
			},
			Action: func(c *cli.Context) {
				if err := serverCommand.Execute(); err != nil {
					log.Error(err)
				}
			},
		},
		{
			Name:  "client",
			Usage: "Create a client for a RAML specification",

			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "language, l",
					Value:       "go",
					Usage:       "Language to construct a client for",
					Destination: &clientCommand.Language,
				},
				cli.StringFlag{
					Name:        "dir",
					Value:       ".",
					Usage:       "target directory",
					Destination: &clientCommand.Dir,
				},
				cli.StringFlag{
					Name:        "ramlfile",
					Value:       ".",
					Usage:       "Source raml file",
					Destination: &clientCommand.RamlFile,
				},
			},
			Action: func(c *cli.Context) {
				if err := clientCommand.Execute(); err != nil {
					log.Error(err)
				}
			},
		}, {
			Name:  "spec",
			Usage: "Generate a RAML specification from a go server",
			Action: func(c *cli.Context) {
				err := errors.New("Not implemented, check the roadmap")
				log.Error(err)
			},
		},
	}

	app.Action = func(c *cli.Context) {
		cli.ShowAppHelp(c)
	}

	app.Run(os.Args)
}
