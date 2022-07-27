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

func NewTrain(name string, date time.Time, volume int, rep int, set int) (*Train, error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	id := uuid.String()
	train := Train{
		name:   TrainName(name),
		id:     TrainId(id),
		date:   date,
		volume: Volume(volume),
		rep:    Rep(rep),
		set:    Set(set),
	}
	return &train, nil
}

type TrainRepository interface {
	Save(t Train)
	FetchById(ti TrainId) Train
}

type TrainName string

type TrainId string

type Volume int32

type Rep int32

type Set int32
