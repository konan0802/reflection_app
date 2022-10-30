package handler

import (
	"fmt"
	"reflection_app/usecase"

	"github.com/aws/aws-lambda-go/events"
)

type Handler interface {
	GetReflection(events.APIGatewayProxyRequest) []byte
	PutReflection(events.APIGatewayProxyRequest) []byte
}

type handler struct {
	usc usecase.Usecase
}

func NewHandler(usc usecase.Usecase) Handler {
	return &handler{
		usc: usc,
	}
}

func (hdr handler) GetReflection(request events.APIGatewayProxyRequest) []byte {
	test2 := request.QueryStringParameters["test"]
	fmt.Printf("test2: %v\n", test2)

	b := []byte("GetReflectionTest " + test2 + "??")
	return b
}

func (hdr handler) PutReflection(request events.APIGatewayProxyRequest) []byte {
	b := []byte("PutReflectionTest")
	return b
}
