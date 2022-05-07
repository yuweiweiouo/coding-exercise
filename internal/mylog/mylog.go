package mylog

import (
	"os"
	"path"

	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Provider = wire.NewSet(New, NewOption)

type Option struct {
	Level      string
	Stdout     bool
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
}

func NewOption(v *viper.Viper) (*Option, error) {
	option := &Option{
		Level:    "info",
		Filename: "./logs/web.log",
	}
	if err := v.UnmarshalKey("log", option); err != nil {
		return nil, err
	}

	return option, nil
}

func New(option *Option) (*zap.Logger, error) {
	os.MkdirAll(path.Dir(option.Filename), os.ModePerm)
	level, _ := zap.ParseAtomicLevel(option.Level)

	fw := zapcore.AddSync(&lumberjack.Logger{
		Filename:   option.Filename,
		MaxSize:    option.MaxSize,
		MaxBackups: option.MaxBackups,
		MaxAge:     option.MaxAge,
	})

	cw := zapcore.Lock(os.Stdout)

	cores := make([]zapcore.Core, 0, 2)
	je := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	cores = append(cores, zapcore.NewCore(je, fw, level))

	if option.Stdout {
		ce := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		cores = append(cores, zapcore.NewCore(ce, cw, level))
	}

	core := zapcore.NewTee(cores...)
	logger := zap.New(core)

	zap.ReplaceGlobals(logger)

	return logger, nil
}
