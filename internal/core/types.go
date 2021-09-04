package core

type Node struct {
	NodeId string `json:"node_id"`
	Name   string `json:"name"`
}

// User : Represents a single user who wants to store his files.
type User struct {
	UserId   string `json:"user_id"`
	Password string `json:"password"`
	Seed     string `json:"seed"`
	Files    []File `json:"files"`
}

// File : A single file which user has uploaded, and wants to store on cloud.
type File struct {
	FileId         string      `json:"file_id"`
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

// FileStages : Store the current state of file i.e what is currently happening with file.
type FileStages struct {

	// POSSIBLE STAGES
	// 	1. BREAKING (file is being broken)
	//  2. DISTRIBUTION (file is being distributed)
	//	3. INNETWORK (file is inside storage nodes)
	//	4. RECOLLECTION (file is being collected from network)
	//	5. MERGING (file is being merged)

	FileID         string // ID of the particular file in processing
	Stage          string // Current stage of the file
	PiecesComplete int    // Number of pieces which have passed a particular stage. Resets when stage changes.
	TotalPieces    int    // Total number of pieces a file is broken into.
}
