package models

import (
	"context"
	"io/ioutil"
	"log"

	"cloud.google.com/go/bigquery"
	"golang.org/x/oauth2/google"
	bq "google.golang.org/api/bigquery/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var PROJECT_ID string
var DATASET string
var GOOGLE_APPLICATION_CREDENTIALS_FILE string

var Client *bigquery.Client
var BqCtx context.Context

func ConnectBq(projectId string, dataset string, googlecredentialfile string) {
	BqCtx = context.Background()
	data, err := ioutil.ReadFile(googlecredentialfile)
	if err != nil {
		log.Fatalf("Open file error, %s", err.Error())
	}
	jwtConfig, err := google.JWTConfigFromJSON(data, bq.BigqueryScope)

	if err != nil {
		log.Fatalf("JWT Config bq error, %s", err.Error())
	}

	ts := jwtConfig.TokenSource(BqCtx)

	Client, err = bigquery.NewClient(BqCtx, PROJECT_ID, option.WithTokenSource(ts))

	if err != nil {
		log.Fatalf("Client not opened")
	}
}

type RowType struct {
	Forecast_hour          int
	Forecasted_temperature float32
	Forecasted_uv          float32
	City                   string
}

func GetQuery(query string) []RowType {
	q := Client.Query(query)

	it, err := q.Read(BqCtx)
	if err != nil {
		log.Fatalf("Error Query : %s", err.Error())
	}
	var valuesTotal []RowType
	for {

		var values RowType
		err := it.Next(&values)
		if err == iterator.Done {
			break
		}
		valuesTotal = append(valuesTotal, values)

		if err != nil {
			log.Fatalf("Error Value : %s", err.Error())
		}

	}
	return valuesTotal
}
