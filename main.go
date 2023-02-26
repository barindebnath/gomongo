package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/barindebnath/gomongo/router"
)

func main() {
	fmt.Println("MONGODB API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listning at port 4000...")
}
