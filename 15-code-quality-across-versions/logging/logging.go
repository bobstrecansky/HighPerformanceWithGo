package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	c := zap.NewProductionConfig()
	c.OutputPaths = []string{"stdout"}
	logger, _ := c.Build()
	logger.Debug("We can use this logging level to debug.  This won't be printed, as the NewProduction logger only prints info and above log levels.")
	logger.Info("This is an INFO message for your code.  We can log indiviual structured things here", zap.String("url", "https://reddit.com"), zap.Int("connectionAttempts", 3), zap.Time("requestTime", time.Now()))
	logger.Warn("This is a WARNING message for your code.  It will not exit your program.")
	logger.Error("This is an ERROR message for your code.  It will not exit your program, but it will print your error message -> ")
	logger.Fatal("This is a Fatal message for your code.  It will exit your program with an os.Exit(1).")
	logger.Panic("This is a panic message for your code.  It will exit your program. We won't see this execute because we have already exited from the above logger.Fatal log message. This also exits with an os.Exit(1)")
}
