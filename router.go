package main

import (
	"github.com/gorilla/mux"
)

func ConfigurateHandlers(r *mux.Router) *mux.Router {
	s := r.PathPrefix("/v1").Subrouter()

	// Sign in in system
	s.HandleFunc("/signin", HeandlerSignIn).Methods("POST")

	// Sign up in system
	s.HandleFunc("/signup/user", HeandlerSignUpUser).Methods("POST")
	s.HandleFunc("/signup/hotel/{step}", HeandlerSignUpHotel).Methods("POST")
	s.HandleFunc("/signup/company/{step}", HeandlerSignUpCompany).Methods("POST")

	// Show user profile
	s.HandleFunc("/profile", HeandlerProfile).Methods("GET")
	s.HandleFunc("/profile", HeandlerProfile).Methods("POST")

	// Show users and there rights
	s.HandleFunc("/users/{limits, offset, company, inorder}", HeandlerUsers).Methods("GET")
	/*
		company - sort by company
		inorder - sort by having order
	*/

	// Show information aboute user
	s.HandleFunc("/user/{id}", HeandlerUser).Methods("GET")
	s.HandleFunc("/user/delete/{id}", HeandlerUserDelete).Methods("DELETE")
	s.HandleFunc("/user/create", HeandlerUserCreate).Methods("POST")
	s.HandleFunc("/user/edit/{id}", HeandlerUserEdit).Methods("PUT")

	// Show hotels
	s.HandleFunc("/hotels/{limits, offset}", HeandlerHotels).Methods("GET")

	// Show information aboute hotel
	s.HandleFunc("/hotel/{id}", HeandlerHotel).Methods("GET")
	s.HandleFunc("/hotel/delete/{id}", HeandlerHotelDelete).Methods("DELETE")
	s.HandleFunc("/hotel/create", HeandlerHotelCreate).Methods("POST")
	s.HandleFunc("/hotel/edit/{id}", HeandlerHotelEdit).Methods("PUT")

	// Show companis
	s.HandleFunc("/companis/{limits, offset}", HeandlerCompanis).Methods("GET")

	// Show information aboute company
	s.HandleFunc("/company/{id}", HeandlerCompany).Methods("GET")
	s.HandleFunc("/company/delete/{id}", HeandlerCompanyDelete).Methods("DELETE")
	s.HandleFunc("/company/create", HeandlerCompanyCreate).Methods("POST")
	s.HandleFunc("/company/edit/{id}", HeandlerCompanyEdit).Methods("PUT")

	// Create order
	s.HandleFunc("/order/{id}", HandlerOrder).Methods("GET")
	s.HandleFunc("/order/delete/{id}", HandlerOrderDelete).Methods("DELETE")
	s.HandleFunc("/order/create", HandlerOrderCreate).Methods("POST")
	s.HandleFunc("/order/edit/{id}", HandlerOrderEdit).Methods("PUT")

	return s
}
