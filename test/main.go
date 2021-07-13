package main

import (
	"fmt"
)

type Event struct {
	Handlers map[int32][]interface{}
	p10_0    func(struct {
		a int
		b string
	})
}

//test Msg
func main() {
	Event := &Event{
		Handlers: make(map[int32][]interface{}),
	}
	Event.AddHandler(10, func(data struct {
		a int
		b string
	}) {
		fmt.Printf("data: %v\n", data)
	})

	Event.printFuncArgs()
	Event.interface2Func()

	Event.p10_0(struct {
		a int
		b string
	}{a: 1234, b: "s"})
}

func (e *Event) AddHandler(packetid int32, f interface{}) {
	e.Handlers[packetid] = append(e.Handlers[packetid], f)
}

func (e *Event) interface2Func() {

	f := e.Handlers[10][0]

	switch v := f.(type) {
	case func(struct {
		a int
		b string
	}):
		e.p10_0 = v
	}

}

func (e *Event) printFuncArgs() {
	for k, v := range e.Handlers {
		for i := 0; i < len(v); i++ {
			fmt.Printf("ID %d: %T\n", k, v[i])
		}
	}
}
