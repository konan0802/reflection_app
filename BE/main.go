package main

import (
	"reflection_app/handler"
	"reflection_app/infra"
	"reflection_app/usecase"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// 依存関係を注入
	dInf := infra.NewDynamoDBInfra()
	usc := usecase.NewUsecase(dInf)
	hdr := handler.NewHandler(usc)

	// ルーティングの設定
	var res []byte
	switch request.Resource {
	case "reflection/get":
		res = hdr.GetReflection(request)
	case "reflection/put":
		res = hdr.PutReflection(request)
	default:
		return events.APIGatewayProxyResponse{
			Body:       `Error: 404 not found`,
			StatusCode: 404,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       string(res),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
