package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Printf("error load enf file")
		os.Exit(1)
	}

	port := os.Getenv("PORT")
	appPort := "0.0.0.0:" + port

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(appPort, nil))
}

// func main() {
// 	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 	// 	fmt.Fprintf(w, "Hello World! faiz baraja")
// 	// })

// 	// envErr := godotenv.Load(".env")
// 	// if envErr != nil {
// 	// 	fmt.Printf("error load enf file")
// 	// 	os.Exit(1)
// 	// }

// 	// port := os.Getenv("PORT")
// 	// appPort := "0.0.0.0:" + port
// 	// fmt.Println(appPort)
// 	// http.ListenAndServe(appPort, nil)
// }
