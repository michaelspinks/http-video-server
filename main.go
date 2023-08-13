package main

import (
	"log"
	"net/http"
)

func video(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "video/mp4")
	http.ServeFile(w, r, "video/SampleVideo_1280x720_1mb.mp4")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/video", video)

	log.Println("Video streaming on :3000")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
