package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "7777777",
		DB:       0,
	})

	duration := 0 * time.Second

	// SIMPLE VALUE
	err := client.Set("key", "value1", duration).Err()
	if err != nil {
		log.Fatal(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(val)
	//////////////////////////////// END SIMPLE VALUE

	// MAP VALUE
	user := map[string]interface{}{
		"name":    "John Doe",
		"age":     30,
		"country": "USA",
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	err = client.HSet("user:100", "person", userJSON).Err()
	if err != nil {
		log.Fatal(err)
	}

	result, err := client.HGetAll("user:100").Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
	/////////////////////////////////// END MAP VALUE

	// STRUCT VALUE

	user1 := User{ID: 1, Name: "John Doe"}

	err = client.Set("user:1", user1, 0).Err()
	if err != nil {
		log.Fatal(err)
	}

	err = client.Get("user:1").Scan(&user1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user1)
	/////////////////////////////////// END STRUCT VALUE
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (u User) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

func (u User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &u)
}
