# Fiber Redoc

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

Made with :heart: of `Swagger` and `React`.
