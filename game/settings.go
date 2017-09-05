package game

type ServerSettings struct {
	URL  string `yaml:"url"`
	Port string `yaml:"port"`
}

type Settings struct {
	Server *ServerSettings `yaml:"server"`
}
