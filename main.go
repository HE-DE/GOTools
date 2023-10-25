package main

import (
	"goscanner/factory"
	"goscanner/utils"
	"os"
)

func main() {
	args := utils.ParseArgs()

	strategy := factory.ProcessorFactory(args[0])
	if strategy != nil {
		strategy.Process(args[1:])
	} else {
		os.Exit(1)
	}
}
