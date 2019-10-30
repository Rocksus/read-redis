package consumer

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Rocksus/read-redis/modules/messaging"
	"github.com/Rocksus/read-redis/modules/rediscount"
	"github.com/nsqio/go-nsq"
)

const (
	defaultConsumerMaxAttempts = 10
	defaultConsumerMaxInFlight = 100
)

// Handler struct handles redis
type Handler struct {
	RDC *rediscount.Handler
}

// RunConsumer runs a goroutine
func (h *Handler) RunConsumer() {
	// initiate consumer
	cfg := messaging.ConsumerConfig{
		Channel:       os.Getenv("redisChannel"),
		LookupAddress: os.Getenv("lookupAddr"),
		Topic:         os.Getenv("redisTopic"),
		MaxAttempts:   defaultConsumerMaxAttempts,
		MaxInFlight:   defaultConsumerMaxInFlight,
		Handler:       h.HandleMessage,
	}
	consumer := messaging.NewConsumer(cfg)

	// run consumer
	consumer.Run()

	// keep app alive until terminated
	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	select {
	case <-term:
		log.Println("Application terminated")
	}
}

// HandleMessage handles consumer data
func (h *Handler) HandleMessage(message *nsq.Message) error {
	log.Println("Received call to increment visitor count")
	h.RDC.UpdateCount()
	return nil
}
