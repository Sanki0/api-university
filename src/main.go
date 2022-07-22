package main

import "github.com/Sanki0/api-university/src/app"



func main() {
	a := app.App{}
	a.Initialize()
	a.Run(":8080")
}