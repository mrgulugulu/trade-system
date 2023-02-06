package log

import (
	"os"
	"sync"
	"trade-system/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
	once   sync.Once
	Sugar  *zap.SugaredLogger
)

func init() {
	once.Do(func() {
		encoder := getEncoder()
		writerSyncer := getLogWriter()
		core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)
		logger = zap.New(core)
		Sugar = logger.Sugar()
	})
	defer Sugar.Sync()
}

// getEncoder 使用zap默认的json编码器
func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

// getLogWriter 指定日志文件路径
func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create(config.LogFilePath)
	return zapcore.AddSync(file)
}
