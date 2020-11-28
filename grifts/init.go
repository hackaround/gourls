package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/hackaround/gourls/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
