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
	GetThread(events.APIGatewayProxyRequest) []byte
	PutThread(events.APIGatewayProxyRequest) []byte
}

type handler struct {
	usc usecase.Usecase
}

func NewHandler(usc usecase.Usecase) Handler {
	return &handler{
		usc: usc,
	}
}

func (hdr handler) GetThread(request events.APIGatewayProxyRequest) []byte {
	param := request.QueryStringParameters
	since, _ := time.Parse("2006/01/02", param["since"])
	until, _ := time.Parse("2006/01/02", param["until"])

	requestGetThread := model.RequestGetThread{
		Since: since,
		Until: until,
	}
	result := hdr.usc.GetThread(requestGetThread)
	fmt.Printf("%+v\n", result)
	s, _ := json.Marshal(result)
	return s
}

func (hdr handler) PutThread(request events.APIGatewayProxyRequest) []byte {
	b := []byte("PutThreadTest")
	return b
}
