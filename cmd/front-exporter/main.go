package main

import (
	f "github.com/rhetoric-coffee/front-exporter/src/front"
	"github.com/urfave/cli"
	"log"
)

func main() {
	app := cli.NewApp()
    app.Name = "front-exporter"
    app.Usage = "metrics from the Front API"
    app.Action = func(c *cli.Context) error {
		frontApi, err := f.New()
		if err != nil {
			return err
		}
		// export teams
		teams, err :=frontApi.ListTeams()
		if err != nil {
			return err
		}
		for _, team := range(*teams) {
			log.Println("Start here next time: %v", team.Name)
		}
		return nil
	}

}

