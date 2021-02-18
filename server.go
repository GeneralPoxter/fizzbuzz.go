package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileServer := http.FileServer(http.Dir("./public"))
	http.Handle("/", fileServer)
	http.HandleFunc("/fizzbuzz", fizzbuzzHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(GetPort(), nil); err != nil {
		log.Fatal(err)
	}
}

func fizzbuzzHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/fizzbuzz" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	start, err := strconv.Atoi(r.URL.Query().Get("start"))
	if err != nil || start < -10000 || start > 10000 {
		fmt.Fprintf(w, "Start value is not an integer between -10000 and 10000")
		return
	}

	number, err := strconv.Atoi(r.URL.Query().Get("number"))
	if err != nil || number < 1 || number > 10000 {
		fmt.Fprintf(w, "Number value is not an integer between 1 and 10000")
		return
	}

	fmt.Fprintf(w, strings.Join(fizzbuzz(start, number,
		cond{3, "fizz"},
		cond{5, "buzz"}), " "))
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
