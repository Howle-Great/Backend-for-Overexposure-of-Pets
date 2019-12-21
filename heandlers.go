package main

import (
	"log"
	"net/http"
	"time"
)

type Handler struct {
	Items ItemRepository
	Auth  AuthManager
}

func (h *Handler) HeandlerSignInPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	params := r.PostFormValue
	nickname, password := params("nickname"), params("password")
	log.Printf("Sign in: user: %s, password: %s", nickname, password)

	token, err := h.Auth.SignIn(&h.Items, &User{
		Nickname: nickname,
		Password: password,
	})

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(
			"Access denied! Wrong nickname or password.",
		))
		return
	}
	cookie := &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(1 * time.Hour),
	}
	http.SetCookie(w, cookie)
}
