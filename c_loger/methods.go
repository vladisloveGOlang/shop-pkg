package c_loger

import (
	"github.com/sirupsen/logrus"
	"os"
)

type StLog struct {
	Func   string
	Action string
	Status string
	Msg    interface{}
}

var (
	Started   = "started"
	Completed = "completed"
	Failed    = "failed"
)

func NewLog(funcName string) StLog {
	return StLog{
		Func: funcName,
	}
}

func (l StLog) Act(actionName string) StLog {
	resLog := l
	resLog.Action = actionName
	return resLog
}

func (l StLog) Started() StLog {
	resLog := l
	resLog.Status = Started
	return resLog
}

func (l StLog) Failed() StLog {
	resLog := l
	resLog.Status = Failed
	return resLog
}

func (l StLog) Completed() StLog {
	resLog := l
	resLog.Status = Completed
	return resLog
}

func (l StLog) Info() {

	log := logrus.New()
	log.SetOutput(os.Stdout)
	//log.SetFormatter(&logrus.JSONFormatter{}) // Для JSON-формата (удобнее для парсинга)

	log.WithFields(logrus.Fields{
		"func":    l.Func,
		"action":  l.Action,
		"status":  l.Status,
		"message": l.Msg,
	}).Info("")
}

func (l StLog) Error(err error) {
	resLog := l
	resLog.Status = Failed
	logrus.Error(err)
}
