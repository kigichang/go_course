package main

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type (
	CustomValidator struct {
		validator *validator.Validate
	}

	UserID struct {
		ID int `json:"id" param:"id" query:"id" form:"id" validate:"gt=0"`
	}

	User struct {
		ID    int    `json:"id"`
		Name  string `json:"name" query:"name" form:"name"`
		Email string `json:"email" query:"email" form:"email"`
	}
)

var (
	counter = 1
	data    = map[int]*User{
		1: &User{
			ID:    1,
			Name:  "test 1",
			Email: "test1@gmail.com",
		},
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	g := e.Group("/api/v1")

	g.GET("/users", func(c echo.Context) error {
		u := new(UserID)
		if err := c.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		users := []*User{}

		for _, v := range data {
			if u.ID > 0 && v.ID == u.ID {
				users = append(users, v)
			} else if u.ID <= 0 {
				users = append(users, v)
			}
		}
		return c.JSONPretty(http.StatusOK, echo.Map{
			"code": "0000",
			"data": users,
		}, "\t")
	})

	g.POST("/users", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		counter += 1
		u.ID = counter
		data[u.ID] = u
		c.Response().Header().Set("Location", fmt.Sprintf("/api/v1/user/%d", u.ID))
		return c.String(http.StatusCreated, "")
	})

	g.GET("/user/:id", func(c echo.Context) error {
		u := new(UserID)
		if err := c.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		if err := c.Validate(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		ret, ok := data[u.ID]
		if ok && ret != nil {
			return c.JSONPretty(
				http.StatusOK,
				echo.Map{
					"code": "0000",
					"data": ret,
				}, "\t")
		}
		return c.JSONPretty(
			http.StatusNotFound,
			echo.Map{
				"code": "0001",
				"msg":  fmt.Sprintf("user [%d] not found", u.ID),
			}, "\t")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
