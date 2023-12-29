package apis

import (
	"log"
	"os"

	"github.com/google/go-github/v57/github"
)

func getToken() string {
	var token, ok = os.LookupEnv("GITHUB_TOKEN")
	if ok {
		return token
	}

	log.Panicln("Token from environment 'GITHUB_TOKEN' is required")
	return ""
}

func NewClient() *github.Client {
	var client = github.NewClient(nil)
	return client.WithAuthToken(getToken())
}

var defaultClient = NewClient()
