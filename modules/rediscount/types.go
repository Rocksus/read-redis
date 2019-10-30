package rediscount

import "github.com/gomodule/redigo/redis"

// Handler struct
type Handler struct {
	Client redis.Conn
}
