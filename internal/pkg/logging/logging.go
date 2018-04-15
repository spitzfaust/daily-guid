package logging

import (
	"fmt"
	"log"
)

// LogWithContext logs to the given logger prepeding the context to the given message.
func LogWithContext(logger *log.Logger, context string, format string, v ...interface{}) {
	logger.Printf(fmt.Sprintf("[%s] ", context)+format, v...)
}
