package event

import "goquake/internal/types"

type Event struct {
	Type     types.EventType
	Author   string
	Target   string
	Position string
	Time     string
	Mean     types.EventMean
	Data     string
}

func (event Event) IsAuthorPlayer() bool {
	entities := types.SpecialEntities()
	for _, entity := range entities {
		if event.Author == entity {
			return false
		}
	}

	return true
}

func (event Event) IsKillEvent() bool {
	return event.Type == types.TypeKill
}

func (event Event) IsScoreEvent() bool {
	return event.Type == types.TypeScore
}
