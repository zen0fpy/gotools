package simple_factory

import "fmt"

type API interface {
	SayHello(name string) string
}

type api struct {
}

func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &helloAPI{}
	}
	return nil
}

type hiAPI struct {
}

func (h *hiAPI) SayHello(name string) string {
	msg := fmt.Sprintf("Hi %s", name)
	return msg
}

type helloAPI struct {
}

func (h *helloAPI) SayHello(name string) string {
	msg := fmt.Sprintf("Hello %s", name)
	return msg
}
