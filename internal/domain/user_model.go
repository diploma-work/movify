package domain

import "github.com/lib/pq"

type UserEntity struct {
	ID       string `gorm:"primaryKey"`
	Name     string
	Email    string
	Password string
	Roles    pq.StringArray `gorm:"type:text[]"`
}
