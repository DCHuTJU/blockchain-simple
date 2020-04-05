package networks

import (
	"blockchain-simple/block"
	"flag"
	"golang.org/x/net/websocket"
)

const (
	QueryLatest = iota
	QueryAll
	ResponseBlockchain
)

var (
	Sockets  []*websocket.Conn
	Blockchain = []*block.Block{block.GenerateGenesisBlock("first block", "1")}
	HttpAddr = flag.String("api", ":3001", "api server address.")
	P2pAddr = flag.String("p2p", ":6001", "p2p server address.")
	InitialPeers = flag.String("peers", "http://localhost:6001", "initial peers.")
)


