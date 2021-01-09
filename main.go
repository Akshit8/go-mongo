package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Print("working")
	time.Tick(5)
	http.ListenAndServe(":8000", nil)
}
