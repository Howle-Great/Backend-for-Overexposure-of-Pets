package main

import (
	"github.com/gorilla/mux"
)

func ConfigurateHandlers(r *mux.Router, h *Handler) *mux.Router {
	s := r.PathPrefix("/v1").Subrouter()

	// Sign in the system
	s.HandleFunc("/signin", h.HeandlerSignInPost).Methods("POST")
	// s.HandleFunc("/signin", HeandlerSignIn).Methods("GET")

	// Show user profile
	// s.HandleFunc("/profile", HeandlerProfile).Methods("GET")
	// s.HandleFunc("/profile", HeandlerProfile).Methods("POST")

	// Show users and there rights
	// s.HandleFunc("/users/{limits, offset, company, inorder}", HeandlerUsers).Methods("GET")
	/*
		company - sort by company
		inorder - sort by having order
	*/

	// Show information aboute user
	// s.HandleFunc("/user", HeandlerUserCreate).Methods("POST")
	// s.HandleFunc("/user/{id}", HeandlerUser).Methods("GET")
	// s.HandleFunc("/user/{id}", HeandlerUserDelete).Methods("DELETE")
	// s.HandleFunc("/user/{id}", HeandlerUserEdit).Methods("PUT")

	// Show hotels
	// s.HandleFunc("/hotels/{limits, offset}", HeandlerHotels).Methods("GET")

	// Show information aboute hotel
	// s.HandleFunc("/hotel", HeandlerHotelCreate).Methods("POST")
	// s.HandleFunc("/hotel/{id}", HeandlerHotel).Methods("GET")
	// s.HandleFunc("/hotel/{id}", HeandlerHotelDelete).Methods("DELETE")
	// s.HandleFunc("/hotel/{id}", HeandlerHotelEdit).Methods("PUT")

	// Show companis
	// s.HandleFunc("/companis/{limits, offset}", HeandlerCompanis).Methods("GET")

	// Show information aboute company
	// s.HandleFunc("/company", HeandlerCompanyCreate).Methods("POST")
	// s.HandleFunc("/company/{id}", HeandlerCompany).Methods("GET")
	// s.HandleFunc("/company/{id}", HeandlerCompanyDelete).Methods("DELETE")
	// s.HandleFunc("/company/{id}", HeandlerCompanyEdit).Methods("PUT")

	// Create order
	// s.HandleFunc("/order", HandlerOrderCreate).Methods("POST")
	// s.HandleFunc("/order/{id}", HandlerOrder).Methods("GET")
	// s.HandleFunc("/order/{id}", HandlerOrderDelete).Methods("DELETE")
	// s.HandleFunc("/order/{id}", HandlerOrderEdit).Methods("PUT")

	return s
}
