package domain

import "github.com/lib/pq"

type MovieEntity struct {
	ID          string `gorm:"primaryKey"`
	Title       string
	Overview    string
	ReleaseDate string
	Image       string
	Duration    int16
	Budget      float64
	Genres      pq.StringArray `gorm:"type:text[]"`
}
