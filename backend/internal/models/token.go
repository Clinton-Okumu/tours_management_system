package models

import (
	"time"
)

type Token struct {
	BaseModel
	Plaintext string    `json:"token"`
	Hash      []byte    `json:"-"`
	UserID    uint      `json:"-"`
	Expiry    time.Time `json:"expiry"`
	Scope     string    `json:"-"`
}
