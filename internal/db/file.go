package db

import (
	"time"
)

type File struct {
	ID           uint64    `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	UserID       uint64    `json:"user_id"`
	User         User
	OriginalName string `json:"original_name"`
	Hash         string `json:"hash"`
	Size         int    `json:"size"`
	ChunkCount   int    `json:"chunk_count"`
}
