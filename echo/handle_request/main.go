package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type (
	User struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
)

var (
	counter = 1
	data    = map[int]*User{
		1: &User{ID: 1, Name: "test1", Email: "test1@gmail.com"},
	}
)

func main() {
	e := echo.New()
	e.Renderer = newRenderEngine()

	e.Any("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, c.QueryParam("a"))
	})

	e.GET("/users", func(c echo.Context) error {
		id, err := strconv.Atoi(c.QueryParam("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		user, ok := data[id]
		if ok && user != nil {
			return c.JSONPretty(
				http.StatusOK,
				echo.Map{
					"code": "0000",
					"data": user,
				}, "\t")
		}

		return c.JSONPretty(
			http.StatusNotFound,
			echo.Map{
				"code": "0001",
				"msg":  fmt.Sprintf("user [%d] not found", id),
			}, "\t")
	})

	e.POST("/users", func(c echo.Context) error {
		name := strings.TrimSpace(c.FormValue("name"))
		email := strings.TrimSpace(c.FormValue("email"))

		if name == "" || email == "" {
			return c.JSONPretty(
				http.StatusBadRequest,
				echo.Map{
					"code": "0002",
					"msg":  "name and email are required",
				}, "\t")
		}

		counter++
		user := new(User)
		user.ID = counter
		user.Name = name
		user.Email = email
		data[user.ID] = user

		return c.JSONPretty(
			http.StatusOK,
			echo.Map{
				"code": "0000",
				"data": user,
			}, "\t")
	})

	e.GET("/user/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		user, ok := data[id]
		if ok && user != nil {
			return c.JSONPretty(
				http.StatusOK,
				echo.Map{
					"code": "0000",
					"data": user,
				}, "\t")
		}

		return c.JSONPretty(
			http.StatusNotFound,
			echo.Map{
				"code": "0001",
				"msg":  fmt.Sprintf("user [%d] not found", id),
			}, "\t")
	})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})

	e.GET("/html", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `<h1>Hello, world</h1>`)
	})

	e.GET("/render", func(c echo.Context) error {
		switch c.QueryParam("tmpl") {
		case "simple":
			return c.Render(http.StatusOK, "simple", data[1])
		case "table":
			return c.Render(http.StatusOK, "table", data[1])
		default:
			return c.JSON(http.StatusOK, data[1])
		}
	})

	e.Logger.Fatal(e.Start(":8080"))
}
