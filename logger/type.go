package logger

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	Layout = "2006-01-02 15:04"
)

type RequestMessage map[string]any

type Timestamp = string

type Logger struct {
	Level         string         `json:"level"`
	RequestURL    string         `json:"requestURL,omitempty"`
	RequestMethod string         `json:"requestMethod,omitempty"`
	Message       RequestMessage `json:"message,omitempty,omitzero"`
	Timestamp     Timestamp      `json:"timestamp,omitempty"`
}

func TimeNow[T Timestamp]() T {
	return T(time.Now().Format(Layout))
}

func (l *Logger) MarshalJSON() ([]byte, error) {
	return json.Marshal(&Logger{
		Level:         l.Level, //TODO: сделать проверку на валидность уровня
		RequestURL:    l.RequestURL,
		RequestMethod: l.RequestMethod,
		Message:       l.Message,
		Timestamp:     TimeNow[Timestamp](),
	})
}

func (l *Logger) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &Logger{
		Level:         l.Level,
		RequestURL:    l.RequestURL,
		RequestMethod: l.RequestMethod,
		Message:       l.Message,
		Timestamp:     TimeNow[Timestamp](),
	})
}

func (l Logger) String() string {
	return fmt.Sprintf("%s %s %s %s %s", l.Level, l.RequestURL, l.RequestMethod, l.Message, l.Timestamp)
}
