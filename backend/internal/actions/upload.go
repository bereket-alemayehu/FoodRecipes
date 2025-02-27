package actions

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
        // Parse multipart form
        err := r.ParseMultipartForm(10 << 20) // 10 MB limit
        if err != nil {
                http.Error(w, "File too large or invalid form", http.StatusBadRequest)
                return
        }

        file, handler, err := r.FormFile("file") // "file" is the form field name
        if err != nil {
                http.Error(w, "Error retrieving file", http.StatusBadRequest)
                return
        }
        defer file.Close()

        // Generate unique filename
        ext := filepath.Ext(handler.Filename)
        filename := uuid.New().String() + ext
        uploadDir := "uploads" // Create this directory if it doesn't exist

        // Create the uploads directory if it doesn't exist
        if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
                os.Mkdir(uploadDir, 0755)
        }

        filePath := filepath.Join(uploadDir, filename)

        // Create the file
        dst, err := os.Create(filePath)
        if err != nil {
                http.Error(w, "Failed to create file", http.StatusInternalServerError)
                return
        }
        defer dst.Close()

        // Copy the uploaded file to the destination
        _, err = io.Copy(dst, file)
        if err != nil {
                http.Error(w, "Failed to save file", http.StatusInternalServerError)
                return
        }

        // Generate URL for the uploaded file
        fileURL := fmt.Sprintf("/%s/%s", uploadDir, filename) // Adjust based on your server setup

        // Respond with the file URL
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{"url": fileURL})
}