package repository

import (
	"muscle/domain/train"

	"github.com/jinzhu/gorm"
)

type TrainRepository struct {
	Conn *gorm.DB
}

func NewTrainRepository(conn *gorm.DB) *TrainRepository {
	return &TrainRepository{Conn: conn}
}

func (tr TrainRepository) save(t train.Train) error {
	return nil
}
