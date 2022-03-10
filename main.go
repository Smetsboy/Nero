package main

import (
	"PeopleService/http"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	e.Use(DBServerPostGres())
	e.GET("/", http.GetList)
	e.GET("/Person/:Id", http.GetPerson)
	e.POST("/Person", http.AddPerson)
	e.DELETE("/Delete/:Id", http.DeletePerson)
	e.PUT("/Update/:Id", http.UpdatePerson)
	e.Logger.Fatal(e.Start(":1323"))

}
