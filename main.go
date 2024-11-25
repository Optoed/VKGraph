package main

import (
	"VKGraph/src"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func main() {
	// Загружаем переменные из .env
	err := godotenv.Load(".env")
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
	src.InitVKClient(ACCESS_TOKEN)

	// Настройка маршрутов
	r := mux.NewRouter()
	r.HandleFunc("/friends/{userID}", src.GetFriendsHandler).Methods("GET")
	r.HandleFunc("/friends/{userIDa}/{userIDb}", src.BuildGraphHandler).Methods("GET")

	r.HandleFunc("/friends/info/{userIDa}/{userIDb}", src.PrintPathHandler).Methods("GET")

	// Сервинг статических файлов
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	fmt.Println("Server started on port 8081")
	http.ListenAndServe(":8081", r)
}
