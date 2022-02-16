package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type ResponseObject struct {
	Msg  string    `json:"msg"`
	Time time.Time `json:"time"`
}

func main() {
	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		fmt.Println(t)
		resp := ResponseObject{Msg: "demo-api", Time: t}
		_ = json.NewEncoder(w).Encode(resp)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
