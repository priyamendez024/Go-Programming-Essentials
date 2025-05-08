// Chapter 15: NoSQL & ORMs in Go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"

    "github.com/go-redis/redis/v8"
    "go.etcd.io/bbolt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func mongoExample() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        log.Fatal(err)
    }
    col := client.Database("myapp").Collection("users")
    res, _ := col.InsertOne(ctx, bson.M{"name": "Alice", "email": "alice@example.com"})
    fmt.Println("Inserted ID:", res.InsertedID)
}

func redisExample() {
    ctx := context.Background()
    rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
    rdb.Set(ctx, "key", "value", 0)
    val, _ := rdb.Get(ctx, "key").Result()
    fmt.Println("Redis key:", val)
}

func boltExample() {
    db, _ := bbolt.Open("my.db", 0666, nil)
    defer db.Close()
    db.Update(func(tx *bbolt.Tx) error {
        b, _ := tx.CreateBucketIfNotExists([]byte("MyBucket"))
        return b.Put([]byte("key"), []byte("value"))
    })
}

func gormExample() {
    dsn := "host=localhost user=pguser dbname=mydb password=secret"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }
    type User struct {
        ID    uint
        Name  string
        Email string
    }
    db.AutoMigrate(&User{})
    db.Create(&User{Name: "Bob", Email: "bob@example.com"})
}

func main() {
    mongoExample()
    redisExample()
    boltExample()
    gormExample()
}
