package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type Response struct {
	Status string `json:"status"`
	Data string `json:"data"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		res := Response {
			Status: "ok",
			Data: "takurinton",
		}
		json, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	})
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		res := Response {
			Status: "ok",
			Data: "pong",
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