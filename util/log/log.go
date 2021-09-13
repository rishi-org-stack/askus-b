package log

import (
	"os"
	"time"

	"github.com/golang/glog"
	"github.com/sirupsen/logrus"
)

type (
	message struct {
		source string
		level  string
		mess   interface{}
	}
)

const (
	infoLog  = "INFO: "
	warnLog  = "WARN: "
	errorLog = "Error: "
)

var (
	Log *logrus.Logger
)

func init() {
	Log = &logrus.Logger{
		Level:     logrus.InfoLevel,
		Out:       os.Stdout,
		Formatter: &logrus.JSONFormatter{},
	}
}
func Init(source, level string) *message {
	return &message{
		source: source,
		level:  level,
	}
}
func (m *message) Mut(mess interface{}) *message {
	m.mess = mess
	return m
}
func (m *message) Info() {
	glog.Infoln(
		infoLog+"\n\tDATE: ",
		time.Now().Day(),
		time.Now().Month(),
		time.Now().Year(),
		" time : ",
		time.Now().Hour(),
		time.Now().Minute(),
		"\n\tMessage :\n", "\t\t",
		m.source,
		"\n\t\t",
		m.level,
		"\n\t\t",
		m.mess)
}
func (m *message) Error() {
	glog.Errorln(
		errorLog+"\n\tDATE: ",
		time.Now().Day(),
		time.Now().Month(),
		time.Now().Year(),
		" time : ",
		time.Now().Hour(),
		time.Now().Minute(),
		"\n\tMessage :\n", "\t\t",
		m.source,
		"\n\t\t",
		m.level,
		"\n\t\t",
		m.mess)
}
func (m *message) Warn() {
	glog.Warningln(
		warnLog+"\n\tDATE: ",
		time.Now().Day(),
		time.Now().Month(),
		time.Now().Year(),
		" time : ",
		time.Now().Hour(),
		time.Now().Minute(),
		"\n\tMessage :\n", "\t\t",
		m.source,
		"\n\t\t",
		m.level,
		"\n\t\t",
		m.mess)
}
