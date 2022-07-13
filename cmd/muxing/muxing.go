package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/
func handleGetName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Hello, %v!", vars["PARAM"])
}

func handleGetBad(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func handlePostData(w http.ResponseWriter, r *http.Request) {
	d, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	//w.Write(d)
	fmt.Fprintf(w, "I got message:\n%s", d)
}

func handlePostHeaders(w http.ResponseWriter, r *http.Request) {
	a := r.Header.Get("a")
	b := r.Header.Get("b")
	an, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	bn, err := strconv.Atoi(b)
	if err != nil {
		panic(err)
	}
	cn := an + bn
	c := strconv.Itoa(cn)
	w.Header().Add("a+b", c)
	//w.WriteHeader("a+b", cn)
}

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{PARAM}", handleGetName).Methods(http.MethodGet)
	router.HandleFunc("/bad", handleGetBad).Methods(http.MethodGet)
	router.HandleFunc("/data", handlePostData).Methods(http.MethodPost)
	router.HandleFunc("/headers", handlePostHeaders).Methods(http.MethodPost)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
