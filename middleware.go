package main

import (
	"context"
	"database/sql"
	"github.com/labstack/echo/v4"
)

func DBServerPostGres() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			ctx := req.Context()
			ctx = context.WithValue(ctx, "Postgres-client", ConnServer())
			c.SetRequest(req.WithContext(ctx))
			return next(c)
		}
	}
}
func ConnServer() *sql.DB {
	connBD := "user = postgres password=74568426 dbname=PersonDB sslmode=disable"
	db, err := sql.Open("postgres", connBD)
	if err != nil {
		panic(err)
	}
	return db
}
