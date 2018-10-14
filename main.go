package main

import (
	"fmt"
	"time"
	"github.com/gwtony/gapi/api"
	"github.com/gstdio/rrb-backend/rrbend"
)

func main() {
	err := api.Init("rrbend.conf")
	if err != nil {
		fmt.Println("[Error] Init api failed")
		return
	}
	config := api.GetConfig()
	log := api.GetLog()

	err = rrbend.InitContext(config, log)
	if err != nil {
		fmt.Println("[Error] Init rrbend failed")
		time.Sleep(time.Second)
		return
	}

	api.Run()
}
