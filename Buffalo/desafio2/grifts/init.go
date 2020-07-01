package grifts

import (
  "github.com/gobuffalo/buffalo"
	"desafio2/actions"
)

func init() {
  buffalo.Grifts(actions.App())
}
