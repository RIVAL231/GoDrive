package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rival231/Go-Drive/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"

	// "github.com/rival231/internal/models"
	// "github.com/go-chi/chi/v5"
	"github.com/rival231/Go-Drive/internal/db"
	// "github.com/go-chi/chi/v5/middleware"
)


func UserAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {
		
		// Extract user ID from URL parameters
      username, password, ok := r.BasicAuth()
	  if !ok {
		  http.Error(w, "Unauthorized", http.StatusUnauthorized)
		  return
	  }
       user,err := GetUserByUsername(username)
	   if err != nil {
		   http.Error(w, "User not found"+err.Error(), http.StatusNotFound)
		   return
	   }
	   err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	   if err != nil {
		   http.Error(w, "Invalid credentials"+err.Error(), http.StatusUnauthorized)
		   return
	   }
	   ctx := context.WithValue(r.Context(), "user", user)
        next.ServeHTTP(w, r.WithContext(ctx))

	}))
}

func JWTAuthMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == ""{
			  http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
            return
		}
		tokenString := authHeader[len("Bearer "):]
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        })
		if err != nil || !token.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }
		 ctx := context.WithValue(r.Context(), "username", claims["username"])
        next.ServeHTTP(w, r.WithContext(ctx))
	})
	
}

func GetUserByUsername(username string) (models.User, error) {
	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	userCollection := db.GetCollection("Go-Drive","users")
	var user models.User
	fmt.Println("Fetching user:", username)
	err := userCollection.FindOne(ctx,bson.M{"username":username}).Decode(&user)
	return user,err

}