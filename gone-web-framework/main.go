package main

import (
	"fmt"
	"gone"
	"log"
	"net/http"
)

func main() {

	engine := gone.New()

	engine.GET("/hello", helloHandler)

	log.Fatal(engine.Run(":9999"))

	// http.HandleFunc("/hello", helloHandler)
	// log.Fatal(http.ListenAndServe(":9999", nil))
}

// handler echoes r.URL.Header
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
