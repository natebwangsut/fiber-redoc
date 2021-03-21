// Package redoc swagger
package redoc

import (
	"path"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/utils"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/swag"
)

var (
	prefix               = ""
	fs     fiber.Handler = filesystem.New(filesystem.Config{Root: swaggerFiles.HTTP})
)

const (
	docPath   = "docs.json"
	indexPath = "index.html"
)

// Handler registers "/index.html" and "/docs.json" endpoint as a form of fiber.Handler
var Handler = New()

// New returns custom handler where it sends redoc HTML page with swagger docs.json if it's registered
func New() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var p string
		if p = utils.ImmutableString(c.Params("*")); p != "" {
			c.Path(p)
		} else {
			p = strings.TrimPrefix(c.Path(), prefix)
			p = strings.TrimPrefix(p, "/")
		}

		switch p {
		// Serve HTML page
		case indexPath:
			return c.Type("html").SendString(redocIndex)
		// Serve JSON
		case docPath:
			doc, err := swag.ReadDoc()
			if err != nil {
				return err
			}
			return c.Type("json").SendString(doc)
		// Redirect /docs/ and /docs to /docs/index.html
		case "", "/":
			return c.Redirect(path.Join(prefix, indexPath), fiber.StatusMovedPermanently)
		}
		return nil
	}
}

////////////////////////////////////////////////////////////////////////////////

const redocIndex = `
<!DOCTYPE html>
<html>
  <head>
    <title>ReDoc</title>
    <!-- needed for adaptive design -->
    <meta charset="utf-8"/>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link href="https://fonts.googleapis.com/css?family=Montserrat:300,400,700|Roboto:300,400,700" rel="stylesheet">

    <!--
    ReDoc doesn't change outer page styles
    -->
    <style>
      body {
        margin: 0;
        padding: 0;
      }
    </style>
  </head>
  <body>
    <redoc spec-url='docs.json'></redoc>
    <script src="https://cdn.jsdelivr.net/npm/redoc@next/bundles/redoc.standalone.js"> </script>
  </body>
</html>`
