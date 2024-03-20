package main

import (
    "html/template"
    "fmt"
    "io"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)


type Templates struct {
    templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplates() *Templates {
    return &Templates {
        templates: template.Must(template.ParseGlob("views/*.html")),
    }
}




func main() {
    fmt.Println("Starting server...")
    e := echo.New()
    e.Renderer = NewTemplates()
    e.Use(middleware.Logger())
    e.Static("/css", "css")
    e.Static("/images", "images")

    e.GET("/", func(c echo.Context) error {
        return c.Render(200, "index", nil)
    })

    e.Logger.Fatal(e.Start(":42069"))
}












