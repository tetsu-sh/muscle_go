package presentation

import (
	"net/http"

	"muscle_go/cmd/domain"
	"muscle_go/cmd/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// handler
func (h *Handler) show(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}

type Handler struct {
	TrainRepository domain.TrainRepository
	TrainUsecase    usecase.TrainUsecase
}

func (h Handler) Register(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", h.show)
	e.POST("/train", h.CreateTrain)
}

func NewHandler(tr domain.TrainRepository) *Handler {
	return &Handler{TrainRepository: tr}
}
