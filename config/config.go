package config

type Configurations struct {
	Rest RestConfigurations
}

type RestConfigurations struct {
	Url            string
	LevelZeroBody  string
	LevelOneBody   string
	LevelTwoBody   string
	LevelThreeBody string
	LevelFourBody  string
	LevelFiveBody  string
}
