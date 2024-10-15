package auth

import (
	"context"
	"encoding/json"
	"gobizdevelop/config"
	"gobizdevelop/helper/watoken"
	"gobizdevelop/model"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

var PrivateKey = "e4cb06d20bcce42bf4ac16c9b056bfaf1c6a5168c24692b38eb46d551777dc4147db091df55d64499fdf2ca85504ac4d320c4c645c9bef75efac0494314cae94"

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

	// Encrypt the token using the symmetric key
	encryptedToken, err := watoken.EncodeforHours(user.No_Telp, user.Nama, PrivateKey, 18)
	if err != nil {
		log.Println("Error generating token:", err, "|", "metric:", PrivateKey, "|", "jsontoken:", user.Nama) // Logging the error
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Send the response with the token
	response := map[string]interface{}{
		"message":  "Login successful",
		"user_id":  user.ID.Hex(),
		"email":    user.Email,
		"username": user.Nama,
		"role":     user.Role,
		"token":    encryptedToken,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
