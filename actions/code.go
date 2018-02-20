package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/google/go-github/github"
)

// CodeIndex default implementation.
func CodeIndex(c buffalo.Context) error {
	return c.Render(200, r.HTML("code/index.html"))
}

// CodeToGetRepos default implementation.
func CodeToGetRepos(c buffalo.Context) error {
	// list public repositories for org "github"
	client := github.NewClient(nil)

	user := "obiknows"
	opt := &github.RepositoryListOptions{Type: "owner", Sort: "updated", Direction: "desc"}
	repos, _, _ := client.Repositories.List(c, user, opt)

	// use the bufflao context to pass in the data  (a go web concept)
	c.Set("repos", []string{"this one", "that one"})

	// returns 'obiknows' repos
	return c.Render(200, r.JSON(repos))
}
