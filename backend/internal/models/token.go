package models

import (
	"time"

	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	Plaintext string    `json:"token"`
	Hash      []byte    `json:"-"`
	UserID    uint      `json:"-"`
	Expiry    time.Time `json:"expiry"`
	Scope     string    `json:"-"`
}
