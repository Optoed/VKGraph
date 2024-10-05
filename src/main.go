package main

import (
	"encoding/json"
	"fmt"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"strconv"
)

type Friend struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Photo string `json:"photo"`
	Sex   int    `json:"sex"` //1 - male, 2 - female
}

var vk *api.VK

func getFriendIDs(userID int) ([]int, error) {
	friendsParams := params.NewFriendsGetBuilder()
	friendsParams.UserID(userID)

	response, err := vk.FriendsGet(friendsParams.Params)
	if err != nil {
		return nil, err
	}

	return response.Items, nil
}

func getUsersDetails(userIDs []int) ([]Friend, error) {
	var ids []string
	for _, id := range userIDs {
		ids = append(ids, strconv.Itoa(id))
	}

	usersParams := params.NewUsersGetBuilder()
	usersParams.UserIDs(ids)
	usersParams.Fields([]string{"photo_50", "sex"})

	response, err := vk.UsersGet(usersParams.Params)
	if err != nil {
		return nil, err
	}

	var friends []Friend
	for _, user := range response {
		friends = append(friends, Friend{
			ID:    user.ID,
			Name:  fmt.Sprintf("%s %s", user.FirstName, user.LastName),
			Photo: user.Photo50,
			Sex:   user.Sex,
		})
	}

	return friends, nil
}

func GetFriendsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["userID"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	friendIDs, err := getFriendIDs(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	friends, err := getUsersDetails(friendIDs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(friends)
}

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	ACCESS_TOKEN := os.Getenv("ACCESS_TOKEN")
	if ACCESS_TOKEN == "" {
		fmt.Println("ACCESS_TOKEN is not set")
		return
	}

	fmt.Println("Starting...")

	vk = api.NewVK(ACCESS_TOKEN)

	r := mux.NewRouter()
	r.HandleFunc("/friends/{userID}", GetFriendsHandler).Methods("GET")

	http.ListenAndServe(":8080", r)
}
