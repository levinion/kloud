package logger

import (
	"kloud/constant"
	"log/slog"
	"os"
)

func Init() {
	file, err := os.OpenFile(constant.LogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}
	handler := slog.NewJSONHandler(file, &slog.HandlerOptions{Level: constant.LogLevel})
	slog.SetDefault(slog.New(handler))
}
