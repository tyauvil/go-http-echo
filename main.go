package main

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	type JSONResponse struct {
		Method        string
		URL           string
		Proto         string
		ProtoMajor    int
		ProtoMinor    int
		Header        http.Header
		Body          io.ReadCloser
		ContentLength int64
		Host          string
		RemoteAddr    string
		RequestURI    string
		TLS           *tls.ConnectionState
	}

	jsondata := JSONResponse{
		Method:        r.Method,
		URL:           r.URL.String(),
		Proto:         r.Proto,
		ProtoMajor:    r.ProtoMajor,
		ProtoMinor:    r.ProtoMinor,
		Header:        r.Header,
		Body:          r.Body,
		ContentLength: r.ContentLength,
		Host:          r.Host,
		RemoteAddr:    r.RemoteAddr,
		RequestURI:    r.RequestURI,
		TLS:           r.TLS,
	}

	jsonbytes, err := json.Marshal(jsondata)
	if err != nil {
		// handle error
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonbytes)
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func main() {
	port := getEnv("PORT", "8080")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
