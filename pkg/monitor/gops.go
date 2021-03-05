package monitor

import (
	"github.com/google/gops/agent"
	"github.com/tal-tech/go-zero/core/logx"
	"time"
)

// TODO: goroutinue会泄露
func StartGoPS() {
	go start()
}

func start() {
	if err := agent.Listen(agent.Options{Addr: ":8082"}); err != nil {
		logx.Errorf("初始化gops失败: %s\n", err.Error())
	}
	time.Sleep(time.Hour)
}
