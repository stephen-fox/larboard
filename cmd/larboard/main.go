package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/stephen-fox/larboard"
	"github.com/stephen-fox/larboard/ipc"
)

const (
	defaultInstanceGameArg = "g"
	defaultInstanceMapArg  = "d"
	humanReadableOutputArg = "human"
	helpArg                = "h"
)

var (
	defaultInstanceGame = flag.String(defaultInstanceGameArg, string(ipc.Halo2), "The game type to use for the default instance")
	defaultInstanceMap  = flag.String(defaultInstanceMapArg, "", "The map file to open for the default instance")
	humanReadableOutput = flag.Bool(humanReadableOutputArg, false, "Output in human readable format")

	printHelp = flag.Bool(helpArg, false, "Print this help page")
)

func main() {
	flag.Parse()

	if *printHelp {
		flag.PrintDefaults()
		os.Exit(0)
	}

	ipcManager, err := ipc.NewManager()
	if err != nil {
		fatal(err.Error())
	}

	if len(strings.TrimSpace(*defaultInstanceMap)) > 0 {
		if len(strings.TrimSpace(*defaultInstanceGame)) == 0 {
			fatal("Please specify a default instance type")
		}

		options := ipc.IoOptions{
			HumanReadableOutput: *humanReadableOutput,
			Source:              ipc.Cli,
		}

		details := ipc.InstanceDetails{
			Game: ipc.Game(*defaultInstanceGame),
			Id:   "default",
			InitialMap: larboard.HaloMap{
				FilePath: *defaultInstanceMap,
			},
		}

		err := setDefaultHaloMap(ipcManager, options, details)
		if err != nil {
			fatal(err.Error())
		}
	}

	inputs := make(chan string)
	results := make(chan ipc.Result)
	stdout := make(chan string)

	go func() {
		for {
			input, err := getUserInput()
			if err != nil {
				fatal(err.Error())
			}

			inputs <- input
		}
	}()

	go func() {
		for {
			output := <-results
			stdout <- output.FormatOutput()
		}
	}()

	go func() {
		for {
			fmt.Println(<-stdout)
		}
	}()

	ioOptions := ipc.IoOptions{
		Source:              ipc.Cli,
		HumanReadableOutput: *humanReadableOutput,
	}

	err = ipcManager.BlockAndParseInput(ioOptions, inputs, results)
	if err != nil {
		fatal(err.Error())
	}
}

func setDefaultHaloMap(manager ipc.Manager, options ipc.IoOptions, details ipc.InstanceDetails) error {
	i, err := ipc.NewInstance(options, details)
	if err != nil {
		return err
	}

	err = manager.AddInstance(i)
	if err != nil {
		return err
	}

	return nil
}

func getUserInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	readD := '\n'

	input, err := reader.ReadString(byte(readD))
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(input, string(readD)), nil
}

func fatal(reason string) {
	log.Fatal("[ERROR] - ", reason)
}