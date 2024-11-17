package process

import (
	"log"
	"os"
	"testing"
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
