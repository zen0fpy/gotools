package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	jlog "github.com/opentracing/opentracing-go/log"
	"gotools/pkg/monitor"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	user := map[string]string{
		"username": "admin",
		"password": "adminxxxx",
	}

	traceConf := monitor.TraceConf{
		ServiceName: "user-api",
		HostPort:    "127.0.0.1:6831",
	}

	closer := monitor.InitTracer(traceConf)
	defer closer.Close()

	rootSpan := opentracing.GlobalTracer().StartSpan("login")
	rootSpan.SetTag("login", "login")
	defer rootSpan.Finish()

	encodeData, err := json.Marshal(user)
	fmt.Printf("%s \n", encodeData)
	if err != nil {
		log.Fatalln(err)
	}
	reader := bytes.NewReader(encodeData)
	url := "http://127.0.0.1:8889/v1/user/login"
	req, err := http.NewRequest("POST", url, reader)
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	if err != nil {
		log.Fatalln(err)
	}

	ctx := opentracing.ContextWithSpan(context.Background(), rootSpan)
	fmt.Printf("client ctx: %s\n", ctx)
	clientSpan, _ := opentracing.StartSpanFromContext(ctx, "send-request")
	defer clientSpan.Finish()

	clientSpan.LogFields(
		jlog.String("event", "string-format"),
		jlog.String("value", "login-log"),
	)

	ext.SpanKindRPCClient.Set(clientSpan)
	ext.HTTPUrl.Set(clientSpan, url)
	ext.HTTPMethod.Set(clientSpan, "POST")
	clientSpan.SetTag("HELLO", "world")
	opentracing.GlobalTracer().Inject(clientSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	clientSpan.LogFields(
		jlog.String("resp", string(data)),
	)
	fmt.Printf("data: %s\n", data)

}
