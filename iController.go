package main

type iController interface{
	init()
	run(port string)
}

type Controller struct {
	controller iController
}

func (c Controller) start(port string) {
	c.controller.init()
	c.controller.run(port)
}
