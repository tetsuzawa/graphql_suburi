package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type Response struct {
	Status string `json:"status"`
	Message string `json:"data"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set( "Access-Control-Allow-Methods","GET, POST, PUT, DELETE, OPTIONS")
		res := Response {
			Status: "ok",
			Message: "takurinton",
		}
		json, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	})
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		res := Response {
			Status: "ok",
			Message: "pong",
		}
		json, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	})
	fmt.Println("listening on :8888 ...")
	if err := http.ListenAndServe(":8888", nil); err !=nil{
		log.Fatalln(err)
	}
}