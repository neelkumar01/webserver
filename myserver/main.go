package main

import (
	"fmt"
	"io"
	"net/http"
)

func runPage(pattern string, msg string) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, msg)
	})
}
func getResponse(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://catfact.ninja/fact")
	if err != nil {
		fmt.Fprint(w, err)
	}
	body,_ := io.ReadAll(response.Body)
	defer response.Body.Close()
	if response.StatusCode > 299 {
		fmt.Fprintf(w, "response failed!")
	}
	fmt.Fprintf(w, string(body))
}

func main() {
	
	// running some web pages
	runPage("/page1", "Hello from web server \n\nThis is page 1")
	runPage("/page2", "Hello from server \n\nThis is page2")

	// running public files
	http.Handle("/", http.FileServer(http.Dir("public/root")))
	http.Handle("/about/", http.StripPrefix("/about/", http.FileServer(http.Dir("public/about"))))
	
	// getting responses
	http.HandleFunc("/cat/", getResponse)

	// starting server
	fmt.Println("starting server at port 7777")
	http.ListenAndServe(":7777", nil)
}