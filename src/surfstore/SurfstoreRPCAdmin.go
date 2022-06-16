package surfstore

import (
	"fmt"
	"net/rpc"
)

type RPCAdmin struct {
	MetaStoreAddr string
}

func (surfAdmin *RPCAdmin) AddNode(nodeAddr string, succ *bool) error {
	// connect to the server
	fmt.Println("[surfAdmin.AddNode]", surfAdmin.MetaStoreAddr, nodeAddr)
	conn, e := rpc.DialHTTP("tcp", surfAdmin.MetaStoreAddr)
	if e != nil {
		return e
	}

	// perform the call
	e = conn.Call("MetaStore.AddNode", nodeAddr, succ)
	if e != nil {
		conn.Close()
		return e
	}

	// close the connection
	return conn.Close()
}

func (surfAdmin *RPCAdmin) RemoveNode(nodeAddr string, succ *bool) error {
	// connect to the server
	conn, e := rpc.DialHTTP("tcp", surfAdmin.MetaStoreAddr)
	if e != nil {
		return e
	}

	// perform the call
	e = conn.Call("MetaStore.RemoveNode", nodeAddr, succ)
	if e != nil {
		conn.Close()
		return e
	}

	// close the connection
	return conn.Close()
}

var _ AdminInterface = new(RPCAdmin)

// Create an Surfstore RPC client
func NewSurfstoreRPCAdmin(hostPort string) RPCAdmin {
	return RPCAdmin{
		MetaStoreAddr: hostPort,
	}
}
