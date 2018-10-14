package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

func getStatusCode() int {
	v, ok := os.LookupEnv("STATUS_CODE")
	if !ok {
		log.Print("STATUS_CODE is not set, using 200")
		return 200
	}

	status, err := strconv.Atoi(v)
	if err != nil {
		log.Panicf("STATUS_CODE is not a valid integer (was '%v')", v)
	}

	if !((status >= 200 && status <= 299) || (status >= 400 && status <= 499) || (status >= 500 && status <= 599)) {
		log.Panicf("STATUS_CODE must be between 200-299, 400-499 or 500-599 (was %v)", status)
	}

	return status
}

func getStatusMessage(statusCode int) string {
	statusMessage, ok := os.LookupEnv("STATUS_MESSAGE")
	if !ok {
		out := http.StatusText(statusCode)
		log.Printf("STATUS_MESSAGE is not set, using '%v'", out)
		return out
	}

	return statusMessage
}

func getAddr() string {
	host, ok := os.LookupEnv("HOST")
	if !ok {
		log.Printf("HOST is not set, using 0.0.0.0")
		host = "0.0.0.0"
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		log.Printf("PORT is not set, using 80")
		port = "80"
	}

	return net.JoinHostPort(host, port)
}

func main() {
	status := getStatusCode()
	statusMessage := []byte(getStatusMessage(status))

	log.Printf("using status of %v with a message of '%v'", status, string(statusMessage))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v - %v", r.Method, r.URL)
		w.WriteHeader(status)
		w.Write(statusMessage)
	})

	addr := getAddr()
	log.Printf("listening on %v", addr)

	log.Fatal(http.ListenAndServe(addr, nil))
}
