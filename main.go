package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/jason0x43/go-alfred"
)

var cacheFile string
var configFile string
var workflow alfred.Workflow

var config struct {
	Service               string `desc:"Service to use"`
	ForecastIOKey         string `desc:"Your API key for forecast.io"`
	WeatherUndergroundKey string `desc:"Your API key for Weather Underground"`
	Icons                 string `desc:"Icon set"`
	Days                  int
	ShowLocaltime         bool
	TimeFormat            string   `desc:"Timestamp format"`
	Location              Location `desc:"Default location"`
	Units                 string   `desc:"Units"`
}

var cache struct {
	Weather Weather
	Time    time.Time
	Service string
}

var dlog = log.New(os.Stderr, "[weather] ", log.LstdFlags)

func main() {
	var err error

	workflow, err = alfred.OpenWorkflow(".", true)
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}

	configFile = path.Join(workflow.DataDir(), "config.json")
	cacheFile = path.Join(workflow.CacheDir(), "cache.json")

	dlog.Println("Using config file", configFile)
	dlog.Println("Using cache file", cacheFile)

	if err := alfred.LoadJSON(configFile, &config); err == nil {
		dlog.Println("loaded config")
	}

	if err := alfred.LoadJSON(cacheFile, &cache); err == nil {
		dlog.Println("loaded cache")
	}

	commands := []alfred.Command{
		ForecastCommand{},
		ConfigCommand{},
	}

	workflow.Run(commands)
}
