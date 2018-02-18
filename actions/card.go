package actions

import "github.com/gobuffalo/buffalo"

// CardIndex default implementation.
func CardIndex(c buffalo.Context) error {
	return c.Render(200, r.HTML("card/index.html"))
}
