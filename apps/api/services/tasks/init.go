package tasks

import (
	"github.com/hibiken/asynq"
	"log"
)

var Client *asynq.Client

const redisAddr = "127.0.0.1:6379"

func Init() {
	Client = asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})

	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{
			Concurrency: 10,
		},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(TypeEmailForgotPassword, HandleEmailForgotPasswordTask)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run asynq server: %v", err)
	}
}

func Close() {
	err := Client.Close()

	if err != nil {
		log.Fatalf("could not close asynq client: %v", err)
	}
}
