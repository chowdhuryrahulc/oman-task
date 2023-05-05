package util

import (
	"net/http"
)

func GetTodayDate() string {
	return "2022-05-05"
}

func Get10DayBackDate() string {
	return "2022-05-01"
}

func CreateHTTPRequest()(*http.Request, error){
	url := "https://api.apilayer.com/exchangerates_data/timeseries?base=INR&symbols=USD,EUR,GBP&start_date=2022-05-01&end_date=2022-05-05"

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("apikey", "0RhIgK5LbHXiIWnGJiUVAnnvbH2aihBS")	// auth token has to come from config || query params set bu above funcs
	// req.Header.Set("base", "INR")
	// req.Header.Set("symbols", "USD,EUR,GBP")
	// req.Header.Set("start_date", "2022-05-01")
	// req.Header.Set("end_date", "2022-05-05")
	return req, err
}

