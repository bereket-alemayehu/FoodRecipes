package actions

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/machinebox/graphql"
)
type Recipe struct {
    ID              string        `json:"id"`
    Title           string        `json:"title"`
    Description     string        `json:"description"`
    CategoryID      string        `json:"category_id"`
    PreparationTime int           `json:"preparation_time"`
    Ingredients     []interface{} `json:"ingredients"` // Adjust type as needed
    Steps           []interface{} `json:"steps"`       // Adjust type as needed
    Images          []string      `json:"images"`
    FeaturedImage   string        `json:"featured_image"`
    UserID          string        `json:"user_id"`
    // ... other recipe fields
}

func CreateRecipeHandler(w http.ResponseWriter, r *http.Request) {
    var recipe Recipe

    if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // GraphQL mutation to create recipe
    req := graphql.NewRequest(`
        mutation ($title: String!, $description: String!, $categoryId: String!, $preparationTime: Int!, $ingredients: jsonb!, $steps: jsonb!, $images: [String!]!, $featuredImage: String!, $userId: String!) {
            insert_recipes_one(object: {title: $title, description: $description, category_id: $categoryId, preparation_time: $preparationTime, ingredients: $ingredients, steps: $steps, images: $images, featured_image: $featuredImage, user_id: $userId}) {
                id
                title
                description
                category_id
                preparation_time
            }
        }
    `)
    req.Var("title", recipe.Title)
    req.Var("description", recipe.Description)
    req.Var("categoryId", recipe.CategoryID)
    req.Var("preparationTime", recipe.PreparationTime)
    req.Var("ingredients", recipe.Ingredients)
    req.Var("steps", recipe.Steps)
    req.Var("images", recipe.Images)
    req.Var("featuredImage", recipe.FeaturedImage)
    req.Var("userId", recipe.UserID)
    req.Header.Set("x-hasura-admin-secret", os.Getenv("HASURA_ADMIN_SECRET"))

    var respData struct {
        InsertRecipesOne Recipe `json:"insert_recipes_one"`
    }

    if err := graphqlClient.Run(context.Background(), req, &respData); err != nil {
        http.Error(w, "Failed to create recipe", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(respData.InsertRecipesOne)
}

func UpdateRecipeHandler(w http.ResponseWriter, r *http.Request) {
    var recipe Recipe

    if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // GraphQL mutation to update recipe
    req := graphql.NewRequest(`
        mutation ($id: uuid!, $title: String!, $description: String!, $categoryId: String!, $preparationTime: Int!, $ingredients: jsonb!, $steps: jsonb!, $images: [String!]!, $featuredImage: String!) {
            update_recipes_by_pk(pk_columns: {id: $id}, _set: {title: $title, description: $description, category_id: $categoryId, preparation_time: $preparationTime, ingredients: $ingredients, steps: $steps, images: $images, featured_image: $featuredImage}) {
                id
                title
                description
                category_id
                preparation_time
            }
        }
    `)
    req.Var("id", recipe.ID)
    req.Var("title", recipe.Title)
    req.Var("description", recipe.Description)
    req.Var("categoryId", recipe.CategoryID)
    req.Var("preparationTime", recipe.PreparationTime)
    req.Var("ingredients", recipe.Ingredients)
    req.Var("steps", recipe.Steps)
    req.Var("images", recipe.Images)
    req.Var("featuredImage", recipe.FeaturedImage)
    req.Header.Set("x-hasura-admin-secret", os.Getenv("HASURA_ADMIN_SECRET"))

    var respData struct {
        UpdateRecipesByPk Recipe `json:"update_recipes_by_pk"`
    }

    if err := graphqlClient.Run(context.Background(), req, &respData); err != nil {
        http.Error(w, "Failed to update recipe", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(respData.UpdateRecipesByPk)
}
func DeleteRecipeHandler(w http.ResponseWriter, r *http.Request) {
    var recipe struct {
        ID string `json:"id"`
    }

    if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // GraphQL mutation to delete recipe
    req := graphql.NewRequest(`
        mutation ($id: uuid!) {
            delete_recipes_by_pk(id: $id) {
                id
            }
        }
    `)
    req.Var("id", recipe.ID)
    req.Header.Set("x-hasura-admin-secret", os.Getenv("HASURA_ADMIN_SECRET"))

    var respData struct {
        DeleteRecipesByPk Recipe `json:"delete_recipes_by_pk"`
    }

    if err := graphqlClient.Run(context.Background(), req, &respData); err != nil {
        http.Error(w, "Failed to delete recipe", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(respData.DeleteRecipesByPk)
}