package logwrapper

import (
	"github.com/sirupsen/logrus"
)

// StandardLogger encapsulates logging format
type StandardLogger struct {
	*logrus.Logger
}

// Event represents a logged event
type Event struct {
	id      int
	message string
}

// New initialize standard logger
func New() *StandardLogger {
	baseLogger := logrus.New()
	standardLogger := &StandardLogger{baseLogger}
	standardLogger.Formatter = &logrus.JSONFormatter{}

	return standardLogger
}

var serviceResponse = Event{1, "Sender service responded with %d. Body: %s"}

// ServiceResponse logs responses from 3rd party services
func (l *StandardLogger) ServiceResponse(body string, statusCode int) {
	l.Errorf(serviceResponse.message, statusCode, body)
}
