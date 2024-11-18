package main

import (
	"github.com/dykoffi/forexauto/src/process"
	"github.com/dykoffi/forexauto/src/scheduler"
)

func main() {
	scheduler.New().RunCrons()
	process.New().CollectIntraDayForex()
	select {}
}
