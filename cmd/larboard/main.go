package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/stephen-fox/larboard/halo2"
	"github.com/stephen-fox/larboard/ipc"
	"github.com/stephen-fox/larboard/ipc/research"
	"github.com/stephen-fox/larboard/ipc/cartography"
	"github.com/stephen-fox/larboard/mapio"
)

const (
	mapFilePathArg  = "m"

	researchActionsArg = "R"
	doResearchArg      = "r"

	helpArg = "h"
)

var (
	mapFilePath = flag.String(mapFilePathArg, "", "The path to the map file")

	doResearch      = flag.String(doResearchArg, "", "Execute a research action")
	researchActions = flag.Bool(researchActionsArg, false, "Print the research actions")

	printHelp = flag.Bool(helpArg, false, "Print this help page")

	cartographerActionsToFuncs = map[string]func(data ipc.Input) ipc.Result{
		"sign": cartography.Sign,
	}

	researchActionsToFuncs = map[string]func(data ipc.Input) ipc.Result{
		"valid":     research.IsMapValid,
		"name":      research.MapName,
		"scenario":  research.Scenario,
		"signature": research.Signature,
	}
)

func main() {
	flag.Parse()

	if len(os.Args) <= 1 || *printHelp {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *researchActions {
		fmt.Println("Available research actions:")

		for action := range researchActionsToFuncs {
			fmt.Println("    " + action)
		}

		os.Exit(0)
	}

	if len(strings.TrimSpace(*mapFilePath)) == 0 {
		fatal("Please specify a map file using '-" + mapFilePathArg + " /path/to/map/file'")
	}

	if len(strings.TrimSpace(*doResearch)) > 0 {
		f, ok := researchActionsToFuncs[*doResearch]
		if !ok {
			fatal("The specified research action does not exist - '" + *doResearch + "'")
		}

		hmap, err := mapio.NewMap(*mapFilePath)
		if err != nil {
			fatal(err.Error())
		}

		h2Researcher, err := halo2.NewResearcher(hmap)
		if err != nil {
			fatal("Failed to load map researcher" + err.Error())
		}

		data := research.Data{
			Researcher: h2Researcher,
		}

		result := f(data)
		if result.IsError() {
			fatal(result.FormatOutput(ipc.SingleRunCli))
		}

		fmt.Println(result.FormatOutput(ipc.SingleRunCli))
	}
}

func fatal(reason string) {
	log.Fatal("[ERROR] - ", reason)
}