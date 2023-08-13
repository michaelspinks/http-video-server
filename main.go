package main

import (
	"log"
	"net/http"
)

func video(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is your video"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/video", video)

	log.Println("Video streaming on :3000")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
