package routes

import (
	"gobizdevelop/controller/auth"
	"gobizdevelop/controller/market"
	"gobizdevelop/controller/menu"
	"gobizdevelop/controller/profile"

	"github.com/gorilla/mux"
)

// InitializeRoutes sets up the router
func InitializeRoutes() *mux.Router {
	router := mux.NewRouter()

	// Define your routes here
	router.HandleFunc("/regis", auth.RegisterUsers).Methods("POST")
	router.HandleFunc("/login", auth.LoginUsers).Methods("POST")

	router.HandleFunc("/profile", profile.GetUsers).Methods("POST")
	router.HandleFunc("/profile-update", profile.UpdateUser).Methods("POST")

	router.HandleFunc("/markets", market.GetMarkets).Methods("GET")
	router.HandleFunc("/add-market", market.AddMarket).Methods("POST")
	router.HandleFunc("/market-id", market.GetMarketByID).Methods("GET")
	router.HandleFunc("/market/update", market.UpdateMarketByID).Methods("PUT")
	router.HandleFunc("/market/delete", market.DeleteMarketByID).Methods("DELETE")

	router.HandleFunc("/toko/menu", menu.AddMenuToToko).Methods("POST")
	router.HandleFunc("/toko/{slug}/menu", menu.GetMenuByMarket).Methods("GET")
	router.HandleFunc("/menu-id", menu.GetMenuByID).Methods("GET")
	router.HandleFunc("/toko/menu/update", menu.UpdateMenu).Methods("PUT")
	router.HandleFunc("/toko/{slug}/menu", menu.DeleteMenu).Methods("DELETE")

	return router
}
