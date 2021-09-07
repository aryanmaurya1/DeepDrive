package db

import (
	"time"
)

type Chunk struct {
	ID                 uint64    `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt          time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt          time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	ChunkOriginalName  string    `gorm:"size:121;not null" json:"chunk_original_name"`
	ChunkEncryptedName string    `gorm:"size:121;not null" json:"chunk_encrypted_name"`
	Hash               string    `json:"hash"`
	FileID             uint64    `json:"file_id"`
	File               File
	UserID             string  `json:"user_id"`
	Size               int     `json:"size"`
	Index              int     `json:"index"`
	Nodes              []*Node `gorm:"many2many:chunk_nodes;"`
}
