package web

import "github.com/gin-contrib/multitemplate"

const (
	base = "web/templates/layouts/base.html"
)

func NewRenderer() multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	// Full pages (base layout + page content)
	r.AddFromFiles("categories/index",
		base,
		"web/templates/categories/index.html",
	)

	// Fragments (HTMX partial responses — no layout)
	r.AddFromFiles("categories/row",
		"web/templates/categories/row.html",
	)

	return r
}
