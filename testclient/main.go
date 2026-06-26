package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func main() {
	http.HandleFunc("/cb", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		state := r.URL.Query().Get("state")

		fmt.Fprintf(w, "Received code: %s\nState: %s\n\n", code, state)

		// Prepare token request
		data := url.Values{}
		data.Set("grant_type", "authorization_code")
		data.Set("code", code)
		data.Set("redirect_uri", "http://localhost:3000/cb")
		data.Set("client_id", "test")

		resp, err := http.Post(
			"http://localhost:8080/token",
			"application/x-www-form-urlencoded",
			bytes.NewBufferString(data.Encode()),
		)
		if err != nil {
			fmt.Fprintf(w, "Token request error: %v\n", err)
			return
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		fmt.Fprintf(w, "Token response:\n%s\n", string(body))
	})

	log.Println("Test client running on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
