package main

import "github.com/nayyyyy/golang-assignment/routes"

func main(){
	var PORT = ":8080"

	routes.StartServer().Run(PORT)
}