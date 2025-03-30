package main

import "time"

type URL struct {
	ID          uint   `gorm:"primaryKey"`
	ShortID     string `gorm:"unique"`
	OriginalUrl string
	Clicks      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
