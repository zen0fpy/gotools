package main

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var w *httptest.ResponseRecorder
var mux *http.ServeMux

func TestMain(m *testing.M) {
	mux = http.NewServeMux()
	mux.HandleFunc("/topic", handleRequest)
	w = httptest.NewRecorder()
	os.Exit(m.Run())
}

func TestHttpPost(t *testing.T) {
	reader := strings.NewReader(`{"title":"The Go Standard Library","Content":"It contains many packages."}`)
	r, err := http.NewRequest(http.MethodPost, "/topic", reader)
	if err != nil {
		log.Fatalln(err)
	}

	mux.ServeHTTP(w, r)
	res := w.Result()
	fmt.Printf("%+v\n", w.Body.Bytes())
	require.Equal(t, 200, res.StatusCode)

	body, err := ioutil.ReadAll(res.Body)
	require.Equal(t, nil, err)
	fmt.Printf("body: %s\n", body)

}

func TestHttpGet(t *testing.T) {
	r, _ := http.NewRequest(http.MethodGet, "/topic/0", nil)
	mux.ServeHTTP(w, r)
	resp := w.Result()

	require.Equal(t, 200, resp.StatusCode)
	topic := new(Topic)
	json.Unmarshal(w.Body.Bytes(), topic)
	require.Equal(t, 0, topic.Id)
}
