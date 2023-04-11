package cmd

import (
	"encoding/json"
	"fmt"
	"goquake/internal/config"
	"goquake/internal/game"
	"goquake/internal/log2memory"
)

func Run() {
	logs, err := log2memory.ReadFileContents()
	if err != nil {
		fmt.Printf("could not parse file %s: %s\n", config.Get().Filepath, err)
		return
	}

	output := game.ParseLog(logs)
	if len(output) == 0 {
		fmt.Println("there's no games to show")
		return
	}

	bytes, err := json.MarshalIndent(output, "", "    ")
	if err != nil {
		fmt.Printf("could not marshal data: %s", err)
		return
	}

	fmt.Printf("There's the statistic of each game present in the logs:\n%s\n", bytes)
}
