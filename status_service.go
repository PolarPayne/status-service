package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

const (
	defaultHost       = "0.0.0.0"
	defaultPort       = 80
	defaultStatusCode = 200
)

var (
	flagHost       = flag.String("host", defaultHost, "host to bind to")
	flagPort       = flag.Int("port", defaultPort, "port to bind to")
	flagStatusCode = flag.Int("code", defaultStatusCode, "status code that is returned to each request")
	flagHelp       = flag.Bool("help", false, "print help message")
)

func getStatusCode() int {
	var (
		status int
		err    error
	)

	v, ok := os.LookupEnv("STATUS_CODE")
	if !ok {
		status = *flagStatusCode
		log.Printf("STATUS_CODE is not set, using %v", status)
	} else {
		status, err = strconv.Atoi(v)
		if err != nil {
			log.Printf("STATUS_CODE is not a valid integer (was '%v')", v)
		}
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

func getHost() string {
	host, ok := os.LookupEnv("HOST")
	if !ok {
		log.Printf("HOST is not set, using %v", *flagHost)
		return *flagHost
	}

	return host
}

func getPort() int {
	v, ok := os.LookupEnv("PORT")
	if !ok {
		log.Printf("PORT is not set, using %v", *flagPort)
		return *flagPort
	}

	port, err := strconv.Atoi(v)
	if err != nil {
	}

	return port
}

func main() {
	flag.Parse()

	if *flagHelp {
		flag.PrintDefaults()
		os.Exit(2)
	}

	status := getStatusCode()
	statusMessage := []byte(getStatusMessage(status))

	log.Printf("using status of %v with a message of '%v'", status, string(statusMessage))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v - %v", r.Method, r.URL)
		w.WriteHeader(status)
		w.Write(statusMessage)
	})

	host, port := getHost(), getPort()
	addr := net.JoinHostPort(host, strconv.Itoa(port))

	log.Printf("listening on %v", addr)

	log.Fatal(http.ListenAndServe(addr, nil))
}
