package actions

import "github.com/gobuffalo/buffalo"

// IpfsIndex default implementation.
func IpfsIndex(c buffalo.Context) error {
	return c.Render(200, r.HTML("ipfs/index.html"))
}
