package repository

import (
	"muscle_go/cmd/domain"

	"github.com/jinzhu/gorm"
)

type TrainRepositoryImpl struct {
	Conn *gorm.DB
}

func NewTrainRepository(conn *gorm.DB) *TrainRepositoryImpl {
	return &TrainRepositoryImpl{Conn: conn}
}

func (tr *TrainRepositoryImpl) Save(t domain.Train) {
	return
}

func (tr *TrainRepositoryImpl) FetchById(ti domain.TrainId) domain.Train {
	return domain.Train{}
}
