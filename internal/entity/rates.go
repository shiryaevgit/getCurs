package entity

import "time"

type Rate struct {
	Id       int
	Currency string  `json:"currency" `
	Value    float64 `json:"value" `
	Time     time.Time
}
