package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackthomson/acronym-decipher/api"
)

func main() {
	api := api.NewAPI()

	server := http.NewServeMux()

	server.HandleFunc("/decipher", api.GetWord)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	fmt.Println("Server is running on port: ", port)

	log.Fatal(http.ListenAndServe(":8080", server))
}
