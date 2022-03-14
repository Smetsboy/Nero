package main

import (
	"PeopleService/http"
	"PeopleService/logic"
	PostBD "PeopleService/postgres"
	"database/sql"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	db, err := ConnServer()

	if err != nil {
		return
	}
	defer db.Close()
	e := echo.New()
	e.Use(DBServerPostGres(db))
	http.NewHttpHandlers(e, logic.NewPersonLogic(PostBD.NewUrlRepository()))
	e.Logger.Fatal(e.Start(":1323"))
}
func ConnServer() (*sql.DB, error) {
	connBD := "user = postgres password=74568426 dbname=PersonDB sslmode=disable"
	db, err := sql.Open("postgres", connBD)
	if err != nil {
		return nil, err
	}
	return db, nil
}
