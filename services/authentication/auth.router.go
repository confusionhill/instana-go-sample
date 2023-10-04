package authentication

import (
	"coffe-shop/services"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"time"
)

var accessMap = map[int]string{
	0:  "default_user",
	5:  "admin",
	10: "owner",
}

func generateAccessToken(user *UserEntity) (string, error) {
	cfg := services.GetConfig()
	claims := services.CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "backend",
			ExpiresAt: &jwt.NumericDate{
				// make it 1 month
				Time: time.Now().Add(365 * 24 * time.Hour),
			},
			IssuedAt: &jwt.NumericDate{
				Time: time.Now(),
			},
		},
		UserID:   user.Id,
		UserRole: user.Access,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JwtSecret))
}

func LoginUser(c echo.Context) error {
	req := LoginUserEntity{}
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	resp, err := PerformLoginUser(req)
	if err != nil {
		log.Info(err.Error())
		return c.String(http.StatusBadRequest, "cannot login user")
	}
	resp.AccessName = accessMap[resp.Access]
	token, err := generateAccessToken(resp)
	if err != nil {
		return c.String(http.StatusInternalServerError, "cannot generate access token")
	}
	return c.JSON(http.StatusOK, services.ResponseDTO{
		ResponseCode: http.StatusOK,
		Message:      "ok",
		Data: LoginUserResponseEntity{
			UserInfo:    resp,
			AccessToken: token,
		},
	})
}

func RegisterUser(c echo.Context) error {
	req := RegisterUserEntity{}
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	role, ok := accessMap[req.RoleID]
	if !ok {
		return c.String(http.StatusUnauthorized, "role is not available")
	}
	resp, err := RegisterNewUser(req)
	if err != nil {
		log.Info(err.Error())
		return c.String(http.StatusBadRequest, "cannot register new user")
	}
	resp.AccessName = role
	token, err := generateAccessToken(resp)
	if err != nil {
		return c.String(http.StatusInternalServerError, "cannot generate access token")
	}
	return c.JSON(http.StatusOK, services.ResponseDTO{
		ResponseCode: http.StatusOK,
		Message:      "ok",
		Data: LoginUserResponseEntity{
			UserInfo:    resp,
			AccessToken: token,
		},
	})
}
