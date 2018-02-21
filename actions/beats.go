package actions

import "github.com/gobuffalo/buffalo"

// BeatsIndex default implementation.
func BeatsIndex(c buffalo.Context) error {
	return c.Render(200, r.HTML("sounds/index.html"))
}
