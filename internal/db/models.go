package db

import (
	"time"
)

type User struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Name      string    `json:"name"`
	Username  string    `gorm:"not null;unique" json:"username"`
	Key       string    `json:"key"`
	Password  string    `json:"password"`
}

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

type Node struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	NodeName  string    `gorm:"not null;unique" json:"username"`
	Password  string    `json:"password"`
	Key       string    `json:"key"`
	Chunks    []*Chunk  `gorm:"many2many:chunk_nodes;"`
}

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
