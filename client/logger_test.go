package client

import (
	"io"

	"go.uber.org/zap/zapcore"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var log = ctrl.Log.WithName("dapr-client")

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
