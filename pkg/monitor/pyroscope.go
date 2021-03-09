package monitor

import "github.com/pyroscope-io/pyroscope/pkg/agent/profiler"

func StartPyroscope(appName, address string) {

	profiler.Start(profiler.Config{
		ApplicationName: appName,
		ServerAddress:   address,
	})
}
