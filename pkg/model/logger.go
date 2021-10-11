package model

import "time"

type Logger struct {
	ID        int64      `json:"id" db:"id"`
	Transport string     `json:"transport" db:"transport"`
	Logging   string     `json:"logging" db:"logging"`
	Level     string     `json:"level" db:"level"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}
