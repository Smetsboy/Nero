package main

import (
	"context"
	"database/sql"
	"github.com/labstack/echo/v4"
)

func DBServerPostGres(db *sql.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			ctx := req.Context()
			ctx = context.WithValue(ctx, "Postgres-client", db)
			c.SetRequest(req.WithContext(ctx))
			return next(c)
		}
	}
}
