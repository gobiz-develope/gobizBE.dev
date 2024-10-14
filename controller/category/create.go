package category

import (
	"context"
	"encoding/json"
	"gobizdevelop/config"
	"gobizdevelop/model"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category model.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	category.ID = primitive.NewObjectID()

	collection := config.Mongoconn.Collection("category")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = collection.InsertOne(ctx, category)
	if err != nil {
		http.Error(w, "Failed to create category", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(category)
}
