package monitor

import "github.com/tal-tech/go-zero/core/prometheus"

func StartPrometheus(c prometheus.Config) {
	go prometheus.StartAgent(c)
}
