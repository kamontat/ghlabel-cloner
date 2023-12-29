package configs

type ConfigLabel struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Color       string `yaml:"color"`
}

type Config struct {
	Labels []ConfigLabel `yaml:"labels"`
}
