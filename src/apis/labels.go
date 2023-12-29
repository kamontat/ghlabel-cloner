package apis

import (
	"context"
	"fmt"

	"github.com/google/go-github/v57/github"
	"github.com/kamontat/ghlabel-cloner/configs"
	"github.com/kamontat/ghlabel-cloner/utils"
)

var labelsCache map[string]map[string]*github.Label = make(map[string]map[string]*github.Label)

func mapLabels(labels []*github.Label) (mapper map[string]*github.Label) {
	mapper = make(map[string]*github.Label)
	for _, label := range labels {
		mapper[*label.Name] = label
	}
	return
}

func saveCache(owner, repo string, labels map[string]*github.Label) {
	labelsCache[fmt.Sprintf("%s/%s", owner, repo)] = labels
}

func getCache(owner, repo string) (labels map[string]*github.Label, ok bool) {
	labels, ok = labelsCache[fmt.Sprintf("%s/%s", owner, repo)]
	return
}

func clearCache(owner, repo string, labels map[string]*github.Label) {
	delete(labelsCache, fmt.Sprintf("%s/%s", owner, repo))
}

func ListLabels(owner, repo string) (map[string]*github.Label, error) {
	if labels, ok := getCache(owner, repo); ok {
		return labels, nil
	}

	var ctx = context.Background()
	var labelMap map[string]*github.Label
	labels, _, err := defaultClient.Issues.ListLabels(ctx, owner, repo, nil)
	if err == nil {
		labelMap = mapLabels(labels)
		saveCache(owner, repo, labelMap)
	}

	return labelMap, err
}

func CreateOrUpdateLabels(owner, repo string, labelConfig configs.ConfigLabel) error {
	var labels, err = ListLabels(owner, repo)
	if err != nil {
		return err
	}

	if _, ok := labels[labelConfig.Name]; ok {
		return UpdateLabel(owner, repo, labelConfig)
	} else {
		return CreateLabel(owner, repo, labelConfig)
	}
}

func DeleteLabels(owner, repo string) error {
	labels, err := ListLabels(owner, repo)
	if err != nil {
		return err
	}

	var errors []error
	for _, label := range labels {
		err = DeleteLabel(owner, repo, *label.Name)
		errors = append(errors, err)
	}
	err = utils.MergeErrors(errors)
	if err != nil {
		return err
	}

	clearCache(owner, repo, labels)
	return nil
}

func CreateLabel(owner, repo string, labelConfig configs.ConfigLabel) error {
	utils.Debug("Creating label: %s (%v)", labelConfig.Name, labelConfig)
	_, _, err := defaultClient.Issues.CreateLabel(
		context.Background(),
		owner,
		repo,
		&github.Label{
			Name:        &labelConfig.Name,
			Description: &labelConfig.Description,
			Color:       &labelConfig.Color,
		})
	return err
}

func UpdateLabel(owner, repo string, labelConfig configs.ConfigLabel) error {
	utils.Debug("Updating label: %s (%v)", labelConfig.Name, labelConfig)
	_, _, err := defaultClient.Issues.EditLabel(
		context.Background(),
		owner,
		repo,
		labelConfig.Name,
		&github.Label{
			Name:        &labelConfig.Name,
			Description: &labelConfig.Description,
			Color:       &labelConfig.Color,
		})
	return err
}

func DeleteLabel(owner, repo, name string) error {
	utils.Debug("Deleting label: %s", name)
	_, err := defaultClient.Issues.DeleteLabel(
		context.Background(),
		owner,
		repo,
		name,
	)
	return err
}
