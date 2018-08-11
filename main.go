package main

import (
	"net/http"
	"fmt"
	"log"
	"strings"
)


func ReturnIP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", strings.Split(r.RemoteAddr, ":")[0])
}

func ReturnPort(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", strings.Split(r.RemoteAddr, ":")[1])
}

func ReturnALL(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", r.RemoteAddr)
}

func main() {
	http.HandleFunc("/", ReturnIP)
	http.HandleFunc("/port", ReturnPort)
	http.HandleFunc("/all", ReturnALL)
	log.Fatal(http.ListenAndServe(":8090", nil))
}