package domain

import (
	"time"

	"github.com/google/uuid"
)

type Train struct {
	name   TrainName
	id     TrainId
	date   time.Time
	volume Volume
	rep    Rep
	set    Set
}

func NewTrain(name string, date time.Time, volume int32, rep int32, set int32) (*Train, error) {
	id, _ := uuid.NewUUID()
	train := Train{
		name:   TrainName(name),
		id:     id,
		date:   date,
		volume: Volume(volume),
		rep:    Rep(rep),
		set:    Set(set),
	}
	return &train, nil
}

type TrainRepository interface {
	FetchById() Train
	save(Train)
}

type TrainName string

type TrainId string

type Volume int32

type Rep int32

type Set int32
