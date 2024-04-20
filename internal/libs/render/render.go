package render

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
)

func Render(c context.Context, w http.ResponseWriter, cmp templ.Component) error {
	w.Header().Set("Content-Type", "text/html")
	return cmp.Render(c, w)
}
