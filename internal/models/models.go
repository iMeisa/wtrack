package models

import (
	"time"
)

type DB struct {
	Types   []Strain `json:"types"`
	Records []Record `json:"records"`
}

type Strain struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Amount float32 `json:"amount"`
}

type Record struct {
	ID     int       `json:"id"`
	Date   time.Time `json:"date"`
	Amount float32   `json:"amount"`
	Strain `json:"strain"`
}
