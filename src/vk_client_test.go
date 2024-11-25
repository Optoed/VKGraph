package src

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func GetAccessToken() (string, error) {
	// Загружаем переменные из .env
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return "", err
	}

	ACCESS_TOKEN := os.Getenv("ACCESS_TOKEN")
	if ACCESS_TOKEN == "" {
		fmt.Println("ACCESS_TOKEN is not set")
		return "", err
	}

	return ACCESS_TOKEN, nil
}

func TestVkClient1(t *testing.T) {
	accessToken, err := GetAccessToken()
	if err != nil {
		t.Error(err)
	}
	// Инициализируем VK API
	InitVKClient(accessToken)

	userIdOptoed := 265240894

	friendsOptoed, err := GetFriendIDs(userIdOptoed)
	if err != nil {
		t.Error(err)
	}

	t.Logf("count of Optoed's friends: %d", len(friendsOptoed))
}
