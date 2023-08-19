package main

import (
	naivelog "github.com/SolarSword/NaiveLog"
)

func main() {
	logger := naivelog.StdLogger()
	logger.Info("test logger")
	logger.SetOptions(naivelog.WithLevel(naivelog.Debug))
	logger.Debug("debug test logger")
	logger.SetOptions(naivelog.WithLevel(naivelog.Warn))
	logger.Debug("by right i shouldn't be printed")

	logger.SetOptions(naivelog.WithFormatter(&naivelog.TextFormatter{IgnoreBasicFields: false}))
	logger.Warn("warn test")
}
