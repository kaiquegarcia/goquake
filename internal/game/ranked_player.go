package game

type RankedPlayer struct {
	Rank  int    `json:"rank"`
	Name  string `json:"name"`
	Kills int    `json:"kills"`
}
