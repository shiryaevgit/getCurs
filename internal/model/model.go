package model

import "time"

type CurrencyData struct {
	ID             int
	Currency       string
	Price          string
	DayHighPrice   string
	DayLowPrice    string
	HourChange     string
	LastUpdateTime time.Time
}
