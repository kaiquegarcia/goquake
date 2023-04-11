package types

type EventType string

const (
	TypeInit                  EventType = "InitGame"
	TypeExit                  EventType = "Exit"
	TypeClientConnect         EventType = "ClientConnect"
	TypeClientDisconnect      EventType = "ClientDisconnect"
	TypeClientUserInfoChanged EventType = "ClientUserinfoChanged"
	TypeClientBegin           EventType = "ClientBegin"
	TypeShutdownGame          EventType = "ShutdownGame"
	TypeItem                  EventType = "Item"
	TypeKill                  EventType = "Kill"
)

func EventTypes() []EventType {
	return []EventType{
		TypeKill,
	}
}
