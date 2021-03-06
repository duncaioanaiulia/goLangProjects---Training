package main

import (
	"net/http"
	"regexp"
)

type userController struct{
	userIDPattern *regexp.Regexp
}

func (uc userController) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func newUserController() *userController{
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}