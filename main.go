package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func video(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(405)
		w.Write([]byte("Method not allowed"))
		return
	}
	w.Header().Set("Content-Type", "video/mp4")
	http.ServeFile(w, r, "video/SampleVideo_1280x720_1mb.mp4")
}

func main() {
	port := os.Getenv("PORT")
	fmt.Println(port)
	mux := http.NewServeMux()
	mux.HandleFunc("/video", video)

	log.Println("Video streaming on :" + port)
	err := http.ListenAndServe(":"+port, mux)
	log.Fatal(err)
}
