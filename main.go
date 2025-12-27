package main

import (
	"{{MODULE_NAME}}/bootstrap"
	_ "{{MODULE_NAME}}/routes"
)

func main() {
	bootstrap.Run()
}
