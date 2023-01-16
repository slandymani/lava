package logger

import (
	"sync"

	"github.com/lavanet/lava/relayer/common/allowList"
	"github.com/rs/zerolog"
)

var (
	instance *Logger
	once     sync.Once
)

const (
	NoPairingAvailableError string = "No pairings available."
)

type LogMessage struct {
	Description string         // a string describing the log message
	Err         error          // an error associated with the log message
	LogEvent    *zerolog.Event // log level
}

var epochErrors = []string{NoPairingAvailableError}

type Logger struct {
	logChan             chan LogMessage // channel to send log messages
	epochErrorAllowList *allowList.AllowList
}

// GetInstance is a function that creates a singleton instance of the Logger struct
// and returns it.
func GetInstance() *Logger {
	once.Do(func() {
		instance = &Logger{
			logChan:             make(chan LogMessage, 1024), // the channel buffer size is 1024
			epochErrorAllowList: allowList.NewErrorAllowList(epochErrors),
		}
		go instance.listen()
	})
	return instance
}

// listen function to listen on the log channel, it will keep running and waiting for new messages
// on the channel and it will print the logs once it receives any log message
func (l *Logger) listen() {
	for {
		// wait for a message
		msg := <-l.logChan

		// if error is not nil and is inside allow-list don't log the message
		if msg.Err != nil && l.epochErrorAllowList.IsErrorSet(msg.Err.Error()) {
			continue
		}

		// if the error is not nil check if it needs to be added in the allow-list
		if msg.Err != nil {
			l.addErrorInAllowList(msg.Err)
		}

		// log the message
		l.printLogs(msg.Description, msg.LogEvent)
	}
}

// addErrorInAllowList adds an error in the epoch error allow-list if needed
func (l *Logger) addErrorInAllowList(err error) {
	// Make sure that error is not already in allow-list
	if l.epochErrorAllowList.IsErrorSet(err.Error()) {
		return
	}

	// If error is inside epoch errors add it into allow-list
	if l.isInsideEpochErrors(err.Error()) {
		l.epochErrorAllowList.SetError("No pairings available.")
	}
}

// ResetErrorAllowList resets epoch error allow-list
func (l *Logger) ResetErrorAllowList() {
	l.epochErrorAllowList.Reset()
}

// isInsideEpochErrors checks if the error is inside epochErrors
func (l *Logger) isInsideEpochErrors(error string) bool {
	for _, a := range epochErrors {
		if a == error {
			return true
		}
	}
	return false
}

// printLogs function is to print the log messages
func (l *Logger) printLogs(description string, logEvent *zerolog.Event) {
	logEvent.Msg(description)
}

// Log function is to push the log messages to the channel
func (l *Logger) Log(msg LogMessage) {
	l.logChan <- msg
}