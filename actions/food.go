package actions

import "github.com/gobuffalo/buffalo"

// FoodIndex default implementation.
func FoodIndex(c buffalo.Context) error {
	return c.Render(200, r.HTML("food/index.html"))
}
