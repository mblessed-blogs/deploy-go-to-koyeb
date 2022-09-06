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
	response := "Home page lol"
	json.NewEncoder(w).Encode(response)
}

// about handler is unexported due to no other package requiring the handler
func about(w http.ResponseWriter, r *http.Request) {
	response := "about lol"
	json.NewEncoder(w).Encode(response)
}

func main() {

	// load our env
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("port number: ", os.Getenv("port"))
		log.Println("Port number: ", os.Getenv("Port"))
		log.Println("PORT number: ", os.Getenv("PORT"))
		log.Fatal("Unable to load env file")
	}

	// create a new router
	router := chi.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/about", about)

	// portNumber
	var portNumber = os.Getenv("port")

	log.Println("Port number: ", portNumber)

	portNumber = ":" + portNumber
	// create our server
	srv := &http.Server{
		Addr:    portNumber,
		Handler: router,
	}

	fmt.Printf("Starting server on port %s\n", portNumber)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
