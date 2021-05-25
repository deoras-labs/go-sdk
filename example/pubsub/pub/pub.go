package main

import (
	"context"
	"fmt"
	"io"
	"os"

	dapr "github.com/dapr/go-sdk/client"
	"go.uber.org/zap/zapcore"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var (
	// set the environment as instructions.
	pubsubName = os.Getenv("DAPR_PUBSUB_NAME")
	topicName  = "neworder"
	log        = ctrl.Log.WithName("dapr-client")
)

// setupLogger configures logger
func setupLogger(out io.Writer) {
	ctrl.SetLogger(
		zap.New(
			zap.Encoder(
				zapcore.NewJSONEncoder(
					zapcore.EncoderConfig{
						MessageKey:     "msg",
						LevelKey:       "level",
						TimeKey:        "time",
						NameKey:        "name",
						CallerKey:      "caller",
						StacktraceKey:  "stacktrace",
						LineEnding:     zapcore.DefaultLineEnding,
						EncodeLevel:    zapcore.LowercaseLevelEncoder,
						EncodeTime:     zapcore.ISO8601TimeEncoder,
						EncodeDuration: zapcore.SecondsDurationEncoder,
						EncodeCaller:   zapcore.ShortCallerEncoder,
						EncodeName:     zapcore.FullNameEncoder,
					},
				),
			),
			zap.UseDevMode(true),
			zap.WriteTo(out),
		),
	)
}

func main() {
	ctx := context.Background()
	data := []byte("ping")

	setupLogger(os.Stderr)
	client, err := dapr.NewClient(log)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	if err := client.PublishEvent(ctx, pubsubName, topicName, data); err != nil {
		panic(err)
	}
	fmt.Println("data published")

	fmt.Println("Done (CTRL+C to Exit)")
}
