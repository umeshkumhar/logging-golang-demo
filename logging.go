package main

import (
	"os"
	"regexp"
	"time"

	log "github.com/sirupsen/logrus"
)

type LogWriter struct{}

var levelRegex *regexp.Regexp

const (
	LevelError   = "error"
	LevelWarning = "warning"
	LevelFatal   = "fatal"
	LevelPanic   = "panic"
)

func init() {
	var err error
	levelRegex, err = regexp.Compile("level=([a-z]+)")
	if err != nil {
		log.WithError(err).Fatal("Cannot setup log level")
	}
}

func (w *LogWriter) detectLogLevel(p []byte) (level string) {
	matches := levelRegex.FindStringSubmatch(string(p))
	if len(matches) > 1 {
		level = matches[1]
	}
	return
}

func (w *LogWriter) Write(p []byte) (n int, err error) {
	level := w.detectLogLevel(p)

	if level == LevelError || level == LevelWarning || level == LevelFatal || level == LevelPanic {
		return os.Stderr.Write(p)
	}
	return os.Stdout.Write(p)
}


func main() {
	var amount, transaction int64
	log.SetOutput(&LogWriter{})
	_, err := os.Stat("file")
	if err != nil {
		log.WithError(err).Error("Transaction failed")
	}
	log.WithField("transaction", "newtranct").Debug("Debug Transaction processed successfully")

	for {
		amount = amount + 15
		transaction= transaction+ 12345
		log.WithFields(log.Fields{
			"user_id": 2384172,
			"transaction_id": transaction,
			"amount": amount,
			"currency": "USD",
		}).Info("Transaction processed successfully for $",amount)
		time.Sleep(2 * time.Second)
	}

}