package main

import (
	"blockchain-simple/api"
	"blockchain-simple/networks"
	"blockchain-simple/util"
	"flag"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"strings"
)

func main() {
 	flag.Parse()
 	// 构建多个 peers
 	networks.ConnectToPeers(strings.Split(*networks.InitialPeers, ","))

 	http.HandleFunc("/block", api.HandleBlocks)
 	http.HandleFunc("/mine_block", api.HandleMineBlock)
 	http.HandleFunc("/peers", api.HandlePeers)
 	http.HandleFunc("/add_peer", api.HandleAddPeer)

 	// 同时处理
 	go func() {
 		log.Println("Listen HTTP on", *networks.HttpAddr)
 		util.ErrFatal("start api server", http.ListenAndServe(*networks.HttpAddr, nil))
	}()

 	http.Handle("/", websocket.Handler(networks.WsHandleP2P))
 	log.Println("Listen P2P on", *networks.P2pAddr)
 	util.ErrFatal("start p2p server", http.ListenAndServe(*networks.P2pAddr, nil))
}
