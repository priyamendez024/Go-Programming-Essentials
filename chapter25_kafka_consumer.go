// Chapter 25: Kafka Event Consumer
package main

import (
    "context"
    "encoding/json"
    "log"

    "github.com/segmentio/kafka-go"
)

type OrderCreated struct {
    OrderID     string  `json:"order_id"`
    CustomerID  string  `json:"customer_id"`
    TotalAmount float64 `json:"total_amount"`
    Timestamp   int64   `json:"timestamp"`
}

func main() {
    reader := kafka.NewReader(kafka.ReaderConfig{
        Brokers: []string{"localhost:9092"},
        Topic:   "orders",
        GroupID: "order-service-group",
    })
    defer reader.Close()

    for {
        msg, err := reader.ReadMessage(context.Background())
        if err != nil {
            log.Println("error:", err)
            continue
        }
        var event OrderCreated
        json.Unmarshal(msg.Value, &event)
        log.Printf("processing order %s for customer %s", event.OrderID, event.CustomerID)
    }
}
