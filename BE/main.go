package main

import (
	"log"
	"os"
	"reflection_app/handler"
	"reflection_app/infra"
	"reflection_app/usecase"
	"strings"

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
	case "/thread/get":
		res = hdr.GetThread(request)
	case "/thread/put":
		res = hdr.PutThread(request)
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
	if strings.HasPrefix(os.Getenv("AWS_EXECUTION_ENV"), "AWS_Lambda") {
		lambda.Start(Handler)
	} else {
		test := events.APIGatewayProxyRequest{
			Resource: "/reflection/get",
			QueryStringParameters: map[string]string{
				"since": "2022-10-01",
				"until": "2022-10-02",
			},
		}

		_, err := Handler(test)
		if err != nil {
			log.Fatal(err)
		}
	}
}
