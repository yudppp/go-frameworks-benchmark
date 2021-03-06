package main

import (
	"github.com/labstack/echo"
)

func init() {
	calcMem("echo", initEcho)
}

func initEcho() {
	h := echo.New()
	h.Get("/", func(c *echo.Context) error {
		return c.String(200, "Hello, World")
	})
	h.Get("/:name", func(c *echo.Context) error {
		return c.String(200, "Hello, %s", c.Param("name"))
	})
	registerHandler("echo", h)
}
