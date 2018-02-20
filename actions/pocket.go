package actions

import "github.com/gobuffalo/buffalo"

// PocketIndex default implementation.
func PocketIndex(c buffalo.Context) error {
	return c.Render(200, r.HTML("pocket/index.html"))
}

// PocketEdit default implementation.
func PocketEdit(c buffalo.Context) error {
	return c.Render(200, r.HTML("pocket/edit.html"))
}
