package actions

import (
	"fmt"
	"github.com/gobuffalo/buffalo"
	"github.com/google/go-github/github" 
)

// CodeIndex default implementation.
func CodeIndex(c buffalo.Context) error {
	return c.Render(200, r.HTML("code/index.html"))
}

// CodeApi default implementation.
func CodeToGetRepos(c buffalo.Context) error {

	// list public repositories for org "github"
	client := github.NewClient(nil)

	user := "obiknows"
	opt := &github.RepositoryListOptions{Type: "owner", Sort: "updated", Direction: "desc"}

	repos, _, err := client.Repositories.List(c, user, opt)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Printf("Recently updated repositories by %q: %v", github.Stringify(repos))

	return c.Render(200, r.JSON(repos))
}
