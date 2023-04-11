package types

type SpecialEntity string

const (
	EntityWorld SpecialEntity = "<world>"
)

func SpecialEntities() []string {
	return []string{
		string(EntityWorld),
	}
}
