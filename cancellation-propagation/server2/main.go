package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/endpoint", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Server 2: Received request")

		// Simulate long processing time
		select {
		case <-time.After(5 * time.Second):
			fmt.Fprintln(w, "Server 2: Completed processing")
			fmt.Println("Server 2: Completed processing")
		case <-r.Context().Done():
			// Detect if the request context is canceled
			fmt.Println("Server 2: Request canceled")
			http.Error(w, "Request canceled", http.StatusRequestTimeout)
		}
	})

	fmt.Println("Server 2 is listening on port 8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Println("Server 2 error:", err)
	}
}
