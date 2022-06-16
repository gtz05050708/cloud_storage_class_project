package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"surfstore"
)

// Usage String
const USAGE_STRING = "./run-admin.sh -s <service_type> <MetaStoreAddr> <BlockStoreAddr>"

// Set of valid services
var SERVICE_TYPES = map[string]bool{"add": true, "remove": true}

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
	}

	service := flag.String("s", "", "(required) Admin Service: add or remove")
	flag.Parse()

	// Valid service type argument
	if _, ok := SERVICE_TYPES[strings.ToLower(*service)]; !ok {
		flag.Usage()
		os.Exit(EX_USAGE)
	}

	args := flag.Args()
	metaHostPort := args[0]
	blockHostPort := args[1]

	rpcAdmin := surfstore.NewSurfstoreRPCAdmin(metaHostPort)
	succ := false
	var err error
	if *service == "add" {
		err = rpcAdmin.AddNode(blockHostPort, &succ)
	} else if *service == "remove" {
		err = rpcAdmin.RemoveNode(blockHostPort, &succ)
	}
	if err != nil {
		log.Fatal(err)
	}
}
