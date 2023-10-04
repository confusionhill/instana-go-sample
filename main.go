package main

import (
	"coffe-shop/database"
	"coffe-shop/services"
	auth "coffe-shop/services/authentication"
	"coffe-shop/services/guard"
	places2 "coffe-shop/services/places"
	"github.com/golang-jwt/jwt/v5"
	instana "github.com/instana/go-sensor"
	"github.com/instana/go-sensor/instrumentation/instaecho"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {
	database.LoadDatabase()

	opt := *instana.DefaultOptions()
	opt.Service = "fake-coffe-shop"
	opt.EnableAutoProfile = true
	instana.StartMetrics(&opt)

	// initialize and configure the logger
	logger := logrus.New()
	logger.Level = logrus.InfoLevel

	// check if INSTANA_DEBUG is set and set the log level to DEBUG if needed
	if _, ok := os.LookupEnv("INSTANA_DEBUG"); ok {
		logger.Level = logrus.DebugLevel
	}

	// use logrus to log the Instana Go Collector messages
	instana.SetLogger(logger)
	sensor := instana.NewSensor("fake-coffe-shop")
	cfg := services.GetConfig()
	e := instaecho.New(sensor)
	// auth
	e.POST("/login", auth.LoginUser)
	e.POST("/register", auth.RegisterUser)

	// test
	e.GET("/", func(c echo.Context) error {
		logger.Error("what the hell!")
		return c.String(http.StatusOK, "Hello, World!")
	})
	places := e.Group("/place")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(services.CustomClaims)
		},
		SigningKey: []byte(cfg.JwtSecret),
	}
	places.GET("/:id", places2.GetPlaceByID)
	places.GET("", places2.GetPlaces)
	places.POST("", guard.OwnerGuard(places2.InsertPlace), echojwt.WithConfig(config))
	places.PUT("", guard.AdminGuard(places2.UpdatePlace), echojwt.WithConfig(config))
	places.DELETE("", guard.OwnerGuard(places2.DeletePlace), echojwt.WithConfig(config))

	// comment
	places.GET("/comment/:id", places2.GetCommentByPlaceID)
	places.POST("/comment", guard.AuthGuard(places2.AddComment), echojwt.WithConfig(config))

	e.Logger.Fatal(e.Start(":1323"))
}
