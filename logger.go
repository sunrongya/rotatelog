package rotatelog

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/sunrongya/lumberjack"
	"log"
)

type PathMap map[logrus.Level]string

// Hook to handle writing to local log files.
type lfsHook struct {
	loggers  map[logrus.Level]*lumberjack.Logger
	levels  []logrus.Level
}

func NewHook(levelMap PathMap) *lfsHook {
	hook := &lfsHook{
		loggers: make(map[logrus.Level]*lumberjack.Logger, len(levelMap) ),
	}
	for level, filename := range levelMap {
		hook.loggers[level] = &lumberjack.Logger{
			Filename:   filename,
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     28,  // days
		}
		hook.levels = append(hook.levels, level)
	}
	
	return hook
}

func NewHookFor(loggers map[logrus.Level]*lumberjack.Logger) *lfsHook {
	hook := &lfsHook{ loggers: loggers }
	for level, _ := range loggers {
		hook.levels = append(hook.levels, level)
	}
	
	return hook
}

func (hook *lfsHook) Fire(entry *logrus.Entry) error {
	logger, ok := hook.loggers[entry.Level]
	if !ok {
		err := fmt.Errorf("no file provided for loglevel: %d", entry.Level)
		log.Println(err.Error())
		return err
	}
	msg, err := entry.String()
	if err != nil {
		log.Println("failed to generate string for entry:", err)
		return err
	}
	_, err = logger.Write([]byte(msg) )
	return err
}

func (hook *lfsHook) Levels() []logrus.Level {
	return hook.levels
}
