package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/stephen-fox/larboard/halo2"
)

const (
	mapFilePathArg = "m"
	isMapArg       = "is-map"
	getNameArg     = "get-name"
	getScenarioArg = "get-scenario"

	helpArg = "h"
)

var (
	mapFilePath = flag.String(mapFilePathArg, "", "The path to the map file")

	isMapValid  = flag.Bool(isMapArg, false, "Verify that the map file is actually a Halo map")
	getName     = flag.Bool(getNameArg, false, "Get the map's name")
	getScenario = flag.Bool(getScenarioArg, false, "Get the map's scenario")

	printHelp = flag.Bool(helpArg, false, "Print this help page")
)

func main() {
	flag.Parse()

	if len(os.Args) <= 1 || *printHelp {
		flag.PrintDefaults()
		os.Exit(0)
	}

	err := validateMap()
	if err != nil {
		fatal("", err)
	}

	h2Researcher, err := halo2.NewResearcher(*mapFilePath)
	if err != nil {
		fatal("Failed to load map researcher", err)
	}

	if *isMapValid {
		err := h2Researcher.IsMap()
		if err != nil {
			fatal("The specified file is not a Halo map", err)
		}

		log.Println("Yep, it's a map")
	}

	if *getName {
		name, err := h2Researcher.Name()
		if err != nil {
			fatal("Failed to get map's name", err)
		}

		fmt.Println(name)
	}

	if *getScenario {
		scenario, err := h2Researcher.Scenario()
		if err != nil {
			fatal("Failed to get map's scenario", err)
		}

		fmt.Println(scenario)
	}

}

func validateMap() error {
	if len(strings.TrimSpace(*mapFilePath)) == 0 {
		return errors.New("Please specify the path to the map file using '-" +
			mapFilePathArg + " /path/to/map/file'")
	}

	return nil
}

func fatal(whatHappened string, err error) {
	if whatHappened == "" {
		log.Fatal("[ERROR] ", err.Error())
	}

	log.Fatal("[ERROR] - ", whatHappened, " - ", err.Error())
}