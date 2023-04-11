package game

import "goquake/internal/event"

type Stat struct {
	TotalKills int            `json:"total_kills"`
	Players    []string       `json:"players"`
	Kills      map[string]int `json:"kills"`
}

func (stat *Stat) initializePlayer(playerName string) {
	if _, exists := stat.Kills[playerName]; !exists {
		stat.Kills[playerName] = 0
		stat.Players = append(stat.Players, playerName)
	}
}

func (stat *Stat) AddKill(event event.Event, considerWorldAsPlayer bool) {
	stat.TotalKills++
	stat.initializePlayer(event.Target)
	isAuthorPlayer := event.IsAuthorPlayer()
	shouldScoreAuthor := considerWorldAsPlayer || isAuthorPlayer

	if shouldScoreAuthor {
		stat.initializePlayer(event.Author)
		stat.Kills[event.Author]++
	}

	if !isAuthorPlayer {
		stat.Kills[event.Target]--
	}
}
