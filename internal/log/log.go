package log

import (
	"sync"

	"go.uber.org/zap"
)

var (
	logger *zap.Logger
	once   sync.Once
	Sugar  *zap.SugaredLogger
)

func init() {
	once.Do(func() {
		logger = zap.NewExample()
		Sugar = logger.Sugar()
	})
}
