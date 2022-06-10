package logs

import (
	"fmt"
	"log"
	"os"
	"time"
)

const LOGFILE = "logs.log"

type Log struct{}

func (l *Log) WriteLog(group, typeMessage, message string) {
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()
	messageData := fmt.Sprintf(
		"%s [%s] [%s] %s\n",
		time.Now().UTC(),
		group,
		typeMessage,
		message,
	)
	if _, err := f.WriteString(messageData); err != nil {
		log.Println(err)
	}
}
