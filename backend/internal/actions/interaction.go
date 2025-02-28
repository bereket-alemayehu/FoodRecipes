package actions

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/machinebox/graphql"
)

func LikeRecipeHandler(w http.ResponseWriter, r *http.Request) {
    var input struct {
        RecipeID string `json:"recipe_id"`
        UserID   string `json:"user_id"`
    }

    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // GraphQL mutation to like a recipe
    req := graphql.NewRequest(`
        mutation ($recipeId: uuid!, $userId: uuid!) {
            insert_likes_one(object: {recipe_id: $recipeId, user_id: $userId}) {
                recipe_id
                user_id
            }
        }
    `)
    req.Var("recipeId", input.RecipeID)
    req.Var("userId", input.UserID)
    req.Header.Set("x-hasura-admin-secret", os.Getenv("HASURA_ADMIN_SECRET"))

    var respData struct {
        InsertLikesOne struct {
            RecipeID string `json:"recipe_id"`
            UserID   string `json:"user_id"`
        } `json:"insert_likes_one"`
    }

    if err := graphqlClient.Run(context.Background(), req, &respData); err != nil {
        http.Error(w, "Failed to like recipe", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(respData.InsertLikesOne)
}

func RateRecipeHandler(w http.ResponseWriter, r *http.Request) {
    var input struct {
        RecipeID string  `json:"recipe_id"`
        UserID   string  `json:"user_id"`
        Rating   float64 `json:"rating"`
    }

    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // GraphQL mutation to rate a recipe
    req := graphql.NewRequest(`
        mutation ($recipeId: uuid!, $userId: uuid!, $rating: float8!) {
            insert_ratings_one(object: {recipe_id: $recipeId, user_id: $userId, rating: $rating}) {
                recipe_id
                user_id
                rating
            }
        }
    `)
    req.Var("recipeId", input.RecipeID)
    req.Var("userId", input.UserID)
    req.Var("rating", input.Rating)
    req.Header.Set("x-hasura-admin-secret", os.Getenv("HASURA_ADMIN_SECRET"))

    var respData struct {
        InsertRatingsOne struct {
            RecipeID string  `json:"recipe_id"`
            UserID   string  `json:"user_id"`
            Rating   float64 `json:"rating"`
        } `json:"insert_ratings_one"`
    }

    if err := graphqlClient.Run(context.Background(), req, &respData); err != nil {
        http.Error(w, "Failed to rate recipe", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(respData.InsertRatingsOne)
}

func CommentRecipeHandler(w http.ResponseWriter, r *http.Request) {
    var input struct {
        RecipeID string `json:"recipe_id"`
        UserID   string `json:"user_id"`
        Comment  string `json:"comment"`
    }

    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // GraphQL mutation to comment on a recipe
    req := graphql.NewRequest(`
        mutation ($recipeId: uuid!, $userId: uuid!, $comment: String!) {
            insert_comments_one(object: {recipe_id: $recipeId, user_id: $userId, comment: $comment}) {
                recipe_id
                user_id
                comment
            }
        }
    `)
    req.Var("recipeId", input.RecipeID)
    req.Var("userId", input.UserID)
    req.Var("comment", input.Comment)
    req.Header.Set("x-hasura-admin-secret", os.Getenv("HASURA_ADMIN_SECRET"))

    var respData struct {
        InsertCommentsOne struct {
            RecipeID string `json:"recipe_id"`
            UserID   string `json:"user_id"`
            Comment  string `json:"comment"`
        } `json:"insert_comments_one"`
    }

    if err := graphqlClient.Run(context.Background(), req, &respData); err != nil {
        http.Error(w, "Failed to comment on recipe", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(respData.InsertCommentsOne)
}

func BookmarkRecipeHandler(w http.ResponseWriter, r *http.Request) {
    var input struct {
        RecipeID string `json:"recipe_id"`
        UserID   string `json:"user_id"`
    }

    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // GraphQL mutation to bookmark a recipe
    req := graphql.NewRequest(`
        mutation ($recipeId: uuid!, $userId: uuid!) {
            insert_bookmarks_one(object: {recipe_id: $recipeId, user_id: $userId}) {
                recipe_id
                user_id
            }
        }
    `)
    req.Var("recipeId", input.RecipeID)
    req.Var("userId", input.UserID)
    req.Header.Set("x-hasura-admin-secret", os.Getenv("HASURA_ADMIN_SECRET"))

    var respData struct {
        InsertBookmarksOne struct {
            RecipeID string `json:"recipe_id"`
            UserID   string `json:"user_id"`
        } `json:"insert_bookmarks_one"`
    }

    if err := graphqlClient.Run(context.Background(), req, &respData); err != nil {
        http.Error(w, "Failed to bookmark recipe", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(respData.InsertBookmarksOne)
}

func BuyRecipeHandler(w http.ResponseWriter, r *http.Request) {
    var input struct {
        RecipeID string `json:"recipe_id"`
        UserID   string `json:"user_id"`
    }

    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // GraphQL mutation to buy a recipe
    req := graphql.NewRequest(`
        mutation ($recipeId: uuid!, $userId: uuid!) {
            insert_purchases_one(object: {recipe_id: $recipeId, user_id: $userId}) {
                recipe_id
                user_id
            }
        }
    `)
    req.Var("recipeId", input.RecipeID)
    req.Var("userId", input.UserID)
    req.Header.Set("x-hasura-admin-secret", os.Getenv("HASURA_ADMIN_SECRET"))

    var respData struct {
        InsertPurchasesOne struct {
            RecipeID string `json:"recipe_id"`
            UserID   string `json:"user_id"`
        } `json:"insert_purchases_one"`
    }

    if err := graphqlClient.Run(context.Background(), req, &respData); err != nil {
        http.Error(w, "Failed to buy recipe", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(respData.InsertPurchasesOne)
}