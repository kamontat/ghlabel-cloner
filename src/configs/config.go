package configs

import (
	"log"
	"os"

	"github.com/kamontat/ghlabel-cloner/utils"
	"gopkg.in/yaml.v3"
)

func Load(path string) Config {
	var buffer, err = os.ReadFile(path)
	utils.MustNotError(err)

	var config Config
	err = yaml.Unmarshal(buffer, &config)
	utils.MustNotError(err)

	return config
}

func Loads(paths []string) Config {
	if len(paths) <= 0 {
		log.Panicln("Invalid config path, required at least one file")
	}

	var config = Config{
		Labels: make([]ConfigLabel, 0),
	}
	for _, path := range paths {
		var conf = Load(path)
		config.Labels = append(config.Labels, conf.Labels...)
	}

	return config
}
