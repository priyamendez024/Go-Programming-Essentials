// Chapter 25: Kafka Event Producer
package main

import (
    "context"
    "encoding/json"
    "log"
    "time"

    "github.com/segmentio/kafka-go"
)

type OrderCreated struct {
    OrderID     string  `json:"order_id"`
    CustomerID  string  `json:"customer_id"`
    TotalAmount float64 `json:"total_amount"`
    Timestamp   int64   `json:"timestamp"`
}

func main() {
    writer := kafka.NewWriter(kafka.WriterConfig{
        Brokers: []string{"localhost:9092"},
        Topic:   "orders",
    })
    defer writer.Close()

    event := OrderCreated{"1234", "cust-5678", 99.95, time.Now().Unix()}
    payload, _ := json.Marshal(event)
    writer.WriteMessages(context.Background(),
        kafka.Message{Key: []byte(event.OrderID), Value: payload},
    )
    log.Println("published OrderCreated event")
}
