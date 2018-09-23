package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	type JSONResponse struct {
		Method           string
		URL              string
		Proto            string
		ProtoMajor       int
		ProtoMinor       int
		Header           http.Header
		Body             io.ReadCloser
		ContentLength    int64
		TransferEncoding []string
		Host             string
		Trailer          http.Header
		RemoteAddr       string
		RequestURI       string
		Response         *http.Response
	}

	jsondata := JSONResponse{
		Method:           r.Method,
		URL:              r.URL.String(),
		Proto:            r.Proto,
		ProtoMajor:       r.ProtoMajor,
		ProtoMinor:       r.ProtoMinor,
		Header:           r.Header,
		Body:             r.Body,
		ContentLength:    r.ContentLength,
		TransferEncoding: r.TransferEncoding,
		Host:             r.Host,
		Trailer:          r.Trailer,
		RemoteAddr:       r.RemoteAddr,
		RequestURI:       r.RequestURI,
		Response:         r.Response,
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
