package main

import (
	"github.com/chowdhuryrahulc/oman-task/src/main/app"
	"github.com/chowdhuryrahulc/oman-task/src/main/config"
	"github.com/chowdhuryrahulc/oman-task/src/main/util"
)

func init() {

}

func main() {
  byte := util.ReturnExchangeApiResponse()      // delete later
  util.MarshalExchangeApiResponseToStruct(byte) // delete later
	app.Start(config.NewAppConfig())
}
