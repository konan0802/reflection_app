package handler

import (
	"encoding/json"
	"fmt"
	"reflection_app/domain/model"
	"reflection_app/usecase"
	"time"

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
	param := request.QueryStringParameters
	since, _ := time.Parse("2006/01/02", param["since"])
	until, _ := time.Parse("2006/01/02", param["until"])

	requestGetReflection := model.RequestGetReflection{
		Since: since,
		Until: until,
	}
	result := hdr.usc.GetReflection(requestGetReflection)
	fmt.Printf("%+v\n", result)
	s, _ := json.Marshal(result)
	return s
}

func (hdr handler) PutReflection(request events.APIGatewayProxyRequest) []byte {
	b := []byte("PutReflectionTest")
	return b
}
