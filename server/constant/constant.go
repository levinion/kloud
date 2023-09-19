package constant

import (
	"log/slog"
)

const (
	ServerIP  = "localhost:9090"
	DBPath    = "db/store"
	LogPath   = "log.log"
	LogLevel  = slog.LevelDebug
	RedisAddr = "localhost:"
)
