package config

func Filepath(filepath string) {
	conf.Filepath = filepath
}

func ConsiderWorldAsPlayer(considerWorldAsPlayer bool) {
	conf.ConsiderWorldAsPlayer = considerWorldAsPlayer
}

func GroupBy(groupBy string) {
	conf.GroupBy = groupBy
}

func Get() Config {
	return conf
}
