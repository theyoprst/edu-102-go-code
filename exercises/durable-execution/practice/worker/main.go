package main

import (
	"log"
	"log/slog"
	"os"

	translation "temporal102/exercises/durable-execution/practice"

	"go.temporal.io/sdk/client"
	tlog "go.temporal.io/sdk/log"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{
		Logger: tlog.NewStructuredLogger(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))),
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, translation.TaskQueueName, worker.Options{})

	w.RegisterWorkflow(translation.SayHelloGoodbye)
	w.RegisterActivity(translation.TranslateTerm)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
