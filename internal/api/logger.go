package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/voblako/TheFlowWork/internal/models"
)

type Logger struct {
	handler http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	jsonMessage, err := json.Marshal(models.LogMessage{Method: r.Method, Path: r.URL.Path, Latency: time.Since(start).String()})
	if err != nil {
		log.Println("Error marshaling log message:", err)
		return
	}
	fmt.Println(string(jsonMessage))
}

func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handlerToWrap}
}
