# Fiber Redoc

[![Go Document](https://pkg.go.dev/badge/github.com/natebwangsut/fiber-redoc?utm_source=godoc)](https://pkg.go.dev/github.com/natebwangsut/fiber-redoc)
[![Go Coverage](https://img.shields.io/badge/go%20coverage-here-brightgreen)](https://gocover.io/github.com/natebwangsut/fiber-redoc)
[![Go Report Card](https://goreportcard.com/badge/github.com/natebwangsut/fiber-redoc)](https://goreportcard.com/report/github.com/natebwangsut/fiber-redoc)

[gofiber](https://github.com/gofiber/fiber) middleware for rendering [Redoc](https://github.com/Redocly/redoc). Compatible with Swagger notation (OpenAPI specification).

## Usage

> Use with fiber v2+, not tested with v1

```go
func main() {
	app := fiber.New()
	app.Get("/docs/*", redoc.Handler)
}
```

## Example

Please see example folder on this repository for simple setup.

## Configuration

| Config      | Default Value | Description |
| ----------- | ------------- | ----------- |
| To be added | To be added   | To be added |

---

Made with :heart: of `swagger`, `fiber` and `react.js`.
