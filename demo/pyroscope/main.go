package main

import (
	"fmt"
	"github.com/pyroscope-io/pyroscope/pkg/agent/profiler"
)

func main() {
	profiler.Start(profiler.Config{
		ApplicationName: "backend.purchases",
		ServerAddress:   "http://192.168.2.105:4040",
	})
	for {

		fmt.Printf("tests \n")
	}
}
