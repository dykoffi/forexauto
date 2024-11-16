package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dykoffi/forexauto/src/config"
	"github.com/dykoffi/forexauto/src/logger"
)

func main() {
	configService := config.New()
	// dataService := data.New()
	loopInterval, _ := strconv.ParseInt(configService.GetOrDefault("LOOP_INTERVAL", "9"), 10, 64)

	for {
		// panic("demain est dimanche")
		fmt.Printf("Test each %d minutes okay \n", loopInterval)
		logger.New().Error("Panic")

		// FullForexQuoteData, err := dataService.GetFullForexQuote(data.EURUSD)
		// HistoricalData, err := dataService.GetHistoricalDailyForex(data.EURUSD)

		// if err != nil {
		// 	fmt.Println(err)
		// }

		// fmt.Println(*HistoricalData)

		// for _, fd := range *HistoricalData {
		// 	fmt.Println(fd.ID)
		// }

		time.Sleep(time.Duration(loopInterval) * time.Second)
	}
}
