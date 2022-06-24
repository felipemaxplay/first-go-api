package response

import "time"

type ErrorPlayer struct {
	Timestamp time.Time `json:"timestamp"`
	Code      int       `json:"code"`
	Error     string    `json:"error"`
	Message   string    `json:"message"`
}

func BuildPlayerError(code int, err string, message string) ErrorPlayer {
	res := ErrorPlayer{
		Timestamp: time.Now(),
		Code:      code,
		Error:     err,
		Message:   message,
	}

	return res
}
