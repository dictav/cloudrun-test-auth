package main

import (
	"net/http"
	"io"
	"os"
)

func main() {
	backend := os.Getenv("BACKEND")
	if backend == "" {
		println("BACKEND is required")
		os.Exit(1)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			rw.Header().Set("Access-Control-Max-Age", "1728000")
			rw.Header().Set("Content-Type", "text/plain charset=UTF-8")
			rw.WriteHeader(http.StatusNoContent)
			return
		}

		u := *r.URL
		u.Scheme = "https"
		u.Host = backend

		req, err := http.NewRequest(r.Method, u.String(), r.Body)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		for k,values := range r.Header{
			for _, v := range values {
				req.Header.Add(k, v)
			}
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		for k,values := range res.Header{
			for _, v := range values {
				rw.Header().Add(k, v)
			}
		}
		
		rw.WriteHeader(res.StatusCode)
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Header().Set("Access-Control-Allow-Credentials", "true")
		rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		rw.Header().Set("Access-Control-Allow-Headers", "Accept,Authorization,Cache-Control,Content-Type,DNT,If-Modified-Since,Keep-Alive,Origin,User-Agent,X-Requested-With")

		if _, err := io.Copy(rw, res.Body); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	if err := http.ListenAndServe(":" + port, nil); err != nil {
		println(err.Error())
		os.Exit(2)
	}
}

