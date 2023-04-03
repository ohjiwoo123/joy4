package setuplogger

import (
	"fmt"
	"io"
	"os"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	log "github.com/sirupsen/logrus"
)

const LOG_NAME = "rtmp_serv.log"

func SetupLogger() {
	dir, err := os.Getwd()

	if err != nil {
		log.Errorln("Failed Getwd(): %s", err)
		return
	}
	/**
	* @brief : rotatelogs의 New 함수
	* @details : New creates a new RotateLogs object. A log filename pattern
				 must be passed. Optional `Option` parameters may be passed
	* @param args : (pattern string, options ...Option)
	* @return : (*RotateLogs, error)
	*/
	logf, err := rotatelogs.New(
		fmt.Sprintf("%s/logs/%s.%s", dir, LOG_NAME, "%Y%m%d%H%M%S"),
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	// Fork writing into two outputs
	multiWriter := io.MultiWriter(os.Stderr, logf)

	logFormatter := new(log.TextFormatter)
	logFormatter.TimestampFormat = "2006-01-02 15:04:05.000"
	logFormatter.FullTimestamp = true

	log.SetFormatter(logFormatter)
	/**
	* @brief : logrus의 SetLevel 함수
	* @details : Trace, Debug, Info, Warn, Error, Fatal, Panic Level이 있다.
				 Debug Level은 Info, Warning, Error Fatal 4가지 레벨이 포함된다.
	* @param args : logger *Logger
	* @return : level
	*/
	//log.SetLevel(log.DebugLevel)
	log.SetOutput(multiWriter)
}
