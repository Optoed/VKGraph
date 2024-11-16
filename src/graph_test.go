// src/pkg/graph_test.go
package src

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func mockGetFriendIDs(userID int) ([]int, error) {
	mockData := map[int][]int{
		1: {2, 3},
		2: {1, 4},
		3: {1, 6},
		4: {2, 6},
		5: {6},
		6: {4, 5},
	}
	return mockData[userID], nil
}

func mockGetUsersDetails(userIDs []int) ([]Friend, error) {
	mockFriends := []Friend{
		{ID: 2, Name: "Alice", Photo: "photo2.jpg", Sex: 1},
		{ID: 3, Name: "Bob", Photo: "photo3.jpg", Sex: 2},
		{ID: 4, Name: "Charlie", Photo: "photo4.jpg", Sex: 1},
		{ID: 5, Name: "David", Photo: "photo5.jpg", Sex: 2},
		{ID: 6, Name: "Gosha", Photo: "photo6.jpg", Sex: 2},
	}

	var result []Friend
	for _, id := range userIDs {
		for _, friend := range mockFriends {
			if friend.ID == id {
				result = append(result, friend)
			}
		}
	}
	return result, nil
}

func TestBuildGraph1(t *testing.T) {
	path, err := bidirectionalSearch(1, 5, mockGetFriendIDs)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	t.Logf("Path: %+v", path)
}

func TestBuildGraph2(t *testing.T) {
	// Загружаем переменные из .env
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

	// Инициализируем VK API
	InitVKClient(ACCESS_TOKEN)

	userIda := 265240894
	userIdb := 178526820
	//через 1 друга

	path, err := bidirectionalSearch(userIda, userIdb, GetFriendIDs)
	if err != nil {
		t.Error(err)
	}

	t.Logf("path: %+v", path)
}

func TestBuildGraph3(t *testing.T) {
	// Загружаем переменные из .env
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

	// Инициализируем VK API
	InitVKClient(ACCESS_TOKEN)

	userIda := 265240894
	userIdb := 2323226
	//через больше чем 1 друга

	path, err := bidirectionalSearch(userIda, userIdb, GetFriendIDs)
	if err != nil {
		t.Error(err)
	}

	t.Logf("path: %+v", path)
}

func TestBuildGraph4(t *testing.T) {
	// Загружаем переменные из .env
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

	// Инициализируем VK API
	InitVKClient(ACCESS_TOKEN)

	userIda := 6482392
	userIdb := 2323226
	//это друзья

	path, err := bidirectionalSearch(userIda, userIdb, GetFriendIDs)
	if err != nil {
		t.Error(err)
	}

	t.Logf("path: %+v", path)
}

func TestBuildGraph5(t *testing.T) {
	// Загружаем переменные из .env
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

	// Инициализируем VK API
	InitVKClient(ACCESS_TOKEN)

	userIda := 6482392
	userIdb := 265240894
	//через больше чем 1 друга

	path, err := bidirectionalSearch(userIda, userIdb, GetFriendIDs)
	if err != nil {
		t.Error(err)
	}

	t.Logf("path: %+v", path)
}
