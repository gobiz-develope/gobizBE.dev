package controller

import (
	"encoding/json"
	"gobizdevelop/config"
	"gobizdevelop/model"
	"net/http"
	"time"

	"context"

	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/writeconcern"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUsers(w http.ResponseWriter, r *http.Request) {
	var user model.Users

	// Decode the JSON request body into the user struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Hash the user's password before saving it to the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// Set other user fields like creation time
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// Insert the user data into the MongoDB collection
	collection := config.Mongoconn.Collection("users")

	// Setup a context with timeout to avoid hanging in case of connection issues
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Insert the user into the users collection
	result, err := collection.InsertOne(ctx, bson.M{
		"username":   user.Nama,
		"password":   user.Password,
		"email":      user.Email,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
	})
	if err != nil {
		http.Error(w, "Failed to insert user", http.StatusInternalServerError)
		return
	}

	// Send back the ID of the newly created user as a response
	response := map[string]interface{}{
		"message": "User registered successfully",
		"user_id": result.InsertedID,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
