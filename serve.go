package main

import "go-gin/routes"

func main() {
	r := routes.ProvideRoutes()
	err := r.Run(":2345")
	if err != nil {
		panic("error !")
	}
}
