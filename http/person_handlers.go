package http

import (
	"PeopleService/app"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type httpHandlers struct {
	person_logic app.Logic
}

func NewHttpHandlers(e *echo.Echo, person_logic app.Logic) {
	h := httpHandlers{person_logic: person_logic}
	e.GET("/", h.GetList)
	e.GET("/Person/:Id", h.Get)
	e.POST("/Person", h.Add)
	e.DELETE("/Delete/:Id", h.Delete)
	e.PUT("/Update/:Id", h.Update)
}
func (h *httpHandlers) GetList(c echo.Context) error {
	ctx := c.Request().Context()
	limit := c.QueryParam("limit")
	search := c.QueryParam("search")
	offset := c.QueryParam("offset")
	var limitint int

	if limit != "" {
		limitint1, err := strconv.Atoi(limit)
		if err != nil {
			c.JSON(505, err)
		}
		limitint = limitint1
	}
	searchstring := search
	var offsetint int

	if offset != "" {
		offsetint1, err := strconv.Atoi(offset)
		if err != nil {
			c.JSON(505, err)
		}
		offsetint = offsetint1
	}
	result, err := h.person_logic.GetList(ctx, limitint, searchstring, offsetint)
	if err != nil {
		c.JSON(505, err)
	}
	return c.JSON(http.StatusOK, result)
}
func (h *httpHandlers) Get(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		c.JSON(505, err)
	}
	result, err := h.person_logic.Get(ctx, id)
	if err != nil {
		c.JSON(505, err)
	}
	return c.JSON(http.StatusOK, result)
}
func (h *httpHandlers) Add(c echo.Context) error {
	ctx := c.Request().Context()
	person := app.Person{Email: c.FormValue("Email"), Phone: c.FormValue("Phone"), FirstName: c.FormValue("FirstName"),
		LastName: c.FormValue("LastName")}
	result, err := h.person_logic.Add(ctx, person)
	if err != nil {
		c.JSON(505, err)
	}
	return c.JSON(http.StatusOK, result)
}
func (h *httpHandlers) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		c.JSON(505, err)
	}
	err = h.person_logic.Delete(ctx, id)
	if err != nil {
		c.JSON(505, err)
	}
	return c.JSON(http.StatusOK, "Удалено")
}
func (h *httpHandlers) Update(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		c.JSON(505, err)
	}
	person := app.Person{Id: id, Email: c.FormValue("Email"), Phone: c.FormValue("Phone"), FirstName: c.FormValue("FirstName"),
		LastName: c.FormValue("LastName")}
	result, err := h.person_logic.Update(ctx, person)
	if err != nil {
		c.JSON(505, err)
	}
	return c.JSON(http.StatusOK, result)
}
