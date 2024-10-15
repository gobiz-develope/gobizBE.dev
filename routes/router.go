package routes

import (
	"gobizdevelop/controller/auth"
	"gobizdevelop/controller/market"
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

	router.HandleFunc("/add-market", market.AddMarket).Methods("POST")
	router.HandleFunc("/toko/{slug}/menu", market.GetMenuByMarket).Methods("GET")
	router.HandleFunc("/toko/menu/update", market.UpdateMenu).Methods("PUT")

	return router
}
