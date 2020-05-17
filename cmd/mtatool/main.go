package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/go-yaml/yaml"

	"github.com/qystishere/mtatool/gta"
	"github.com/qystishere/mtatool/gta/parser/resource"
	"github.com/qystishere/mtatool/mta"
	mtaItem "github.com/qystishere/mtatool/mta/parser/item"
)

var config struct {
	GTA struct {
		Path string `yaml:"path"`
	} `yaml:"gta"`
	MTA struct {
		Path string `yaml:"path"`
		Maps struct {
			Include []string `yaml:"include"`
			Exclude []string `yaml:"exclude"`
			Only    []string `yaml:"only"`
			Sync    bool     `yaml:"sync"`
		} `yaml:"maps"`
	} `yaml:"mta"`
}

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetHandler(cli.New(os.Stdout))
	log.Info("MTA Tool | github.com/qystishere/mtatool")

	file, err := os.Open("mtatool.yml")
	if err != nil {
		log.WithError(err).
			Fatal("Can't open config file")
	}

	if err = yaml.NewDecoder(file).Decode(&config); err != nil {
		log.WithError(err).
			Fatal("Can't decode config file")
	}
}

func main() {
	game, err := gta.Load(config.GTA.Path)
	if err != nil {
		log.WithError(err).
			Fatal("Can't load gta")
	}

	server, err := mta.Load(config.MTA.Path)
	if err != nil {
		log.WithError(err).
			Fatal("Can't load mta")
	}

	log.Debugf("GTA SA path: %s", game.Path)
	log.Debugf("MTA path: %s", server.Path)

	var (
		path = game.Data.Path + "maps/test"
		_    = os.MkdirAll(path, os.ModePerm)
	)

	for _, m := range server.Mod.Maps {
		name := strings.Split(m.FileInfo.Name(), ".")[0]
		log.Debugf("MTA Map: %s", name)

		ipl, err := resource.CreateIPL(path + fmt.Sprintf("/%s.ipl", name))
		if err != nil {
			log.WithError(err).
				Fatalf("Can't create ipl (%s)", name)
		}

		for _, item := range m.Items {
			switch item := item.(type) {
			case mtaItem.Object:
				fmt.Println(item.Type(), item.ModelID, item.ModelID)
			}
			if err := ipl.Add(item); err != nil {
				log.WithError(err).
					Warnf("Can't add item (%s)", item.Type())
			}
		}

		if err := ipl.Save(); err != nil {
			log.WithError(err).
				Fatalf("Can't save IPL (%s)", name)
		}
	}

	/*
		if err := game.Data.Save(); err != nil {
			log.WithError(err).
				Fatal("Can't save gta data")
		}
	*/

}
