package http

import (
	"PeopleService/logic"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func GetList(c echo.Context) error {
	ctx := c.Request().Context()
	limit := c.QueryParam("limit")
	search := c.QueryParam("search")
	offset := c.QueryParam("offset")
	var limitint int

	if limit != "" {
		limitint1, err := strconv.Atoi(limit)
		if err != nil {
			panic(err)
		}
		limitint = limitint1
	}
	searchstring := search
	var offsetint int

	if offset != "" {
		offsetint1, err := strconv.Atoi(offset)
		if err != nil {
			panic(err)
		}
		offsetint = offsetint1
	}

	return c.JSON(http.StatusOK, logic.GetList(limitint, searchstring, offsetint, ctx))
}
func GetPerson(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		panic(err)
	}
	person := logic.GetPerson(id, ctx)
	return c.JSON(http.StatusOK, person)
}
func AddPerson(c echo.Context) error {
	ctx := c.Request().Context()
	Id := c.FormValue("Id")
	Email := c.FormValue("Email")
	Phone := c.FormValue("Phone")
	FirstName := c.FormValue("FirstName")
	LastName := c.FormValue("LastName")
	id, err := strconv.Atoi(c.FormValue("Id"))
	if err != nil {
		panic(err)
	}
	logic.AddPerson(id, Email, Phone, FirstName, LastName, ctx)
	return c.JSON(http.StatusOK, "Id:"+Id+", Email:"+Email+", Phone:"+Phone+", FirstName:"+FirstName+", LastName:"+LastName)
}
func DeletePerson(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		panic(err)
	}
	person := logic.DeletePerson(id, ctx)
	return c.JSON(http.StatusOK, person)
}
func UpdatePerson(c echo.Context) error {
	ctx := c.Request().Context()
	Email := c.FormValue("Email")
	Phone := c.FormValue("Phone")
	FirstName := c.FormValue("FirstName")
	LastName := c.FormValue("LastName")
	id, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		panic(err)
	}
	person := logic.UpdatePerson(id, Email, Phone, FirstName, LastName, ctx)
	return c.JSON(http.StatusOK, person)
}
