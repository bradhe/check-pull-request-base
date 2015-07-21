package main

import (
	"flag"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"log"
	"os"
	"strings"
)

var (
	GithubAccessToken = flag.String("access-token", "", "The token to use when authenticating with Github.")
	RepoName          = flag.String("repo", "", "The repository to check.")
	PullRequestNumber = flag.Int("pull-request", 0, "The pull request number to fetch data for.")
)

func Usage(message string) {
	fmt.Println(message)
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	flag.Parse()

	if *GithubAccessToken == "" {
		Usage("Please specify an access-token")
	}

	if *RepoName == "" {
		Usage("Please specify a repo")
	}

	if *PullRequestNumber < 1 {
		Usage("Please specify a pull-request")
	}

	comps := strings.Split(*RepoName, "/")

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: *GithubAccessToken},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)
	pr, _, err := client.PullRequests.Get(comps[0], comps[1], *PullRequestNumber)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(*pr.Base.Ref)
}
