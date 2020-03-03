package main

import "go.uber.org/zap"

func main() {
	zapLogger := zap.NewExample()
	defer zapLogger.Sync()
	zapLogger.Debug("Hi Gophers - from our Zap Logger")
}
