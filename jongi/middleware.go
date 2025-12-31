package jongi

import (
	"context"
	"net/http"
	"slices"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sugandasu/ruru/tolo"
)

func EchoAuthMiddleware(secret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing authorization header")
			}
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			token, err := ValidateToken(tokenString, secret)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid or expired token", err.Error())
			}

			auth, ok := token.Claims.(*AuthClaims)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid token claims")
			}

			c.SetRequest(c.Request().WithContext(context.WithValue(c.Request().Context(), AuthContext{}, auth)))

			return next(c)
		}
	}
}

func EchoRoleMiddleware(secret string, levels []int) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth := GetAuthFromContext(c.Request().Context())
			if auth == nil {
				return echo.NewHTTPError(http.StatusForbidden, "unauthorized")
			}

			if slices.Contains(levels, auth.Role.Level) {
				return next(c)
			}

			return echo.NewHTTPError(http.StatusForbidden, "unauthorized")
		}
	}
}

func EchoErrorMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err == nil {
			return nil
		}

		tolo.ResponseError(c.Response().Writer, err)
		return nil
	}
}
