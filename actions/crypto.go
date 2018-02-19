package actions

import "github.com/gobuffalo/buffalo"

// CryptoIndex default implementation.
func CryptoIndex(c buffalo.Context) error {
	return c.Render(200, r.HTML("crypto/index.html"))
}
