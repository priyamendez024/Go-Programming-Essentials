// Chapter 24: Real‑World Chat Server
package main

import (
    "context"
    "log"
    "net/http"

    "github.com/go-redis/redis/v8"
    "github.com/gorilla/websocket"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{Addr: "localhost:6379"})
var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
var broadcast = make(chan []byte)
var clients = make(map[*websocket.Conn]bool)

func handleMessages() {
    sub := rdb.Subscribe(ctx, "chat")
    ch := sub.Channel()
    for msg := range ch {
        for client := range clients {
            client.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
        }
    }
}

func handleWS(w http.ResponseWriter, r *http.Request) {
    conn, _ := upgrader.Upgrade(w, r, nil)
    clients[conn] = true
    defer conn.Close()
    for {
        _, msg, err := conn.ReadMessage()
        if err != nil {
            delete(clients, conn)
            break
        }
        rdb.Publish(ctx, "chat", msg)
    }
}

func main() {
    go handleMessages()
    http.HandleFunc("/ws", handleWS)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
