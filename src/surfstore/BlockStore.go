package surfstore

import (
	"net/rpc"
)

type BlockStore struct {
	BlockMap map[string]Block
	RingSize int
}

// Get the BlockMap of the BlockStore for debugging with run-debug.sh
func (bs *BlockStore) GetBlockMap(succ *bool, serverBlockInfoMap *map[string]Block) error {
	for k, v := range bs.BlockMap {
		(*serverBlockInfoMap)[k] = v
	}

	return nil
}

func (bs *BlockStore) GetBlock(blockHash string, blockData *Block) error {
	block, exist := bs.BlockMap[blockHash]

	if exist {
		blockData.BlockData = block.BlockData
		blockData.BlockSize = block.BlockSize
	} else {
		blockData = nil
	}

	return nil
}

func (bs *BlockStore) PutBlock(block Block, succ *bool) error {
	blockHash := GetBlockHashString(block.BlockData)
	bs.BlockMap[blockHash] = block
	*succ = true

	return nil
}

func (bs *BlockStore) hasBlock(blockHash string, hasBlock *bool) error {
	_, *hasBlock = bs.BlockMap[blockHash]

	return nil
}

//Given a list of hashes “in”, returns a list containing the
//subset of in that are stored in the key-value store
func (bs *BlockStore) HasBlocks(blockHashesIn []string, blockHashesOut *[]string) error {
	hasBlocksSlice := make([]string, 0)

	for _, blockHash := range blockHashesIn {
		hasBlock := false
		_ = bs.hasBlock(blockHash, &hasBlock)

		if hasBlock {
			hasBlocksSlice = append(hasBlocksSlice, blockHash)
		}
	}

	blockHashesOut = &hasBlocksSlice

	return nil
}

// Helper function for modulo operation
func mod(a int, b int) int {
	m := a % b
	if a < 0 && b < 0 {
		m -= b
	}
	if a < 0 && b > 0 {
		m += b
	}
	return m
}

// Migrate specified blocks from this node to another node.
func (bs *BlockStore) MigrateBlocks(inst MigrationInstruction, succ *bool) error {
	// connect to the server
	conn, e := rpc.DialHTTP("tcp", inst.DestAddr)
	if e != nil {
		return e
	}

	// migrate the blocks with ring index between inst.LowerIndex and inst.UpperIndex (in modulo sense)
	// in this BlockStore server to another BlockStore server with address inst.DestAddr
	// For each block to migrate, you could do:
	// e = conn.Call("BlockStore.PutBlock", block, succ)
	panic("todo")

	// close the connection
	return conn.Close()
}

// This line guarantees all method for BlockStore are implemented
var _ BlockStoreInterface = new(BlockStore)

func NewBlockStore(ringSize int) BlockStore {
	return BlockStore{
		BlockMap: map[string]Block{},
		RingSize: ringSize,
	}
}
