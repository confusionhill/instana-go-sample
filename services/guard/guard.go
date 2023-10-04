package guard

import (
	"coffe-shop/services"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthGuardRequest struct {
	EchoCtx echo.Context
	Claims  *services.CustomClaims
}

func AuthGuard(funcHandler func(a AuthGuardRequest) error) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		if user == nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		claims := user.Claims.(*services.CustomClaims)
		jwt := AuthGuardRequest{
			EchoCtx: c,
			Claims:  claims,
		}
		return funcHandler(jwt)
	}
}

func AdminGuard(funcHandler func(a AuthGuardRequest) error) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		if user == nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		claims := user.Claims.(*services.CustomClaims)
		if claims.UserRole < 5 {
			return c.String(http.StatusUnauthorized, "unauthorized : user does not have admin access")
		}
		jwt := AuthGuardRequest{
			EchoCtx: c,
			Claims:  claims,
		}
		return funcHandler(jwt)
	}
}

func OwnerGuard(funcHandler func(a AuthGuardRequest) error) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		if user == nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		claims := user.Claims.(*services.CustomClaims)
		if claims.UserRole < 10 {
			return c.String(http.StatusUnauthorized, "unauthorized : user does not have owner access")
		}
		jwt := AuthGuardRequest{
			EchoCtx: c,
			Claims:  claims,
		}
		return funcHandler(jwt)
	}
}
