package main

import (
	"PeopleService/http"
	"PeopleService/logic"
	DataBase "PeopleService/postgres"
	"github.com/gocraft/dbr/v2"
	_ "github.com/gocraft/dbr/v2"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	db, err := Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	e := echo.New()
	e.Use(PostGresSession(db))
	http.NewHttpHandlers(e, logic.NewLogic(DataBase.NewUrlRepository()))
	e.Logger.Fatal(e.Start(":1323"))
}
func Connection() (*dbr.Session, error) {
	connBD, err := dbr.Open("postgres", "user = postgres password=74568426 dbname=PersonDB sslmode=disable", nil)
	if err != nil {
		log.Fatal(err)
	}
	connBD.SetMaxOpenConns(10)
	sess := connBD.NewSession(nil)
	sess.Begin()
	return sess, err
}
