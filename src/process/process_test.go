package process

import (
	"fmt"
	"log"
	"math"
	"os"
	"testing"
	"time"

	"github.com/dykoffi/forexauto/src/data"
)

func BenchmarkCollectFullForexQuote(b *testing.B) {

	err := os.Chdir("../../") // Remonte d'un niveau pour arriver à la racine
	if err != nil {
		log.Fatalf("Erreur lors du changement de répertoire : %v", err)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		New().CollectIntraDayForex()
	}
}

func TestProcess(t *testing.T) {
	err := os.Chdir("../../")
	if err != nil {
		log.Fatalf("Erreur lors du changement de répertoire : %v", err)
	}

	from, err := time.Parse("2006-01-02", "2024-01-01")
	if err != nil {

	}
	to := time.Now()

	days := math.Ceil(to.Sub(from).Hours() / 24)

	newDate := from

	var dates []string

	for i := 1; i < int(days); i++ {
		newDate = newDate.Add(24 * time.Hour)
		if newDate.Weekday() == time.Saturday || newDate.Weekday() == time.Sunday {
			continue
		}

		newDateFormat := newDate.Format("2006-01-02")
		dates = append(dates, newDateFormat)

		go data.New().GetIntraDayForex(newDateFormat, newDateFormat)

	}

	time.Sleep(10 * time.Second)
	fmt.Println(days)
	fmt.Println(len(dates))
	fmt.Println(dates)

}
