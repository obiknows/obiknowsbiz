package actions

import "github.com/gobuffalo/buffalo"

// IndexHandler is a default handler to serve up the directory page.
func IndexHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("index.html"))
}

// KubeHandler is a default handler to serve
func KubeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("kubed.html"))
}
