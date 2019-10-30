package rediscount

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
)

// Init function
func (h *Handler) Init() {
	redisString := os.Getenv("redisConn")
	redipool := redis.Pool{
		MaxActive:   20,
		MaxIdle:     5,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisString)
			if err != nil {
				return nil, err
			}
			return c, err
		},
	}
	conn := redipool.Get()
	h.Client = conn
	_, err := h.Client.Do("SET", "visitors", 0)
	if err != nil {
		fmt.Println("Error Redis Set!")
		log.Println(err)
	}
	fmt.Println("Successfully initialized redis")
}

// GetLatestCount function to access redis
func (h *Handler) GetLatestCount() int {
	val, err := redis.Int(h.Client.Do("GET", "visitors"))
	if err != nil {
		log.Println(err)
	}
	return val
}

// UpdateCount increments visitor count
func (h *Handler) UpdateCount() {
	h.Client.Do("INCR", "visitors")
}
