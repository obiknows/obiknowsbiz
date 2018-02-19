package actions

import "github.com/gobuffalo/buffalo"

// ResearchIndex default implementation.
func ResearchIndex(c buffalo.Context) error {
	return c.Render(200, r.HTML("research/index.html"))
}
