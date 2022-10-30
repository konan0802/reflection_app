package usecase

import (
	"reflection_app/domain/repository"
)

type Usecase interface {
}

type usecase struct {
	drp repository.DynamoDBRepository
}

func NewUsecase(drp repository.DynamoDBRepository) Usecase {
	return &usecase{
		drp: drp,
	}
}
