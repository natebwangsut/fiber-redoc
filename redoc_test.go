package redoc

import (
	"errors"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/stretchr/testify/mock"
	"github.com/swaggo/swag"
)

////////////////////////////////////////////////////////////////////////////////

type mockDocsOKInterface struct {
	mock.Mock
}

type mockDocsErrInterface struct {
	mock.Mock
}

////////////////////////////////////////////////////////////////////////////////

//
func (m *mockDocsOKInterface) ReadDoc() (string, error) {
	return "", nil
}

func (m *mockDocsErrInterface) ReadDoc() (string, error) {
	return "", errors.New("unknown error")
}

////////////////////////////////////////////////////////////////////////////////

func Test_DocsService(t *testing.T) {
	ds := docsService{}

	t.Run("random should OK", func(t *testing.T) {
		swagDoc, swagErr := swag.ReadDoc()
		actualDoc, actualErr := ds.ReadDoc()

		utils.AssertEqual(t, swagDoc, actualDoc)
		utils.AssertEqual(t, swagErr, actualErr)
	})
}

////////////////////////////////////////////////////////////////////////////////

func Test_DocsRedirect(t *testing.T) {
	app := fiber.New()
	app.Get("/docs/*", Handler)
	t.Run("docs should redirect", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/docs", nil)
		res, err := app.Test(req)
		if err != nil {
			t.Fatalf("%s", err)
		}

		//
		if res.StatusCode != http.StatusMovedPermanently {
			t.Fatalf("status code %d should be 301 (Moved Permanently)", res.StatusCode)
		}
	})
}

func Test_DocsTrailingSlashRedirect(t *testing.T) {
	app := fiber.New()
	app.Get("docs/*", Handler)
	t.Run("docs/ should redirect", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/docs/", nil)
		res, err := app.Test(req)
		if err != nil {
			t.Fatalf("%s", err)
		}

		//
		if res.StatusCode != http.StatusMovedPermanently {
			t.Fatalf("status code %d should be 301 (Moved Permanently)", res.StatusCode)
		}
	})
}

func Test_IndexHtml(t *testing.T) {
	app := fiber.New()
	app.Get("/docs/*", Handler)
	t.Run("index.html should OK", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/docs/index.html", nil)
		res, err := app.Test(req)
		if err != nil {
			t.Fatalf("%s", err)
		}

		//
		if res.StatusCode != http.StatusOK {
			t.Fatalf("status code %d should be 200 (OK)", res.StatusCode)
		}
	})
}

func Test_DocsJSON_OK(t *testing.T) {
	app := fiber.New()
	app.Get("/docs/*", Handler)

	docs = &mockDocsOKInterface{}
	t.Run("docs.json should OK", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/docs/docs.json", nil)
		res, err := app.Test(req)
		if err != nil {
			t.Fatalf("%s", err)
		}

		//
		if res.StatusCode != http.StatusOK {
			t.Fatalf("status code %d should be 200 (OK)", res.StatusCode)
		}
	})
}

func Test_DocsJSON_Error(t *testing.T) {
	app := fiber.New()
	app.Get("/docs/*", Handler)

	docs = &mockDocsErrInterface{}
	t.Run("docs.json should Err", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/docs/docs.json", nil)
		res, err := app.Test(req)
		if err != nil {
			t.Fatalf("%s", err)
		}

		//
		if res.StatusCode != http.StatusInternalServerError {
			t.Fatalf("status code %d should be 500 (InternalServerError)", res.StatusCode)
		}
	})
}

func Test_OutOfScope(t *testing.T) {
	app := fiber.New()
	app.Get("/docs/*", Handler)

	docs = &mockDocsErrInterface{}
	t.Run("random should OK", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/docs/random", nil)
		res, err := app.Test(req)
		if err != nil {
			t.Fatalf("%s", err)
		}

		//
		if res.StatusCode != http.StatusOK {
			t.Fatalf("status code %d should be 200 (OK)", res.StatusCode)
		}
	})
}
