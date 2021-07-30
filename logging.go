package main
import (
	"os"
	"regexp"
	"time"
	uuid "github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/rand"
	"k8s.io/apimachinery/pkg/util/sets"
)
type LogWriter struct{}
var (
	levelRegex *regexp.Regexp
	userIDs    = sets.NewInt(787801, 787802, 787803, 787805)
)
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
	log.SetOutput(&LogWriter{})
	log.SetReportCaller(true)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
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
func transact(userID, amount int) {
	if userID == 787805 {
		log.WithFields(log.Fields{
			"user_id":  userID,
			"amount":   amount,
			"currency": "USD",
		}).Fatal("Couldn't find the user")
	}
	log.WithFields(log.Fields{
		"user_id":        userID,
		"transaction_id": uuid.New(),
		"amount":         amount,
		"currency":       "USD",
	}).Info("Transaction processed successfully")
}
func checkAndTransact(userID, amount int) {
	if !userIDs.Has(userID) {
		log.WithFields(log.Fields{
			"user_id":  userID,
			"amount":   amount,
			"currency": "USD",
		}).Warn("User is not allowed to perform transaction")
		return
	}
	if amount < 85 {
		log.WithFields(log.Fields{
			"user_id":  userID,
			"amount":   amount,
			"currency": "USD",
		}).Error("Transaction failed because of insufficient balance")
		return
	}
	transact(userID, amount)
}
func main() {

	log.Error("Error while processing")
	for {
		userID := rand.IntnRange(787801, 787806)
		amount := rand.IntnRange(80, 100)
		log.WithFields(log.Fields{
			"user_id":  userID,
			"amount":   amount,
			"currency": "USD",
		}).Info("User requested for transaction")
		checkAndTransact(userID, amount)
		time.Sleep(2 * time.Second)
	}
}