package market

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"gobizdevelop/config"
	"gobizdevelop/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Fungsi untuk mengupdate toko berdasarkan ID
func UpdateMarketByID(w http.ResponseWriter, r *http.Request) {
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

	// Mengurai JSON dari body request
	var updatedMarket model.Market
	err = json.NewDecoder(r.Body).Decode(&updatedMarket)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Mengatur field yang akan diupdate
	updateFields := bson.M{
		"nama_toko": updatedMarket.MarketName,
		"alamat":    updatedMarket.Alamat,
		"updatedAt": time.Now(),
	}

	// Koneksi ke koleksi MongoDB market
	collection := config.Mongoconn.Collection("market")

	// Membuat context dengan timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Melakukan update berdasarkan ID
	result, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": objID},
		bson.M{"$set": updateFields},
	)
	if err != nil {
		http.Error(w, "Failed to update market", http.StatusInternalServerError)
		return
	}

	// Memeriksa apakah data berhasil diupdate
	if result.MatchedCount == 0 {
		http.Error(w, "Market not found", http.StatusNotFound)
		return
	}

	// Mengirimkan respons sukses
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Market updated successfully",
	})
}
