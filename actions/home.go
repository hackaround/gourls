package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// RoutesHandler is a default handler to serve up
// a home page.
func RoutesHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("index.html"))
}
