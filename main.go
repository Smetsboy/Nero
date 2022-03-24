package main

import (
	"PeopleService/http"
	"PeopleService/logic"
	PostBD "PeopleService/postgres"
	"github.com/gocraft/dbr/v2"
	_ "github.com/gocraft/dbr/v2"
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
func ConnServer() (*dbr.Session, error) {
	//connBD := "user = postgres password=74568426 dbname=PersonDB sslmode=disable"
	connBD, err := dbr.Open("postgres", "user = postgres password=74568426 dbname=PersonDB sslmode=disable", nil)
	connBD.SetMaxOpenConns(10)
	sess := connBD.NewSession(nil)
	sess.Begin()
	return sess, err
}
