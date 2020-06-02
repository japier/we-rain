package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type timezone struct {
	Dt        int64   `json:"dt"`
	Sunrise   int64   `json:"sunrise"`
	Sunset    int64   `json:"sunset"`
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
	DewPoint  float64 `json:"dew_point"`
}

type response struct {
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	Timezone       string  `json:"timezone"`
	TimezoneOffset string  `json:"timezone_offset"`
}

/*
{
"lat": 35.69,
"lon": 139.69,
"timezone": "Asia/Tokyo",
"timezone_offset": 32400,
"current":{
"dt": 1590415791,
"sunrise": 1590348592,
"sunset": 1590400008,
"temp": 293.91,
"feels_like": 291.83,
"pressure": 1009,
"humidity": 82,
"dew_point": 290.73,
"uvi": 9.39,
"clouds": 75,
"visibility": 10000,
"wind_speed": 6.7,
"wind_deg": 190,
"weather":[
{
"id": 803,
"main": "Clouds",
"description": "broken clouds",
"icon": "04n"
}
]
},
"hourly":[
*/

func run(request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello World",
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		},
	}, nil
}
func main() {
	lambda.Start(run)
}
