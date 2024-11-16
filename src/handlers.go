package src

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetFriendsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["userID"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	friendIDs, err := GetFriendIDs(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	friends, err := GetUsersDetails(friendIDs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(friends)
}

func BuildGraphHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDa, err := strconv.Atoi(vars["userIDa"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userIDb, err := strconv.Atoi(vars["userIDb"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	path, err := bidirectionalSearch(userIDa, userIDb, GetFriendIDs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(path)
}

func PrintPathHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDa, err := strconv.Atoi(vars["userIDa"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userIDb, err := strconv.Atoi(vars["userIDb"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	path, err := bidirectionalSearch(userIDa, userIDb, GetFriendIDs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pathDetails, err := GetUsersDetails(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pathDetails)
}
