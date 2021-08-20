package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	TZ = "Asia/Tokyo"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	loc, err := time.LoadLocation(TZ)
	if err != nil {
		return events.APIGatewayProxyResponse{}, fmt.Errorf("wrong TZ string: %v", TZ)
	}
	now := time.Now().In(loc)
	year := now.Year()
	month := int(now.Month()) % 12 + 1
	uri := fmt.Sprintf("https://p.secure.freee.co.jp/#work_records/%v/%v/employees/%v", year, month, os.Getenv("employees"))
	return events.APIGatewayProxyResponse{
		Headers: map[string]string{"location": uri},
		Body:       uri,
		StatusCode: http.StatusFound,
	}, nil
}

func main() {
	lambda.Start(handler)
}
