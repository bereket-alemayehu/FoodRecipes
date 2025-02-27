package main

import (
	"fmt"
	"food-recipes/backend/internal/actions"
	"net/http"
	"os"
	"strconv"
)

func main() {
    http.HandleFunc("/signup", actions.SignupHandler)
    http.HandleFunc("/login", actions.LoginHandler)
    http.HandleFunc("/recipes", actions.CreateRecipeHandler)
    http.HandleFunc("/recipes/update", actions.UpdateRecipeHandler)
    http.HandleFunc("/recipes/delete", actions.DeleteRecipeHandler)
    http.HandleFunc("/upload", actions.UploadFileHandler)

    portStr := os.Getenv("PORT")
    if portStr == "" {
        fmt.Println("PORT environment variable not set, using default port 8081")
        portStr = "8081" // Default port if not set
    }

    port, err := strconv.Atoi(portStr)
    if err != nil {
        fmt.Printf("Error converting PORT to integer: %v\n", err)
        return // Exit if port conversion fails
    }

    fmt.Printf("Server listening on port %d\n", port)
    err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
    if err != nil {
        fmt.Printf("Error starting server: %v\n", err)
    }
}