package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusInternalServerError)
		return
	}
	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading the file", http.StatusInternalServerError)
		return
	}

	result, err := service.Convert(string(data))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rawTime := time.Now().UTC().String()
	ext := filepath.Ext(header.Filename)
	parts := strings.Split(rawTime, " ")
	timePart := strings.Split(parts[1], ".")[0]
	timePart = strings.ReplaceAll(timePart, ":", "-")
	filename := parts[0] + "_" + timePart + ext

	out, err := os.Create(filename)
	if err != nil {
		http.Error(w, "Error creating the output file", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	_, err = out.WriteString(result)
	if err != nil {
		http.Error(w, "Error writing to the output file", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(result))
}
