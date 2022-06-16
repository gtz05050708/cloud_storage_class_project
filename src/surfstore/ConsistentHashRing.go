package surfstore

import (
	"crypto/sha256"
	"encoding/hex"
	"math/big"
	"sort"
)

type Node struct {
	Addr  string
	Index int
}

type ConsistentHashRing struct {
	RingSize int
	Nodes    []Node
}

// Perform a modulo operation on a hash string.
// The hash string is assumed to be hexadecimally encoded.
func HashMod(hashString string, ringSize int) int {
	hashBytes, _ := hex.DecodeString(hashString)
	hashInt := new(big.Int).SetBytes(hashBytes[:])
	ringSizeInt := big.NewInt(int64(ringSize))

	indexInt := new(big.Int).Mod(hashInt, ringSizeInt)

	return int(indexInt.Int64())
}

// Compute a block’s index on the ring from its hash value.
func (ms *ConsistentHashRing) ComputeBlockIndex(blockHash string) int {
	return HashMod(blockHash, ms.RingSize)
}

// Compute a node’s index on the ring from its address string.
func (ms *ConsistentHashRing) ComputeNodeIndex(nodeAddr string) int {
	hashBytes := sha256.Sum256([]byte(nodeAddr))
	hashString := hex.EncodeToString(hashBytes[:])
	return HashMod(hashString, ms.RingSize)
}

// Find the hosting node for the given ringIndex. It’s basically the first node on the ring with node.Index >= ringIndex (in a modulo sense).
func (ms *ConsistentHashRing) FindHostingNode(ringIndex int) Node {
	// Try to implement a O(log N) solution here using binary search.
	// It's also fine if you can't because we don't test your perforrmance.
	//panic("todo")
	end := len(ms.Nodes) - 1
	start := 0
	// if block index is in between the first and last node, use first node
	if (ringIndex > ms.Nodes[end].Index) {
		return ms.Nodes[0]
	}
	// else find the node in linear sense
	bestDistance := ms.Nodes[end].Index - ringIndex
	successorNode := ms.Nodes[end]
	for {
		if (start > end) {
			break
		}
		middle := (end - start) / 2
		middle += start
		middleNode := ms.Nodes[middle]
		distance := middleNode.Index - ringIndex
		if (distance < 0) {
			start = middle + 1
			continue
		} else {
			end = middle - 1
		}
		if (distance < bestDistance) {
			bestDistance = distance
			successorNode = middleNode
		}
	}
	return successorNode
}

// Add the given nodeAddr to the ring.
func (ms *ConsistentHashRing) AddNode(nodeAddr string) {
	// O(N) solution is totally fine here.
	// O(log N) solution might be overly complicated.
	//panic("todo")
	nodeIdx := ms.ComputeNodeIndex(nodeAddr)
	newNode := Node{Index: nodeIdx, Addr: nodeAddr}
	ms.Nodes = append(ms.Nodes, newNode)
	sort.Slice(ms.Nodes, func(i, j int) bool { 
		return ms.Nodes[i].Index < ms.Nodes[j].Index 
	})
}

// Remove the given nodeAddr from the ring.
func (ms *ConsistentHashRing) RemoveNode(nodeAddr string) {
	// O(N) solution is totally fine here.
	// O(log N) solution might be overly complicated.
	nodeIdx := ms.ComputeNodeIndex(nodeAddr)
	if (len(ms.Nodes) == 1 && ms.Nodes[0].Index == nodeIdx) {
		ms.Nodes = make([]Node, 0)
	}
	nodes := ms.Nodes
	end := len(ms.Nodes) - 1
	start := 0
	idx := 0
	for {
		if (start > end) {
			break
		}
		middle := (end - start) / 2
		middle += start
		if (ms.Nodes[middle].Index == nodeIdx) {
			idx = middle
			break
		} else if (ms.Nodes[middle].Index > nodeIdx) {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}
	ms.Nodes = append(nodes[:idx], nodes[idx+1:]...)
}

// Create consistent hash ring struct with a list of blockstore addresses
func NewConsistentHashRing(ringSize int, blockStoreAddrs []string) ConsistentHashRing {
	// You can not use ComputeNodeIndex method to compute the ring index of blockStoreAddr in blockStoreAddrs here.
	// You will need to use HashMod function, remember to hash the blockStoreAddr before calling HashMod
	// Hint: refer to ComputeNodeIndex method on how to hash the blockStoreAddr before calling HashMod
	//panic("todo")
	var storeNodes []Node
	for _, nodeAddr := range (blockStoreAddrs) {
		hashBytes := sha256.Sum256([]byte(nodeAddr))
		hashString := hex.EncodeToString(hashBytes[:])
		nodeIdx := HashMod(hashString, ringSize)
		newNode := Node{Index: nodeIdx, Addr: nodeAddr}
		storeNodes = append(storeNodes, newNode)
	}
	sort.Slice(storeNodes, func(i, j int) bool { 
		return storeNodes[i].Index < storeNodes[j].Index 
	})
	return ConsistentHashRing{RingSize: ringSize, Nodes: storeNodes}
}
