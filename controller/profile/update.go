package profile

import (
	"context"
	"encoding/json"
	"gobizdevelop/config"
	"gobizdevelop/model"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UpdateUser updates the user's data in the MongoDB collection
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Get user_id from URL params (e.g., /users/{id})
	vars := r.URL.Query()
	userID := vars.Get("id")

	// Convert the user_id from string to MongoDB ObjectID
	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Parse the body to get the updated data
	var updatedUser model.Users
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Set the updated timestamp
	updatedUser.UpdatedAt = time.Now()

	// Set up the MongoDB collection
	collection := config.Mongoconn.Collection("users")

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Define the update query using $set to update only specific fields
	update := bson.M{
		"$set": bson.M{
			"nama":       updatedUser.Nama,
			"no_telp":    updatedUser.No_Telp,
			"email":      updatedUser.Email,
			"alamat":     updatedUser.Alamat,
			"updated_at": updatedUser.UpdatedAt,
		},
	}

	// Perform the update operation
	result, err := collection.UpdateOne(ctx, bson.M{"_id": oid}, update)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	// Check if any document was actually updated
	if result.MatchedCount == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Return success response
	response := map[string]interface{}{
		"message": "User updated successfully",
		"user_id": userID,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
