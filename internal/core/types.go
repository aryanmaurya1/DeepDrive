package core

type Node struct {
	NodeId int    `json:"node_id"`
	Name   string `json:"name"`
}

// User : Represents a single user who wants to store his files.
type User struct {
	UserId   int    `json:"user_id"`
	Password string `json:"password"`
	Seed     string `json:"seed"`
	Files    []File `json:"files"`
}

// File : A single file which user has uploaded, and wants to store on cloud.
type File struct {
	FileId         int         `json:"file_id"`
	Name           string      `json:"name"`
	NumberOfPieces int         `json:"number_of_pieces"`
	CreatedAt      int64       `json:"created_at"` // Unixnano time
	Chunks         []FileChunk `json:"chunks"`
}

// FileChunk : A single encrypted piece of file, stored on nodes.
type FileChunk struct {
	ChunkId       string `json:"chunck_id"`
	OriginalName  string `json:"original_name"`
	EncryptedName string `json:"encrypted_name"`
	Index         int    `json:"index"`
	NodeIds       []int  `json:"node_ids"` // NodeId of all the nodes which have this piece
}
