package main

import (
	"os"
	"net/http"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		rw.Write([]byte(`{"hello":"world"}`))
	})

	if err := http.ListenAndServe(":" + port, nil); err != nil {
		println(err.Error())
	}
}

