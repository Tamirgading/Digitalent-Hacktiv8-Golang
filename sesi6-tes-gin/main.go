package main

import "sesi6-tes-gin/routers"

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
