package main

import (
	"context"
	"github.com/gocraft/dbr/v2"
	"github.com/labstack/echo/v4"
)

func PostGresSession(db *dbr.Session) echo.MiddlewareFunc {
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
