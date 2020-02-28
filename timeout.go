package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {

	// create a new context with deadline
	// based on previous base context
	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Millisecond)

	// make request
	req, _ := http.NewRequest(http.MethodGet, "http://google.com", nil)

	// associate context with request
	req = req.WithContext(ctx)

	// create client and do a request
	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}

	fmt.Println("Response received, status code:", res.StatusCode)
}
