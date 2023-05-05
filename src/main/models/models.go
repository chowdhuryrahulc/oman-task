package models

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	// log "github.com/sirupsen/logrus"
	"github.com/chowdhuryrahulc/oman-task/src/main/util"
)


type IExchangeApiResponseModel interface {
	Get() []byte
}

type ExchangeAPIResponse struct {
	Success    bool   `json:"success"`
	Timeseries bool   `json:"timeseries"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
	Base       string `json:"base"`
	Rates      Rate   `json:"rates"`
}

type Rate struct {
	Dates Currency
}

type Currency struct {
	USD float64 `json:"USD"`
	EUR float64 `json:"EUR"`
	GBP float64 `json:"GBP"`
}

func (exe ExchangeAPIResponse) Get() []byte {
	client := &http.Client{}
	req, err := util.CreateHTTPRequest()
	if err != nil {
		fmt.Println("error creating exchange api url request: " + err.Error())
	}

	// Timeout added of 3 sec
	ctx, cancel := context.WithTimeout(context.Background(), 9 * time.Second)
	defer cancel()
	res, err := client.Do(req.WithContext(ctx))
	if err != nil {
		fmt.Println("error fetching exchange api url response: " + err.Error())
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error reading exchange api url response: " + err.Error())
	}

	fmt.Println("Body", string(body))
	return body
	// byte := util.ReturnExchangeApiResponse()
	// util.MarshalExchangeApiResponseToStruct(byte)
}
