package db

import "time"

type Node struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	NodeName  string    `gorm:"not null;unique" json:"username"`
	Password  string    `json:"password"`
	Key       string    `json:"key"`
	Chunks    []*Chunk  `gorm:"many2many:chunk_nodes;"`
}
