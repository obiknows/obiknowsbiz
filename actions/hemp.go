package actions

import "github.com/gobuffalo/buffalo"

// HempIndex default implementation.
func HempIndex(c buffalo.Context) error {
	return c.Render(200, r.HTML("hemp/index.html"))
}
