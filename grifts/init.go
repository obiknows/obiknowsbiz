package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/obiknows/obiknowsbiz/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
