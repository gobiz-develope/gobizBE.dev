package controller

import (
	"context"
	"encoding/json"
	"gobizdevelop/config"
	"gobizdevelop/model"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func LoginUsers(w http.ResponseWriter, r *http.Request) {
	var credentials model.Users
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Setup a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get the MongoDB collection
	collection := config.Mongoconn.Collection("users")

	// Find the user by email
	var user model.Users
	err = collection.FindOne(ctx, bson.M{"email": credentials.Email}).Decode(&user)
	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Compare the provided password with the hashed password stored in the database
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	// If the password is correct, create a response
	response := map[string]interface{}{
		"message":  "Login successful",
		"user_id":  user.ID,
		"email":    user.Email,
		"username": user.Nama,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
