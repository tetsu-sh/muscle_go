package usecase

import (
	"fmt"
	"time"

	"muscle_go/cmd/domain"
)

type TrainUsecase struct {
	TrainRepository domain.TrainRepository
}

func (tu *TrainUsecase) New(r domain.TrainRepository) *TrainUsecase {
	return &TrainUsecase{r}
}

func (tu *TrainUsecase) CreateTrain(name string, volume int, rep int, set int) error {
	date := time.Now()
	train, err := domain.NewTrain(
		name, date, volume, rep, set,
	)
	fmt.Printf("%v", train)
	if err != nil {
		return err
	}
	return nil
}
