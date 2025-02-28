package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/bereket-alemayehu/food-recipes/backend/internal/actions"
)

func main() {
    http.HandleFunc("/login", actions.LoginHandler)
    http.HandleFunc("/signup", actions.SignupHandler)
    http.HandleFunc("/create-recipe", actions.CreateRecipeHandler)
    http.HandleFunc("/update-recipe", actions.UpdateRecipeHandler)
    http.HandleFunc("/delete-recipe", actions.DeleteRecipeHandler)
    http.HandleFunc("/like-recipe", actions.LikeRecipeHandler)
    http.HandleFunc("/rate-recipe", actions.RateRecipeHandler)
    http.HandleFunc("/comment-recipe", actions.CommentRecipeHandler)
    http.HandleFunc("/bookmark-recipe", actions.BookmarkRecipeHandler)
    http.HandleFunc("/buy-recipe", actions.BuyRecipeHandler)
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