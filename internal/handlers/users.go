package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	// "github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rival231/Go-Drive/internal/db"
	"github.com/rival231/Go-Drive/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w, "http method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "couldnt decode user data", http.StatusBadRequest)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "couldnt hash password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)
	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	userCollection := db.GetCollection("Go-Drive","users")
	_,err = userCollection.InsertOne(ctx,bson.M{
		"username":user.Username,
		"password":user.Password,
	})
	if err != nil {
		http.Error(w, "couldnt create user"+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("user created successfully"))
}
var jwtKey = []byte("sankalp231")
func UserLogin(w http.ResponseWriter, r *http.Request){
	user, ok := r.Context().Value("user").(models.User)
	if !ok {
		http.Error(w, "user not found", http.StatusUnauthorized)
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"username":user.Username,
		"exp":time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString(jwtKey)
	 if err != nil {
        http.Error(w, "Could not generate token", http.StatusInternalServerError)
        return
    }

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string] interface{}{
		"message":"login successful",
		"token":tokenString,
		"user": user.Username,
	})
}