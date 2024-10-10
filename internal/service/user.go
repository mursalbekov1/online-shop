package service

import (
	"encoding/json"
	"net/http"
	"online-shop/internal/models"
	"online-shop/internal/repository/postgres"
	"strconv"
)

type UserHandler struct {
	UserRepo postgres.UserRepository
}

func (h *UserHandler) GetUser(w http.ResponseWriter, req *http.Request) {
	idstr := req.Header.Get("id")
	if idstr == "" {
		http.Error(w, "Missing id header", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.UserRepo.GetById(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, req *http.Request) {
	users, err := h.UserRepo.GetAll()
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, req *http.Request) {
	var user models.User
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	if err := h.UserRepo.Create(&user); err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, req *http.Request) {
	var user models.User
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	idstr := req.Header.Get("id")
	if idstr == "" {
		http.Error(w, "Missing id header", http.StatusBadRequest)
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
	}

	user.Id = id
	if err := h.UserRepo.Update(&user); err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, req *http.Request) {
	idstr := req.Header.Get("id")
	if idstr == "" {
		http.Error(w, "Missing id header", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if err := h.UserRepo.Delete(id); err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
