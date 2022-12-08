package models

import "time"

type Strain struct {
	ID     int
	Name   string
	Amount float32
}

type Record struct {
	Date   time.Time
	Amount float32
	Strain
}
