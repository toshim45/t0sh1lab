package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

func info() {
	app.Name = "Site downloader"
	app.Usage = "Follows an URL and downloads rendered html content. Follows and download anchor links to child pages recursively. Only follows relative links."
}

func commands() {
	var url string

	urlFlag := cli.StringFlag{
		Name:        "url",
		Usage:       "Site url",
		Destination: &url,
	}

	app.Flags = []cli.Flag{
		&urlFlag,
	}
}

func main() {
	info()
	commands()

	app.Action = func(c *cli.Context) error {
		fmt.Printf("Passed url: %v", c.String("url"))
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
