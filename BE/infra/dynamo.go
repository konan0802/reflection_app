package infra

import "reflection_app/domain/repository"

type dynamoDBInfra struct{}

func NewDynamoDBInfra() repository.DynamoDBRepository {
	return &dynamoDBInfra{}
}
