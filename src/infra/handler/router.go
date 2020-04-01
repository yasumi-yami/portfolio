package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const indent = "  "

// Router ルーティング
func Router() *echo.Echo {
	e := echo.New()

	e.HideBanner = true

	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.JSONPretty(
			http.StatusOK,
			map[string]interface{}{"text": "hello yami world"},
			indent,
		)
	})

	podcastAPI := &MyPodcastAPI{}
	podcast := e.Group("/podcast")
	podcast.GET("/template", podcastAPI.GetTemplate)
	return e
}
