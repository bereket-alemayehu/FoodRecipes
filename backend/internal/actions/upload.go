package actions

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
    // Parse the multipart form
    err := r.ParseMultipartForm(10 << 20) // 10 MB
    if err != nil {
        http.Error(w, "Unable to parse form", http.StatusBadRequest)
        return
    }

    // Retrieve the file from form data
    file, handler, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "Unable to retrieve file", http.StatusBadRequest)
        return
    }
    defer file.Close()

    // Create a directory to store the uploaded files
    uploadDir := "./uploads"
    if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
        os.Mkdir(uploadDir, os.ModePerm)
    }

    // Create the file on the server
    filePath := filepath.Join(uploadDir, handler.Filename)
    dst, err := os.Create(filePath)
    if err != nil {
        http.Error(w, "Unable to create file", http.StatusInternalServerError)
        return
    }
    defer dst.Close()

    // Copy the uploaded file to the server
    if _, err := io.Copy(dst, file); err != nil {
        http.Error(w, "Unable to save file", http.StatusInternalServerError)
        return
    }

    // Respond with the file path
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, `{"file_path": "%s"}`, filePath)
}