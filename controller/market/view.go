package market

import (
	"context"
	"encoding/json"
	"gobizdevelop/config"
	"gobizdevelop/model"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Fungsi untuk mendapatkan daftar toko (market) saja
func GetMarkets(w http.ResponseWriter, r *http.Request) {
	// Koneksi ke koleksi MongoDB toko
	collection := config.Mongoconn.Collection("market")

	// Membuat context dengan timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Membuat opsi untuk mengabaikan field 'menu'
	projection := options.Find().SetProjection(bson.M{"menu": 0})

	// Menemukan semua toko, tanpa array 'menu'
	cursor, err := collection.Find(ctx, bson.M{}, projection)
	if err != nil {
		http.Error(w, "Failed to retrieve markets", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	// Dekode hasil query ke dalam slice dari model Toko
	var markets []model.Market
	if err = cursor.All(ctx, &markets); err != nil {
		http.Error(w, "Failed to parse markets", http.StatusInternalServerError)
		return
	}

	// Kirim daftar toko (market) sebagai respons
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(markets)
}

func GetMarketByID(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan ID toko dari parameter URL
	id := r.URL.Query().Get("id")

	// Memastikan ID tidak kosong
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	// Mengubah ID dari string ke ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// Koneksi ke koleksi MongoDB market
	collection := config.Mongoconn.Collection("market")

	// Membuat context dengan timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var market model.Market
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&market)
	if err != nil {
		http.Error(w, "Market not found", http.StatusNotFound)
		return
	}

	// Kirim data toko sebagai respons
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(market)
}
