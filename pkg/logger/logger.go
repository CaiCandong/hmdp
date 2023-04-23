package logger

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitializeLogger() {
	Logger, _ = zap.NewProduction()
}

//func Zap() (logger *zap.Logger) {
//	if ok, _ := utils.PathExists(global.GVA_CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
//		fmt.Printf("create %v directory\n", global.GVA_CONFIG.Zap.Director)
//		_ = os.Mkdir(global.GVA_CONFIG.Zap.Director, os.ModePerm)
//	}
//	// 调试级别
//	debugPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
//		return lev == zap.DebugLevel
//	})
//	// 日志级别
//	infoPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
//		return lev == zap.InfoLevel
//	})
//	// 警告级别
//	warnPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
//		return lev == zap.WarnLevel
//	})
//	// 错误级别
//	errorPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
//		return lev >= zap.ErrorLevel
//	})
//
//	cores := [...]zapcore.Core{
//		getEncoderCore(fmt.Sprintf("./%s/server_debug.log", global.GVA_CONFIG.Zap.Director), debugPriority),
//		getEncoderCore(fmt.Sprintf("./%s/server_info.log", global.GVA_CONFIG.Zap.Director), infoPriority),
//		getEncoderCore(fmt.Sprintf("./%s/server_warn.log", global.GVA_CONFIG.Zap.Director), warnPriority),
//		getEncoderCore(fmt.Sprintf("./%s/server_error.log", global.GVA_CONFIG.Zap.Director), errorPriority),
//	}
//	logger = zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())
//
//	if global.GVA_CONFIG.Zap.ShowLine {
//		logger = logger.WithOptions(zap.AddCaller())
//	}
//	return logger
//}
