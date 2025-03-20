package main

import "TESTGIN/router"

func main() {
	r := router.Router()
	r.Run(":8080")
}
