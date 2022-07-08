package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/houseme/mobiledetect/ua"
)

func main() {
	handler := http.HandlerFunc(handleRequest)
	http.Handle("/is-mobile", handler)
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {

	userAgent := r.UserAgent()
	ua := ua.New(userAgent)
	isMobile := ua.Mobile()

	fmt.Printf("UserAgent:: %s", userAgent)
	fmt.Println(ua.Mobile())

	resp := make(map[string]string)

	resp["userAgent"] = userAgent
	resp["isMobile"] = strconv.FormatBool(isMobile)
	jsonResp, err := json.Marshal(resp)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
