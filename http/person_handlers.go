package http

import (
	"PeopleService/app"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type httpHandlers struct {
	person_logic app.PersonLogic
}

func NewHttpHandlers(e *echo.Echo, person_logic app.PersonLogic) {
	h := httpHandlers{person_logic: person_logic}
	e.GET("/", h.GetList)
	e.GET("/Person/:Id", h.GetPerson)
	e.POST("/Person", h.AddPerson)
	e.DELETE("/Delete/:Id", h.DeletePerson)
	e.PUT("/Update/:Id", h.UpdatePerson)
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
	err, result := h.person_logic.GetList(limitint, searchstring, offsetint, ctx)
	if err != nil {
		c.JSON(505, err)
	}
	return c.JSON(http.StatusOK, result)
}
func (h *httpHandlers) GetPerson(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		c.JSON(505, err)
	}
	err, result := h.person_logic.GetPerson(id, ctx)
	if err != nil {
		c.JSON(505, err)
	}
	return c.JSON(http.StatusOK, result)
}
func (h *httpHandlers) AddPerson(c echo.Context) error {
	ctx := c.Request().Context()
	Email := c.FormValue("Email")
	Phone := c.FormValue("Phone")
	FirstName := c.FormValue("FirstName")
	LastName := c.FormValue("LastName")
	id, err := strconv.Atoi(c.FormValue("Id"))
	if err != nil {
		c.JSON(505, err)
	}
	err, result := h.person_logic.AddPerson(id, Email, Phone, FirstName, LastName, ctx)
	if err != nil {
		c.JSON(505, err)
	}
	return c.JSON(http.StatusOK, result)
}
func (h *httpHandlers) DeletePerson(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		c.JSON(505, err)
	}
	err, result := h.person_logic.DeletePerson(id, ctx)
	if err != nil {
		c.JSON(505, err)
	}
	return c.JSON(http.StatusOK, result)
}
func (h *httpHandlers) UpdatePerson(c echo.Context) error {
	ctx := c.Request().Context()
	Email := c.FormValue("Email")
	Phone := c.FormValue("Phone")
	FirstName := c.FormValue("FirstName")
	LastName := c.FormValue("LastName")
	id, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		c.JSON(505, err)
	}
	err, result := h.person_logic.UpdatePerson(id, Email, Phone, FirstName, LastName, ctx)
	if err != nil {
		c.JSON(505, err)
	}
	return c.JSON(http.StatusOK, result)
}
