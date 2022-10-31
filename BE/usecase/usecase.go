package usecase

import (
	"reflection_app/domain/model"
	"reflection_app/domain/repository"
)

type Usecase interface {
	GetReflection(model.RequestGetReflection) model.ResponseGetReflection
}

type usecase struct {
	drp repository.DynamoDBRepository
}

func NewUsecase(drp repository.DynamoDBRepository) Usecase {
	return &usecase{
		drp: drp,
	}
}

func (usc usecase) GetReflection(req model.RequestGetReflection) model.ResponseGetReflection {
	reflection := model.Reflection{
		Type: "a",
		Text: "aiueo",
	}
	reflections := []model.Reflection{reflection}

	daysReflections := model.DaysReflections{
		Date:        "2022-10-31",
		Reflections: reflections,
	}
	data := []model.DaysReflections{daysReflections}

	return model.ResponseGetReflection{
		Data: data,
	}
}
