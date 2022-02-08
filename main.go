package main

import (
	"go-forum-thingy/database"
	"go-forum-thingy/webhandler"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"log/syslog"
	"os"
)

func logInit() {
	logLevel := "error"
	syslogger, err := syslog.New(syslog.LOG_INFO|syslog.LOG_DEBUG|syslog.LOG_WARNING|syslog.LOG_ERR, "TEKKSTATT")
	syslogwriter := zerolog.SyslogLevelWriter(syslogger)
	multi := zerolog.MultiLevelWriter(syslogwriter, os.Stdout)
	log.Logger = zerolog.New(multi).With().Timestamp().Caller().Logger()
	level, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		level = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(level)
}

func main() {
	logInit()
	database.Connect()
	webhandler.NewApp()

}
