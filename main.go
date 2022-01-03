package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type UserRate struct {
	IdUser       int
	IdRate       int
	Email        string
	UserTimeZone string
	Rate         float64
}

func HandleRequest(ctx context.Context) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{Body: string(getRatesPerUser()), StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
	//fmt.Print(string(getRatesPerUserPerProject()))

}

func getRatesPerUser() []byte {
	var rates []UserRate
	users := getHarvestActiveUsers()
	for _, user := range users {
		idUser := int(user.(map[string]interface{})["id"].(float64))
		email := user.(map[string]interface{})["email"].(string)
		timeZone := user.(map[string]interface{})["timezone"].(string)
		ratesPerUser := getHarvestRatesPerUser(idUser)
		for _, rate := range ratesPerUser {
			idRate := int(rate.(map[string]interface{})["id"].(float64))
			amount := rate.(map[string]interface{})["amount"].(float64)
			if rate.(map[string]interface{})["end_date"] != nil {
				continue
			}
			userRate := UserRate{idUser, idRate, email, timeZone, amount}
			rates = append(rates, userRate)
		}
	}
	jsonRates, _ := json.Marshal(rates)
	return jsonRates
}
