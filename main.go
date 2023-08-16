package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func video(w http.ResponseWriter, r *http.Request) {

	log.Println("Checking environment variables...")

	// Retrieve environment variables
	VIDEO_STORAGE_HOST, ok1 := os.LookupEnv("VIDEO_STORAGE_HOST")
	VIDEO_STORAGE_PORT, ok2 := os.LookupEnv("VIDEO_STORAGE_PORT")

	// Ensure environment variables are set
	if !ok1 || !ok2 {
		http.Error(w, "Environment variables not set", http.StatusInternalServerError)
		log.Println("Environment variables not set!")
		return
	}

	serviceURL := fmt.Sprintf("http://%s:%s/video?path=SampleVideo_1280x720_1mb.mp4", VIDEO_STORAGE_HOST, VIDEO_STORAGE_PORT)

	req, err := http.NewRequest(r.Method, serviceURL, r.Body)
	if err != nil {
		http.Error(w, "Error creating request", http.StatusInternalServerError)
		log.Println("Error creating request!")
		return
	}
	req.Header = r.Header

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "Error forwarding request", http.StatusInternalServerError)
		log.Println("Error forwarding request!")
		return
	}
	log.Printf("Sending Request to %s", serviceURL)
	defer resp.Body.Close()

	// Copy the response from the other service to the original response writer
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/video", video)

	log.Println("Video streaming on :4001")
	err := http.ListenAndServe(":80", mux)
	log.Fatal(err)
}
