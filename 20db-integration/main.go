package main

import (
	"fmt"
	"net/http"

	"github.com/SohamChatterG/mod/router"
)

func main() {
	fmt.Println("first go crud app")
	r := router.Router()
	http.ListenAndServe(":4000", r)
	fmt.Println("server is listening on port 4000")

}
