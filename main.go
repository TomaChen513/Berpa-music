package main

// @title berpar-Music

import (
	"berpar/model"
	"berpar/routes"
)


func main() {
	model.InitDb()
	routes.InitRouter()
}