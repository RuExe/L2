package main

import "fmt"

type Context struct {
	state State
}

func (c *Context) SetState(state State) {
	c.state = state
}

func (c *Context) req1() {
	c.state.handle1()
}

func (c *Context) req2() {
	c.state.handle2()
}

type State interface {
	handle1()
	handle2()
}

type StateA struct {
}

func (s StateA) handle1() {
	fmt.Println("StateA handles request1.")
}

func (s StateA) handle2() {
	fmt.Println("StateA handles request2.")
}

type StateB struct {
}

func (s StateB) handle1() {
	fmt.Println("StateB handles request1.")
}

func (s StateB) handle2() {
	fmt.Println("StateB handles request2.")
}

func main() {
	context := Context{state: StateA{}}
	context.req1()
	context.req2()

	context.SetState(StateB{})
	context.req1()
	context.req2()
}
