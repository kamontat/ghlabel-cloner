package configs

import (
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
