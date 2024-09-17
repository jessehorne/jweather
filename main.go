package main

import (
	"fmt"
	"github.com/jessehorne/jweather/internal/api/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/points", handlers.PointsQueryHandler)
	fmt.Println("serving at 127.0.0.1:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}
