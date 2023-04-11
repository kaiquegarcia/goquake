package main

import (
	"goquake/cmd"
)

func main() {
	if keepRunning := cmd.Bootstrap(); keepRunning {
		cmd.Run()
	}
}
