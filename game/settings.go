package game

// ServerSettings are the yaml representation of the redis settings.
type ServerSettings struct {
	URL  string `yaml:"url"`
	Port string `yaml:"port"`
}

// Settings is the model that stores all the different configuration models.
type Settings struct {
	Server *ServerSettings `yaml:"server"`
}
