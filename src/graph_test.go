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
		3: {1, 5},
		4: {2},
		5: {3},
	}
	return mockData[userID], nil
}

func mockGetUsersDetails(userIDs []int) ([]Friend, error) {
	mockFriends := []Friend{
		{ID: 2, Name: "Alice", Photo: "photo2.jpg", Sex: 1},
		{ID: 3, Name: "Bob", Photo: "photo3.jpg", Sex: 2},
		{ID: 4, Name: "Charlie", Photo: "photo4.jpg", Sex: 1},
		{ID: 5, Name: "David", Photo: "photo5.jpg", Sex: 2},
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
	graph, path, err := BuildGraph(1, 5, mockGetFriendIDs, mockGetUsersDetails)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if _, ok := graph[1][2]; !ok {
		t.Errorf("Expected friend with ID 2 in graph[1]")
	}
	if _, ok := graph[3][5]; !ok {
		t.Errorf("Expected friend with ID 5 in graph[3]")
	}

	t.Logf("Graph: %+v", graph)
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

	optoed := 265240894
	pikmike := 7834725
	//через 1 друга

	friendsIDs, err := GetFriendIDs(optoed)
	if err != nil {
		t.Error(err)
	}

	t.Log(friendsIDs)

	_, path, err := BuildGraph(optoed, pikmike, GetFriendIDs, GetUsersDetails)
	if err != nil {
		t.Error(err)
	}

	t.Logf("path: %+v", path)
}
