package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Declare port
const PORT = "8080"

// Web server example
func main() {
	// Serve web page with handler
	fileServer := http.FileServer(http.Dir("./public"))
	http.Handle("/", fileServer)
	http.HandleFunc("/fizzbuzz", fizzbuzzHandler)

	fmt.Printf("Starting server at port %s\n", PORT)

	// Error handling example
	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		log.Fatal(err)
	}
}

// Request handler example
func fizzbuzzHandler(w http.ResponseWriter, r *http.Request) {
	// Check request path
	if r.URL.Path != "/fizzbuzz" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	// Check request method
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	query := r.URL.Query()

	// Check request query parameters
	start, err := strconv.Atoi(query.Get("start"))
	if err != nil || start < -10000 || start > 10000 {
		fmt.Fprintf(w, "Start value is not an integer between -10000 and 10000")
		return
	}

	number, err := strconv.Atoi(query.Get("number"))
	if err != nil || number < 1 || number > 10000 {
		fmt.Fprintf(w, "Number value is not an integer between 1 and 10000")
		return
	}

	keys := query["cond-key"]
	strs := query["cond-str"]

	if len(keys) != len(strs) {
		fmt.Fprintf(w, "Number of keys and number of strs do not match")
		return
	}

	if len(keys) > 10 {
		fmt.Fprintf(w, "Number of conditions exceed 10")
		return
	}

	conds := make([]cond, len(keys))
	for i := 0; i < len(keys); i++ {
		key, err := strconv.Atoi(keys[i])
		if err != nil || key < 1 || key > 10000 {
			fmt.Fprintf(w, "Key #%v is not an integer between 1 and 10000", i+1)
			return
		}

		str := strs[i]
		if len(str) < 1 || len(str) > 4 {
			fmt.Fprintf(w, "Str #%v does not have length between 1 and 4", i+1)
			return
		}

		conds[i] = cond{
			key: key,
			str: str,
		}
	}

	// Respond with fizzbuzz output
	fmt.Fprintf(w, strings.Join(fizzbuzz(start, number, conds...), " "))
}
