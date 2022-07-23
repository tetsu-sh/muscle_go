package usecase

import (
	"muscle/domain/train"
	"time"

	"github.com/labstack/echo"
)

type TrainUsecase struct {
	c echo.Context
	r train.TrainRepository
}

func (tu *TrainUsecase) CreateTrain(name string, volume int32, rep int32, set int32) error {
	date := time.Now()
	train, err := train.NewTrain(
		name, date, volume, rep, set,
	)
	if err != nil {
		return nil, err
	}
	return train, nil
}
