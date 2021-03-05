package monitor

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
	"io"
)

type TraceConf struct {
	ServiceName string
	HostPort    string
}

func InitTracer(c TraceConf) io.Closer {

	cfg := &jaegerConfig.Configuration{
		Sampler: &jaegerConfig.SamplerConfig{
			Type:  "const", //固定采样
			Param: 1,       //1=全采样、0=不采样
		},

		Reporter: &jaegerConfig.ReporterConfig{
			LogSpans:           false,
			LocalAgentHostPort: c.HostPort,
		},

		ServiceName: c.ServiceName,
	}

	var closer io.Closer
	var err error
	Tracer, closer, err := cfg.NewTracer(jaegerConfig.Logger(jaeger.StdLogger))

	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.SetGlobalTracer(Tracer)
	return closer
}
