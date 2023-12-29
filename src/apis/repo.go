package apis

import (
	"context"

	"github.com/google/go-github/v57/github"
)

func ListReposName(owner string) ([]string, error) {
	var ctx = context.Background()
	var opt = &github.RepositoryListByUserOptions{
		ListOptions: github.ListOptions{PerPage: 10},
	}
	var repositoryNames []string = make([]string, 0)
	for {
		repos, resp, err := defaultClient.Repositories.ListByUser(ctx, owner, opt)
		if err != nil {
			return repositoryNames, err
		}
		for _, repo := range repos {
			repositoryNames = append(repositoryNames, *repo.FullName)
		}
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return repositoryNames, nil
}
