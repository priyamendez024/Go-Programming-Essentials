// Chapter 26: Data‑Processing Pipeline
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/segmentio/kafka-go"
)

type Event struct {
    UserID string
    Type   string
}

type Summary struct {
    UserID string
    Count  int
}

func ingest(ctx context.Context, jobs chan<- Event) {
    reader := kafka.NewReader(kafka.ReaderConfig{
        Brokers: []string{"localhost:9092"},
        Topic:   "events",
        GroupID: "pipeline-ingest",
    })
    defer reader.Close()
    for {
        m, err := reader.FetchMessage(ctx)
        if err != nil {
            log.Println("ingest error:", err)
            continue
        }
        jobs <- Event{UserID: string(m.Key), Type: string(m.Value)}
        reader.CommitMessages(ctx, m)
    }
}

func filter(in <-chan Event, out chan<- Event) {
    for e := range in {
        if e.Type == "click" {
            out <- e
        }
    }
    close(out)
}

func aggregate(in <-chan Event, out chan<- Summary) {
    counts := make(map[string]int)
    for e := range in {
        counts[e.UserID]++
    }
    for user, cnt := range counts {
        out <- Summary{UserID: user, Count: cnt}
    }
    close(out)
}

func main() {
    ctx := context.Background()
    ch1 := make(chan Event, 100)
    ch2 := make(chan Event, 100)
    ch3 := make(chan Summary, 100)

    go ingest(ctx, ch1)
    go filter(ch1, ch2)
    go aggregate(ch2, ch3)

    for s := range ch3 {
        fmt.Println("Summary:", s)
    }
}
