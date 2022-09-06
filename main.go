package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

// home handler is unexported due to no other package requiring the handler
func home(w http.ResponseWriter, r *http.Request) {
	response := "This is the Home Page"
	json.NewEncoder(w).Encode(response)
}

// loadEnv loads our .env file: we will use this to test locally
func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load env file")
	}
}

func main() {

	// load our env: to be commented out when we push live
	// loadEnv()

	// get PORT number from our environmental variable
	var portNumber = os.Getenv("PORT")
	portNumber = ":" + portNumber

	// create a new router
	router := chi.NewRouter()
	router.HandleFunc("/", home)

	// create our server
	srv := &http.Server{
		Addr:    portNumber,
		Handler: router,
	}

	fmt.Printf("Starting server on port %s\n", portNumber)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
