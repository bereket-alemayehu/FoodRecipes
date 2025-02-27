package actions

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/machinebox/graphql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
        ID       string `json:"id"`
        Username string `json:"username"`
        Email    string `json:"email"`
        Password string `json:"password"`
        // ... other user fields
}

type LoginOutput struct {
        Token  string `json:"token"`
        UserID string `json:"user_id"`
}

var graphqlClient *graphql.Client

func init() {
        graphqlClient = graphql.NewClient(os.Getenv("HASURA_GRAPHQL_ENDPOINT")) // Example: http://localhost:8080/v1/graphql
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
        var signupInput struct {
                Username string `json:"username"`
                Email    string `json:"email"`
                Password string `json:"password"`
        }

        if err := json.NewDecoder(r.Body).Decode(&signupInput); err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
                return
        }

        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signupInput.Password), bcrypt.DefaultCost)
        if err != nil {
                http.Error(w, "Failed to hash password", http.StatusInternalServerError)
                return
        }

        // GraphQL mutation to create user
        req := graphql.NewRequest(`
                mutation ($username: String!, $email: String!, $password: String!) {
                        insert_users_one(object: {username: $username, email: $email, password: $password}) {
                                id
                                username
                                email
                        }
                }
        `)
        req.Var("username", signupInput.Username)
        req.Var("email", signupInput.Email)
        req.Var("password", string(hashedPassword))
        req.Header.Set("x-hasura-admin-secret", os.Getenv("HASURA_ADMIN_SECRET")) // Example: myadminsecretkey

        var respData struct {
                InsertUsersOne User `json:"insert_users_one"`
        }

        if err := graphqlClient.Run(context.Background(), req, &respData); err != nil {
                http.Error(w, "Failed to create user", http.StatusInternalServerError)
                return
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(respData.InsertUsersOne)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
        var loginInput struct {
                Email    string `json:"email"`
                Password string `json:"password"`
        }

        if err := json.NewDecoder(r.Body).Decode(&loginInput); err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
                return
        }

        // GraphQL query to get user by email
        req := graphql.NewRequest(`
                query ($email: String!) {
                        users(where: {email: {_eq: $email}}) {
                                id
                                username
                                password
                        }
                }
        `)
        req.Var("email", loginInput.Email)
        req.Header.Set("x-hasura-admin-secret", os.Getenv("HASURA_ADMIN_SECRET"))

        var respData struct {
                Users []User `json:"users"`
        }

        if err := graphqlClient.Run(context.Background(), req, &respData); err != nil {
                http.Error(w, "Failed to fetch user", http.StatusInternalServerError)
                return
        }

        if len(respData.Users) == 0 {
                http.Error(w, "Invalid email or password", http.StatusUnauthorized)
                return
        }

        user := respData.Users[0]

        if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInput.Password)); err != nil {
                http.Error(w, "Invalid email or password", http.StatusUnauthorized)
                return
        }

        // Generate JWT token
        token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
                "user_id": user.ID,
                "exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
        })

        tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET"))) // Example: myjwtsecretkey
        if err != nil {
                http.Error(w, "Failed to generate token", http.StatusInternalServerError)
                return
        }

        output := LoginOutput{
                Token:  tokenString,
                UserID: user.ID,
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(output)
}

