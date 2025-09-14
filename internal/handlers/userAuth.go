package handlers

import (
	"context"
	"net/http"
	"time"

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
		   http.Error(w, "User not found", http.StatusNotFound)
		   return
	   }
	   err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	   if err != nil {
		   http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		   return
	   }
	   ctx := context.WithValue(r.Context(), "user", user)
        next.ServeHTTP(w, r.WithContext(ctx))

	}))
}

func GetUserByUsername(username string) (models.User, error) {
	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	userCollection := db.GetCollection("Go-Drive","users")
	var user models.User
	err := userCollection.FindOne(ctx,bson.M{"username":username}).Decode(&user)
	return user,err

}