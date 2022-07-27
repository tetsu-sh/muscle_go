package presentation

import "github.com/labstack/echo/v4"

func (h *Handler) CreateTrain(c echo.Context) error {
	rec := &trainCreateRequest{}
	rec.bind(c)
	tu := h.TrainUsecase.New(h.TrainRepository)
	tu.CreateTrain(rec.Train.TrainName, rec.Train.Volume, rec.Train.Rep, rec.Train.Set)
	return nil
}
