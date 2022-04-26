package main

import (
	"fmt"
	http "backend/pkg/http"
)

func main() {
	fmt.Println("The beer server is on tap now: http://localhost:8080")
	r := http.SetupRouter()
	r.Run(":8080")
}