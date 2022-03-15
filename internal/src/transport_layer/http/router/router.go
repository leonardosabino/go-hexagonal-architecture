package router

import (
	"hexagonal/template/internal/src/transport_layer/http"

	"github.com/labstack/echo/v4"
)

func Register(server *echo.Echo) {
	server.GET("/v1/dummies", http.ConstructorDummyController().GetDummy)
}
