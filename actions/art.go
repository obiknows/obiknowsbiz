package actions

import "github.com/gobuffalo/buffalo"

// ArtIndex default implementation.
func ArtIndex(c buffalo.Context) error {
	return c.Render(200, r.HTML("art/index.html"))
}
