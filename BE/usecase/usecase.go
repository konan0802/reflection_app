package usecase

import (
	"reflection_app/domain/model"
	"reflection_app/domain/repository"
)

type Usecase interface {
	GetThread(model.RequestGetThread) model.ResponseGetThread
}

type usecase struct {
	drp repository.DynamoDBRepository
}

func NewUsecase(drp repository.DynamoDBRepository) Usecase {
	return &usecase{
		drp: drp,
	}
}

func (usc usecase) GetThread(req model.RequestGetThread) model.ResponseGetThread {
	reflection := model.Refl{
		Type: "a",
		Text: "aiueo",
	}
	reflections := []model.Refl{reflection}

	daysThread := model.Thread{
		Date:  "2022-10-31",
		Refls: reflections,
	}
	data := []model.Thread{daysThread}

	return model.ResponseGetThread{
		Data: data,
	}
}
