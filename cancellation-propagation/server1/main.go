package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/endpoint", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Server 1: Received request")

		req, err := http.NewRequestWithContext(r.Context(), "GET", "http://localhost:8081/endpoint", nil)
		if err != nil {
			fmt.Println("Server 1: Error creating request:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		client := &http.Client{
			Timeout: 2 * time.Second, // Timeout for the request
		}

		resp, err := client.Do(req)
		if err != nil {
			// Check if the error is due to context cancellation
			if err == context.Canceled {
				fmt.Println("Server 1: Client canceled the request")
			} else if err, ok := err.(net.Error); ok && err.Timeout() {
				fmt.Println("Server 1: Request to Server 2 timed out")
			} else {
				fmt.Println("Server 1: Error making request to Server 2:", err)
			}
			http.Error(w, "Error contacting Server 2", http.StatusGatewayTimeout)
			return
		}

		defer resp.Body.Close()

		fmt.Println("Server 1: Received response from Server 2:", resp.Status)
		fmt.Fprintln(w, "Server 1: Handler completed")
	})

	fmt.Println("Server 1 is listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server 1 error:", err)
	}
}
