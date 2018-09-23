package main

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func getEnv(envvar, def string) string {
	val, env := os.LookupEnv(envvar)
	if !env {
		val = def
	}
	return val
}

func handler(w http.ResponseWriter, r *http.Request) {
	type JSONResponse struct {
		Method        string
		URL           *url.URL
		Protocol      string
		Header        http.Header
		Body          io.ReadCloser
		ContentLength int64
		Host          string
		RemoteAddr    string
		TLS           *tls.ConnectionState
	}

	jsondata := JSONResponse{
		Method:        r.Method,
		URL:           r.URL,
		Protocol:      r.Proto,
		Header:        r.Header,
		Body:          r.Body,
		ContentLength: r.ContentLength,
		Host:          r.Host,
		RemoteAddr:    r.RemoteAddr,
		TLS:           r.TLS,
	}

	jsonbytes, err := json.Marshal(jsondata)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonbytes)
}

func main() {
	port := getEnv("PORT", "8080")
	http.HandleFunc("/", handler)
	log.Println("Starting go-http-echo")
	log.Println("Listening on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
