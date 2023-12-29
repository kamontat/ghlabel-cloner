package main

import (
	"flag"

	"github.com/kamontat/ghlabel-cloner/apis"
	"github.com/kamontat/ghlabel-cloner/configs"
	"github.com/kamontat/ghlabel-cloner/utils"
)

var (
	configPath string
	owner      string
	repo       string
	replace    bool
)

func main() {
	var config = configs.Load(configPath)
	utils.Info("Cloning %d labels to repository '%s/%s'", len(config.Labels), owner, repo)

	var err error
	if replace {
		utils.Info("Deleting all existed labels")
		err = apis.DeleteLabels(owner, repo)
		utils.MustNotError(err)
	}

	var errors []error
	for i, label := range config.Labels {
		utils.Info("Processing... label #%d", i+1)
		err = apis.CreateOrUpdateLabels(owner, repo, label)
		errors = append(errors, err)
	}

	err = utils.MergeErrors(errors)
	utils.MustNotError(err)
}

func init() {
	flag.StringVar(&configPath, "config", "", "Config path")
	flag.StringVar(&owner, "owner", "kc-workspace", "Repository owner")
	flag.StringVar(&repo, "repo", "", "Repository name")
	flag.BoolVar(&replace, "replace", false, "Replace existed labels with config")

	flag.Parse()
}
