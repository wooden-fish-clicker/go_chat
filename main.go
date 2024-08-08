package main

import (
	"log"
	"net/http"

	"github.com/wooden-fish-clicker/chat/configs"
	"github.com/wooden-fish-clicker/chat/internal/websocket"
	"github.com/wooden-fish-clicker/chat/pkg/db"
	"github.com/wooden-fish-clicker/chat/pkg/logger"
	"github.com/wooden-fish-clicker/chat/pkg/redis"
)

func init() {
	configs.Setup()
	logger.Setup()
	db.ConnectMongoDB()
	redis.ConnectRedis()
}

func chatPage(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/chat" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "chat.html")
}
func main() {
	hub := websocket.NewHub()
	go hub.Run()
	http.HandleFunc("/ws/chat", func(w http.ResponseWriter, r *http.Request) {
		websocket.RecvFunc(hub, w, r)
	})

	http.HandleFunc("/chat", chatPage)
	log.Println("server start at: ", configs.C.App.ServerAddress)

	err := http.ListenAndServe(configs.C.App.ServerAddress, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
