package main

import (
	http "backend/pkg/http"
	"fmt"
)

func main() {
	fmt.Println("The beer server is on tap now: http://localhost:8888")
	r := http.SetupRouter()
	r.Run(":8888")
}
