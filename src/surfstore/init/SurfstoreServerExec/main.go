package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"strconv"
	"strings"
	"surfstore"
)

// Usage String
const USAGE_STRING = "./run-server.sh -s <service_type> -p <port> -r <ring_size> -l -d (BlockStoreAddr*)"

// Set of valid services
var SERVICE_TYPES = map[string]bool{"meta": true, "block": true, "both": true}

// Exit codes
const EX_USAGE int = 64

func main() {
	// Custom flag Usage message
	flag.Usage = func() {
		w := flag.CommandLine.Output()
		fmt.Fprintf(w, "Usage of %s:\n", USAGE_STRING)
		flag.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(w, "  -%s: %v\n", f.Name, f.Usage)
		})
		fmt.Fprintf(w, "  (BlockStoreAddr*): BlockStore Addresses (include self if service type is both)\n")
	}

	// Parse command-line argument flags
	service := flag.String("s", "", "(required) Service Type of the Server: meta, block, both")
	port := flag.Int("p", 8080, "(default = 8080) Port to accept connections")
	localOnly := flag.Bool("l", false, "Only listen on localhost")
	debug := flag.Bool("d", false, "Output log statements")
	ringSize := flag.Int("r", 128, "(default = 128) Consistent hashing ring size")
	flag.Parse()

	// Use tail arguments to hold variable number of BlockStore addresses
	blockStoreAddrs := flag.Args()

	// Valid service type argument
	if _, ok := SERVICE_TYPES[strings.ToLower(*service)]; !ok {
		flag.Usage()
		os.Exit(EX_USAGE)
	}

	// Add localhost if necessary
	addr := ""
	if *localOnly {
		addr += "localhost"
	}
	addr += ":" + strconv.Itoa(*port)

	// Disable log outputs if debug flag is missing
	if !(*debug) {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}

	log.Fatal(startServer(addr, strings.ToLower(*service), *ringSize, blockStoreAddrs))
}

func startServer(hostAddr string, serviceType string, ringSize int, blockStoreAddrs []string) error {
	// Create a new Server
	rpcServer := rpc.NewServer()

	// Register rpc services
	if serviceType != "block" {
		metastore := surfstore.NewMetaStore(surfstore.NewConsistentHashRing(ringSize, blockStoreAddrs))
		rpcServer.RegisterName("MetaStore", &metastore)
	}

	if serviceType != "meta" {
		blockstore := surfstore.NewBlockStore(ringSize)
		rpcServer.RegisterName("BlockStore", &blockstore)
	}

	l, e := net.Listen("tcp", hostAddr)
	if e != nil {
		return e
	}

	return http.Serve(l, rpcServer)
}
