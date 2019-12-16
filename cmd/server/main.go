package main

import (
	"fmt"
	"net/http"
	"os"

	demo "github.com/tekkamanendless/google-function-logging-demo"
)

func main() {
	port := "8080"
	if value := os.Getenv("PORT"); value != "" {
		port = value
	}

	fmt.Printf("Port: %s\n", port)

	address := ":" + port

	handler := http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("REQUEST [IN]\n")
			demo.CloudFunction(w, r)
			fmt.Printf("REQUEST [OUT]\n")
		})

	serveMux := http.NewServeMux()
	serveMux.Handle("/", handler)

	server := &http.Server{
		Addr:    address,
		Handler: serveMux,
	}

	fmt.Printf("Listening on: %s\n", address)
	fmt.Printf("URL: http://localhost%s/\n", address)

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
