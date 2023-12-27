package main

import (
	"GroupieTracker/models"
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"html/template"
	"io"
	"net/http"
)

type TemplateRegistry struct {
	templates *template.Template
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	r := echo.New()
	r.Use(middleware.Logger())
	r.Use(middleware.Recover())

	templates := &TemplateRegistry{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}

	r.Renderer = templates

	r.GET("/", func(c echo.Context) error {
		resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
		if err != nil {
			return c.String(http.StatusInternalServerError, "failed to fetch artists")
		}
		if resp.StatusCode != http.StatusOK {
			return c.String(http.StatusInternalServerError, "unexpected status from api")
		}
		defer resp.Body.Close()

		artists := []models.Artist{}
		if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
			return c.String(http.StatusInternalServerError, "failed to decode response")
		}

		return c.Render(http.StatusOK, "index.html", map[string]interface{}{
			"artists": artists,
		})
	})

	r.GET("/artist/:id", func(c echo.Context) error {
		id := c.Param("id")
		resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + id)
		if err != nil {
			return c.String(http.StatusInternalServerError, "failed to fetch artist")
		}
		if resp.StatusCode != http.StatusOK {
			return c.String(http.StatusInternalServerError, "unexpected status from api")
		}
		defer resp.Body.Close()

		artist := new(models.Artist)
		if err := json.NewDecoder(resp.Body).Decode(&artist); err != nil {
			return c.String(http.StatusInternalServerError, "failed to decode response")
		}

		// Make another API call to fetch the locations
		locResp, err := http.Get(artist.Locations)
		if err != nil {
			return c.String(http.StatusInternalServerError, "failed to fetch locations")
		}
		defer locResp.Body.Close()

		locations := new(models.Location)
		if err := json.NewDecoder(locResp.Body).Decode(locations); err != nil {
			return c.String(http.StatusInternalServerError, "failed to decode locations")
		}

		// And another to fetch concert dates
		dateResp, err := http.Get(artist.ConcertDates)
		if err != nil {
			return c.String(http.StatusInternalServerError, "failed to fetch concertDates")
		}
		defer dateResp.Body.Close()

		concertDates := new(models.Date)
		if err := json.NewDecoder(dateResp.Body).Decode(concertDates); err != nil {
			return c.String(http.StatusInternalServerError, "failed to decode concertDates")
		}

		// Now we return a new anonymous struct which includes the locations and dates
		return c.Render(http.StatusOK, "artist.html", struct {
			*models.Artist
			Locations    []string
			ConcertDates []string
		}{
			Artist:       artist,
			Locations:    locations.Locations,
			ConcertDates: concertDates.Dates,
		})
	})

	r.Logger.Fatal(r.Start(":3000"))
}
