package game

import "goquake/internal/event"

type Data struct {
	Events []event.Event `json:"events"`
}

func New(events []event.Event) Data {
	return Data{
		Events: events,
	}
}

func (game Data) Stats(considerWorldAsPlayer bool) *Stat {
	stats := &Stat{
		TotalKills: 0,
		Players:    make([]string, 0),
		Kills:      make(map[string]int),
	}

	for index := 0; index < len(game.Events); index++ {
		event := game.Events[index]
		if event.IsKillEvent() {
			stats.AddKill(event, considerWorldAsPlayer)
		}
	}

	return stats
}

func (game Data) KillsByMeans() map[string]int {
	kills := make(map[string]int)
	for _, event := range game.Events {
		if event.IsKillEvent() {
			mean := string(event.Mean)
			if _, exists := kills[mean]; !exists {
				kills[mean] = 0
			}

			kills[mean]++
		}
	}

	return kills
}
