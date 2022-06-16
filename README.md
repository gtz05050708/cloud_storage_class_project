# Scalable SurfStore

This is the starter code for Module 4: Scalable Surfstore.

## Data Types

In this project, we add the following type in `SurfstoreInterfaces.go`:

```go
type MigrationInstruction struct {
	LowerIndex int
	UpperIndex int
	DestAddr   string
}
```

We also add the following types for consistent hash ring structure in `ConsistentHashRing.go`

```go
type Node struct {
	Addr  string
	Index int
}

type ConsistentHashRing struct {
	RingSize int
	Nodes    []Node
}
```
## Consistent Hash Ring
`ConsistentHashRing.go` provides a skeleton implementation of the consistent hash ring structure. **You must implement the methods in this file which have `panic("todo")` as their body.**

## Server
`BlockStore.go` provides a skeleton implementation of the `BlockStoreInterface` and `MetaStore.go` provides a skeleton implementation of the `MetaStoreInterface` 
**You must implement the methods in these 2 files which have `panic("todo")` as their body.**

## User Client
The user client refers to the same client we implemented in Project 3 that does the sync operation. If you have implemented the client properly, assuming `GetBlockStoreMap()` could return multiple BlockStore addresses, then that client shall work automatically for this project. In case you haven’t made it to implement a correct client in Project 3, we have provided a user client binary to you in the starter code.

## Admin Client
Admin is a new client role we introduce to this project. An admin can manage the nodes in the cluster by adding/removing nodes dynamically. `SurfstoreRPCAdmin.go` provides the rpc admin stub for the AddNode and RemoveNode. **you don't need to modify this file**

## Setup
You will need to setup your runtime environment variables so that you can build your code and also use the executables that will be generated.
1. **You need to first download and install Docker (https://docs.docker.com/get-docker/).**
2. If you are using a Mac, open `~/.bash_profile` or if you are using a unix/linux machine, open `~/.bashrc`. Then add the following:
```
export GOPATH=<path to starter code>
export PATH=$PATH:$GOPATH/bin
export GO111MODULE=off
```
3. Run `source ~/.bash_profile` or `source ~/.bashrc`
## Usage
1. Only after you have implemented all the methods and completed the `Setup` steps, run the `build.sh` script provided with the starter code. This should create 4 executables in the `bin` folder inside your starter code directory.
```shell
> ./build.sh
> ls bin
SurfstoreClientExec SurfstoreServerExec SurfstoreAdminExec SurfstoreDebugExec
```

2. Run your server using the script provided in the starter code.
```shell
./run-server.sh -s <service> -p <port> -r <ring_size> -l -d (BlockStoreAddr*)
```
Here, `service` should be one of three values: meta, block, or both. This is used to specify the service provided by the server. `port` defines the port number that the server listens to (default=8080). `ring_size` defines the ring size we use for consistent hash ring (default=128). `-l` configures the server to only listen on localhost. `-d` configures the server to output log statements. Lastly, (BlockStoreAddr\*) is zero or more initial BlockStore addresses that the server is configured with. For module 3, the MetaStore should always start with 1 BlockStore address and if `service=both` then the BlockStoreAddr should be the `ip:port` of this server.

Examples:

```shell
Run the commands below on separate terminals (or nodes)
> ./run-server.sh -s block -p 8081 -l
> ./run-server.sh -s block -p 8082 -l
> ./run-server.sh -s meta -l localhost:8081 localhost:8082
```
The first line starts the first BlockStore server and listens only to localhost on port 8081. The ring size we use for consistent hash ring is 128 (by default).
The Second line starts the Second BlockStore server and listens only to localhost on port 8082. The ring size we use for consistent hash ring is 128 (by default).
The third line starts a MetaStore Server, listens only to localhost on port 8080, and references the BlockStore we created as the underlying BlockStore. (Note: if these are on separate nodes, then you should use the public ip address and remove `-l`). The ring size we use for consistent hash ring is 128 (by default).

3. From a new terminal (or a new node), run the user client using the script provided in the starter code (if using a new node, build using step 1 first). Use a base directory with some files in it.
```shell
> mkdir dataA
> cp ~/pic.jpg dataA/ 
> ./run-client.sh server_addr:port dataA 4096
```
This would sync pic.jpg to the server hosted on `server_addr:port`, using `dataA` as the base directory, with a block size of 4096 bytes.

4. From another terminal (or a new node), run the user client to sync with the server. (if using a new node, build using step 1 first)
```shell
> ls dataB/
> ./run-client.sh server_addr:port dataB 4096
> ls dataB/
pic.jpg index.txt
```
We observe that pic.jpg has been synced to this client.

5. Run the admin client to add or remove a BlockStore server.
```shell
./run-admin.sh -s <service> <MetaStoreAddr> <BlockStoreAddr>
```
Here, `service` should be one of two values: add or remove. This is used to specify the service provided by the admin. `MetaStoreAddr` is the address of the MetaStore server you have started. `BlockStoreAddr` should be the address of the BlockStore server you want to add or remove.

Examples:

```shell
Run the commands below on separate terminals (or nodes)
> ./run-server.sh -s block -p 8083 -l
> ./run-admin.sh -s add localhost:8080 localhost:8083
> ./run-admin.sh -s remove localhost:8080 localhost:8081
```

## Testing 
We have provided you with a script for printing the BlockMap of the BlockStore. You can use this for checking whether the blocks are stored in the right BlockStore.
```shell
./run-debug.sh -r <ring_size> <BlockStoreAddr>
```
Here, `ring_size` defines the ring size we use for consistent hash ring (default=128). `BlockStoreAddr` should be the address of the BlockStore server you want to print the BlockMap.

Examples:
```shell
> ./run-debug.sh localhost:8081
localhost:8081
--------BEGIN PRINT MAP--------
	 f5ad3def7576b054cdd88c3747437f4bfc07bc0352ed219accf3308d2ad8a2ac 44
---------END PRINT MAP--------
```

In this example, `f5ad3def7576b054cdd88c3747437f4bfc07bc0352ed219accf3308d2ad8a2ac` is the blockHash, and `44` is the ring index of this block.
