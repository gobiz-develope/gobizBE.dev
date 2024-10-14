package auth

import (
	"context"
	"encoding/json"
	"gobizdevelop/config"
	"gobizdevelop/model"
	"log"
	"net/http"
	"time"

	"github.com/o1egl/paseto"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

var symmetricKey = []byte("this_is_a_32_byte_secret_key_123")

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

	// Generate PASETO token
	now := time.Now()
	expiration := now.Add(24 * time.Hour) // Token valid for 24 hours

	token := paseto.NewV2() // Gunakan Paseto V2
	jsonToken := paseto.JSONToken{
		Subject:    user.ID.Hex(),
		IssuedAt:   now,
		Expiration: expiration,
	}
	footer := "some-footer-info"

	// Encrypt the token using the symmetric key
	encryptedToken, err := token.Encrypt(symmetricKey, jsonToken, footer)
	if err != nil {
		log.Println("Error generating token:", err, "|", "metric:", symmetricKey, "|", "jsontoken:", jsonToken) // Logging the error
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Send the response with the token
	response := map[string]interface{}{
		"message":  "Login successful",
		"user_id":  user.ID.Hex(),
		"email":    user.Email,
		"username": user.Nama,
		"token":    encryptedToken,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
