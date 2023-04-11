package game

import (
	"goquake/internal/event"
	"goquake/internal/types"
	"regexp"
)

var logPattern = regexp.MustCompile(`(?s)(\d{1,3}):(\d{2}) (\w{2,15}): (.{1,120})\n`)
var killPattern = regexp.MustCompile(`(\d{1,4}) (\d{1,4}) (\d{1,3}): (.{2,15}) killed (.{2,15}) by ([A-Z_]{4,30})`)

func parseGameLog(gameLog string) []event.Event {
	var events []event.Event
	// classic loop to keep events order
	matches := logPattern.FindAllStringSubmatch(gameLog, -1)
	for index := 0; index < len(matches); index++ {
		match := matches[index]
		event := event.Event{
			Time: match[1] + ":" + match[2],
			Type: types.EventType(match[3]),
			Data: match[4],
		}

		if event.Type == types.TypeKill {
			cmdMatch := killPattern.FindAllStringSubmatch(match[4], -1)
			event.Position = cmdMatch[0][1] + " " + cmdMatch[0][2] + " " + cmdMatch[0][3]
			event.Author = cmdMatch[0][4]
			event.Target = cmdMatch[0][5]
			event.Mean = types.EventMean(cmdMatch[0][6])
		}
		events = append(events, event)
	}

	return events
}
