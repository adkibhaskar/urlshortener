package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zapLog *zap.Logger
func init(){

	encoderConfig:=zap.NewProductionEncoderConfig()

	encoderConfig.EncodeTime=zapcore.ISO8601TimeEncoder

	fileEnoder:=zapcore.NewJSONEncoder(encoderConfig)

	file,err:=os.Create("logfile.log")

	if err != nil{
		panic(err)
	}

	// defer file.Close()

	fileCore:=zapcore.NewCore(fileEnoder,zapcore.AddSync(file),zap.DebugLevel)

	core:=zapcore.NewTee(fileCore)

	zapLog=zap.New(core)
}

func GetLogger() *zap.Logger{

	return zapLog

}




