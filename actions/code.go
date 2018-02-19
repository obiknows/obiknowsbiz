package actions

import "github.com/gobuffalo/buffalo"

// CodeIndex default implementation.
func CodeIndex(c buffalo.Context) error {
	return c.Render(200, r.HTML("code/index.html"))
}
