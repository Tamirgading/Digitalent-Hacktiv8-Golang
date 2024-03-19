package main

import (
	"sesi6-belajar-gin/routers"
)

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
