package actions

import "github.com/gobuffalo/buffalo"

// IndexHandler is a default handler to serve up the directory page.
func IndexHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("index.html"))
}
