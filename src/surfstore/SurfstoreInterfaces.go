package surfstore

type Block struct {
	BlockData []byte
	BlockSize int
}

type MigrationInstruction struct {
	LowerIndex int
	UpperIndex int
	DestAddr   string
}

type FileMetaData struct {
	Filename      string
	Version       int
	BlockHashList []string
}

type MetaStoreInterface interface {
	// Retrieves the server's FileInfoMap
	GetFileInfoMap(succ *bool, serverFileInfoMap *map[string]FileMetaData) error

	// Update a file's fileinfo entry
	UpdateFile(fileMetaData *FileMetaData, latestVersion *int) (err error)

	// Retrieve the mapping of BlockStore addresses to block hashes
	GetBlockStoreMap(blockHashesIn []string, blockStoreMap *map[string][]string) error

	// Add a BlockStore node
	AddNode(nodeAddr string, succ *bool) error

	// Remove a BlockStore node
	RemoveNode(nodeAddr string, succ *bool) error
}

type BlockStoreInterface interface {

	// Get a block based on
	GetBlock(blockHash string, block *Block) error

	// Put a block
	PutBlock(block Block, succ *bool) error

	// Provide a block hash and the result will be stored in hasBlock boolean value
	HasBlocks(blockHashesIn []string, blockHashesOut *[]string) error
}

type ClientInterface interface {
	// MetaStore
	GetFileInfoMap(succ *bool, serverFileInfoMap *map[string]FileMetaData) error
	UpdateFile(fileMetaData *FileMetaData, latestVersion *int) (err error)
	GetBlockStoreMap(blockHashesIn []string, blockStoreMap *map[string][]string) error

	// BlockStore
	GetBlock(blockHash string, blockStoreAddr string, block *Block) error
	PutBlock(block Block, blockStoreAddr string, succ *bool) error
	HasBlocks(blockHashesIn []string, blockStoreAddr string, blockHashesOut *[]string) error
}

type AdminInterface interface {
	AddNode(nodeAddr string, succ *bool) error
	RemoveNode(nodeAddr string, succ *bool) error
}
