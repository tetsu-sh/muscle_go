package presentation

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/muscle_go/domain"
	"github.com/muscle_go/usecase"
)

// handler
func (h *Handler) show(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}
func (h *Handler) CreateTrain(c echo.Context) error {
	usecase.TrainUsecase.CreateTrain()

}

type Handler struct {
	TrainRepository domain.TrainRepository
}

func (h Handler) Register(e *echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", h.show)
	e.POST("/train", h.CreateTrain)
}

func NewHandler(tr domain.TrainRepository) *Handler {
	return &Handler{TrainRepository: tr}
}
