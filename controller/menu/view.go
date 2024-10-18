package menu

import (
	"context"
	"encoding/json"
	"gobizdevelop/config"
	"gobizdevelop/model"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetMenuByMarket(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	slug := params["slug"]

	collection := config.Mongoconn.Collection("market")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var market model.Market
	err := collection.FindOne(ctx, bson.M{"slug": slug}).Decode(&market)
	if err != nil {
			http.Error(w, "Market not found", http.StatusNotFound)
			return
	}

	json.NewEncoder(w).Encode(market.Menu)
}

func GetMenuByID(w http.ResponseWriter, r *http.Request) {
	menuID := r.URL.Query().Get("menu_id")

	// Konversi menuID dari string ke primitive.ObjectID
	objectID, err := primitive.ObjectIDFromHex(menuID)
	if err != nil {
			http.Error(w, "Invalid menu ID format", http.StatusBadRequest)
			return
	}

	// Cari menu berdasarkan menu_id dari database
	var menu model.Menu
	collection := config.Mongoconn.Collection("market") // Pastikan koleksi yang benar
	err = collection.FindOne(context.TODO(), bson.M{"menu._id": objectID}).Decode(&menu)

	if err != nil {
			http.Error(w, "Menu not found", http.StatusNotFound)
			return
	}

	json.NewEncoder(w).Encode(menu)
}

