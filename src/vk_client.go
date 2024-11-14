package src

import (
	"fmt"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"strconv"
)

var Vk *api.VK

// Инициализация VK клиента
func InitVKClient(accessToken string) {
	Vk = api.NewVK(accessToken)
}

// Получение списка ID друзей
func GetFriendIDs(userID int) ([]int, error) {
	friendsParams := params.NewFriendsGetBuilder()
	friendsParams.UserID(userID)

	response, err := Vk.FriendsGet(friendsParams.Params)
	if err != nil {
		return nil, err
	}

	return response.Items, nil
}

// Получение детальной информации о друзьях
func GetUsersDetails(userIDs []int) ([]Friend, error) {
	var ids []string
	for _, id := range userIDs {
		ids = append(ids, strconv.Itoa(id))
	}

	usersParams := params.NewUsersGetBuilder()
	usersParams.UserIDs(ids)
	usersParams.Fields([]string{"photo_50", "sex"})

	response, err := Vk.UsersGet(usersParams.Params)
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
