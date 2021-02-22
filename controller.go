package main

type Controller interface{
	init()
	run(port string)
}
