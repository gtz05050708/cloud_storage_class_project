package main

import (
	"flag"
	"fmt"
	"net/rpc"
	"os"
	"surfstore"
)

// Usage String
const USAGE_STRING = "./run-debug.sh -r <ring_size> <BlockStoreAddr>"

const ARG_COUNT = 1
const EX_USAGE int = 64

func main() {
	// Custom flag Usage message
	flag.Usage = func() {
		w := flag.CommandLine.Output()
		fmt.Fprintf(w, "Usage of %s:\n", USAGE_STRING)
		flag.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(w, "  -%s: %v\n", f.Name, f.Usage)
		})
	}

	ringSize := flag.Int("r", 128, "(default = 128) Consistent hashing ring size")
	flag.Parse()
	args := flag.Args()

	if len(args) != ARG_COUNT {
		flag.Usage()
		os.Exit(EX_USAGE)
	}

	hostPort := args[0]

	fmt.Println(hostPort)
	// connect to the server
	conn, e := rpc.DialHTTP("tcp", hostPort)
	if e != nil {
		return
	}

	blockMap := make(map[string]surfstore.Block)
	// perform the call
	succ := false
	e = conn.Call("BlockStore.GetBlockMap", succ, &blockMap)
	if e != nil {
		conn.Close()
		return
	}

	PrintBlockMap(blockMap, *ringSize)

	// close the connection
	conn.Close()
}

func PrintBlockMap(blockMap map[string]surfstore.Block, ringSize int) {

	fmt.Println("--------BEGIN PRINT MAP--------")

	// fmt.Println(blockMap)

	for blockHash := range blockMap {
		fmt.Println("\t", blockHash, int(surfstore.HashMod(blockHash, ringSize)))
	}

	fmt.Println("---------END PRINT MAP--------")

}
