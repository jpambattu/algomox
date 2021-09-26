package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	// "path/filepath"
	// "time"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	http.ServeFile(w, r, "index.html")
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// truncated for brevity

	// The argument to FormFile must match the name attribute
	// of the file input on the frontend
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	//

	// buff := make([]byte, 512)
	// _, err = file.Read(buff)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// filetype := http.DetectContentType(buff)
	// if filetype != "application/json" {
	// 	http.Error(w, filetype+"The provided file format is not allowed. Please upload a JPEG or PNG image", http.StatusBadRequest)
	// 	return
	// }

	// _, errr := file.Seek(0, io.SeekStart)
	// if errr != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// Create the uploads folder if it doesn't
	// already exist
	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a new file in the uploads directory
	dst, err := os.Create(fmt.Sprintf("./uploads/%s", fileHeader.Filename))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer dst.Close()

	// Copy the uploaded file to the filesystem
	// at the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Upload successful")

}

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/upload", uploadHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
