package db

import (
	"log"

	"github.com/prometheus/prometheus/tsdb"
)

var tsdbClient *tsdb.DB
var writer *TSDBWriter

func init() {
	var err error
	tsdbClient, err = tsdb.Open("./data.db", nil, nil, tsdb.DefaultOptions(), nil)
	if err != nil {
		log.Fatal(err)
	}
	writer = NewTSDBWriter(tsdbClient, 1000)
}

func CloseDBClient() {
	tsdbClient.Close()
}
