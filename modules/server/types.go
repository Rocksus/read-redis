package server

import (
	"github.com/Rocksus/read-redis/modules/rediscount"
	"github.com/Rocksus/read-redis/modules/userdata"
)

// Handler handles rediscounthandler and userdatahandler
type Handler struct {
	RDC *rediscount.Handler
	UDT *userdata.Handler
}
