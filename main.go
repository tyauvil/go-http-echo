package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
		Path          string
		RawQuery      string
		Protocol      string
		Header        http.Header
		Body          string
		ContentLength int64
		Host          string
		RemoteAddr    string
		Referer       string
		TLS           *tls.ConnectionState
	}

	defer r.Body.Close()

	bytesBody, err := ioutil.ReadAll(r.Body)
	strBody := fmt.Sprintf("%s", bytesBody)

	if err != nil {
		log.Println(err)
	}

	jsondata := JSONResponse{
		Method:        r.Method,
		Path:          r.URL.Path,
		RawQuery:      r.URL.RawQuery,
		Protocol:      r.Proto,
		Header:        r.Header,
		Body:          strBody,
		ContentLength: r.ContentLength,
		Host:          r.Host,
		Referer:       r.Referer(),
		RemoteAddr:    r.RemoteAddr,
		TLS:           r.TLS,
	}

	jsonbytes, err := json.Marshal(jsondata)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonbytes)
	log.Println(fmt.Sprintf("%s", jsonbytes))
}

func main() {
	port := getEnv("HTTP_PORT", "8080")
	tlsPort := getEnv("TLS_PORT", "8443")
	tls := getEnv("TLS", "")

	http.HandleFunc("/", handler)

	log.Println("Starting go-http-echo")

	if tls == "enabled" {
		log.Println("Listening for http/s on port:", tlsPort)
		go func() {
			log.Fatal(http.ListenAndServeTLS(":"+tlsPort, "server.crt", "server.key", nil))
		}()
	}

	log.Println("Listening for http on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
