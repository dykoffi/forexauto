package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dykoffi/forexauto/src/config"
	"github.com/dykoffi/forexauto/src/data"
	"github.com/dykoffi/forexauto/src/logger"
)

func main() {
	configService := config.New()
	dataService := data.New()
	loopInterval, _ := strconv.ParseInt(configService.GetOrDefault("LOOP_INTERVAL", "9"), 10, 64)

	for {
		// panic("demain est dimanche")
		fmt.Printf("Test each %d minutes okay \n", loopInterval)
		logger.New().Error("Panic")

		FullForexQuoteData, err := dataService.GetFullForexQuote(data.EURUSD)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(*FullForexQuoteData)

		for _, fd := range *FullForexQuoteData {
			fmt.Println(fd.Volume)
		}

		time.Sleep(time.Duration(loopInterval) * time.Minute)
	}
}
