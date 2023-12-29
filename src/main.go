package main

import (
	"flag"
	"log"
	"strings"

	"github.com/kamontat/ghlabel-cloner/apis"
	"github.com/kamontat/ghlabel-cloner/configs"
	"github.com/kamontat/ghlabel-cloner/utils"
)

type ArrayArg []string

func (a *ArrayArg) Set(value string) error {
	*a = append(*a, value)
	return nil
}
func (a *ArrayArg) String() string {
	return strings.Join(*a, ", ")
}

var (
	configPaths ArrayArg
	owner       string
	repo        string
	replace     bool
)

var (
	name    string = "ghlabel"
	version string = "dev"
	date    string = "unknown"
)

func main() {
	utils.Info("Start %s version %s (%s)", name, version, date)

	var config = configs.Loads(configPaths)
	if owner == "" {
		log.Panicln("Owner (--owner) is required options")
	}

	var err error
	var repositories []string = make([]string, 0)
	if repo == "" {
		repositories, err = apis.ListReposName(owner)
		utils.MustNotError(err)
	} else {
		repositories = append(repositories, repo)
	}

	utils.Info("Updating %d repositories", len(repositories))
	utils.Debug("Repositories: %v", repositories)
	for _, repository := range repositories {
		utils.Info("Cloning %d labels to repository '%s/%s'", len(config.Labels), owner, repository)

		if replace {
			labels, err := apis.ListLabels(owner, repository)
			utils.MustNotError(err)

			utils.Info("Deleting all existed %d labels", len(labels))
			err = apis.DeleteLabels(owner, repository)
			utils.MustNotError(err)

			utils.Info("Creating %d labels", len(config.Labels))
		} else {
			utils.Info("Updating %d labels", len(config.Labels))
		}

		var dedup map[string]bool = make(map[string]bool)
		var errors []error
		for _, label := range config.Labels {
			if key, ok := dedup[label.Name]; key && ok {
				utils.Error("Duplicated label found: %s", label.Name)
				continue
			}

			err = apis.CreateOrUpdateLabel(owner, repository, label)
			errors = append(errors, err)
			dedup[label.Name] = true
		}

		err = utils.MergeErrors(errors)
		utils.MustNotError(err)
	}
}

func init() {
	flag.Var(&configPaths, "configs", "Config path can contains multiple files")
	flag.StringVar(&owner, "owner", "", "Repository owner")
	flag.StringVar(&repo, "repo", "", "Repository name; if not exist, will updates all repository from owner")
	flag.BoolVar(&replace, "replace", false, "Replace existed labels with config by delete all labels and recreate it")

	flag.Parse()
}
