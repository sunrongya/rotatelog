package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/sunrongya/rotatelog"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
	log.Formatter = new(logrus.JSONFormatter)
	
	log.Hooks.Add(rotatelog.NewHook(rotatelog.PathMap{
		logrus.DebugLevel : `debug.log`,
		logrus.InfoLevel :  `info.log`,
		logrus.WarnLevel :  `warning.log`,
		logrus.ErrorLevel : `error.log`,
		logrus.FatalLevel : `fatal.log`,
		logrus.PanicLevel : `panic.log`,
	}) )
	log.Level = logrus.DebugLevel
}

func main() {
	defer func() {
		err := recover()
		if err != nil {
			log.WithFields(logrus.Fields{
				"omg":    true,
				"err":    err,
				"number": 100,
			}).Fatal("The ice breaks!")
		}
	}()

	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"number": 8,
	}).Debug("Started observing beach")

	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(logrus.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(logrus.Fields{
		"temperature": -4,
	}).Debug("Temperature changes")

	log.WithFields(logrus.Fields{
		"animal": "orca",
		"size":   9009,
	}).Panic("It's over 9000!")
}

