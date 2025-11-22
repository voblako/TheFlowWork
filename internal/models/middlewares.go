package models

type LogMessage struct {
	Method  string `json:"method"`
	Path    string `json:"path"`
	Latency string `json:"latency"`
}
