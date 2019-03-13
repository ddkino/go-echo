package main

import (
	"math/rand"
	"net/http"
	"time"

	gohttp "github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := gohttp.New()
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{
			gohttp.HeaderOrigin, 
			gohttp.HeaderContentType, 
			gohttp.HeaderAccept, 
			gohttp.HeaderAuthorization, 
			gohttp.HeaderXCSRFToken,
		},
	}))

	e.Static("/", "public")

	// JSONP
	e.GET("/jsonp", func(c gohttp.Context) error {
		callback := c.QueryParam("callback")
		var content struct {
			Response  string    `json:"response"`
			Timestamp time.Time `json:"timestamp"`
			Random    int       `json:"random"`
		}
		content.Response = "Sent via JSONP"
		content.Timestamp = time.Now().UTC()
		content.Random = rand.Intn(1000)
		return c.JSONP(http.StatusOK, callback, &content)
	})

	// Start server
	e.Logger.Fatal(e.Start(":3333"))
}
