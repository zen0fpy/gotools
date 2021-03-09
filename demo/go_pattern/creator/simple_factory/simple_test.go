package simple_factory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHelloAPI(t *testing.T) {
	api := helloAPI{}
	name := "world"
	expect := "Hello world"
	result := api.SayHello(name)
	assert.Equalf(t, expect, result, "wo huo, %s\n", "not ok!")
}

func TestHiAPI(t *testing.T) {
	api := hiAPI{}
	name := "world"
	expect := "Hi world"
	result := api.SayHello(name)
	assert.Equalf(t, expect, result, "wo huo, %s\n", "not ok!")
}

func TestNewAPI(t *testing.T) {
	assert.IsType(t, &hiAPI{}, NewAPI(1))
	assert.IsType(t, &helloAPI{}, NewAPI(2))
	assert.IsType(t, nil, NewAPI(3))
}
