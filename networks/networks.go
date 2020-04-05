package networks

import (
	"blockchain-simple/block"
	"blockchain-simple/util"
	"encoding/json"
	"golang.org/x/net/websocket"
	"io"
	"log"
)

func ConnectToPeers(peersAddr []string) {
	for _, peer := range peersAddr {
		if peer == "" {
			continue
		}
		ws, err := websocket.Dial(peer, "", peer)
		if err != nil {
			log.Println("dial to peer", err)
			continue
		}
		InitConnection(ws)
	}
}

func InitConnection(ws *websocket.Conn) {
	go WsHandleP2P(ws)

	log.Println("query lastest block.")
	ws.Write(QueryLatestMsg())
}

func WsHandleP2P(ws *websocket.Conn) {
	var (
		v    = &block.ResponseBlockchain{}
		peer = ws.LocalAddr().String()
	)
	Sockets = append(Sockets, ws)
	for {
		var msg []byte
		err := websocket.Message.Receive(ws, &msg)
		if err == io.EOF {
			log.Printf("p2p Peer[%s] shutdown, remove it form peers pool.\n", peer)
			break
		}
		if err != nil {
			log.Println("Can't receive p2p msg from ", peer, err.Error())
			break
		}

		log.Printf("Received[from %s]: %s.\n", peer, msg)
		err = json.Unmarshal(msg, v)
		util.ErrFatal("invalid p2p msg", err)

		switch v.Type {
		case QueryLatest:
			v.Type = ResponseBlockchain

			bs := ResponseLatestMsg()
			log.Printf("responseLatestMsg: %s\n", bs)
			ws.Write(bs)

		case QueryAll:
			d, _ := json.Marshal(Blockchain)
			v.Type = ResponseBlockchain
			v.Data = string(d)
			bs, _ := json.Marshal(v)
			log.Printf("responseChainMsg: %s\n", bs)
			ws.Write(bs)

		case ResponseBlockchain:
			HandleBlockchainResponse([]byte(v.Data))
		}

	}
}

// 广播 📢
func Broadcast(msg []byte) {
	for n, socket := range Sockets {
		_, err := socket.Write(msg)
		if err != nil {
			log.Printf("peer [%s] disconnected.", socket.RemoteAddr().String())
			Sockets = append(Sockets[0:n], Sockets[n+1:]...)
		}
	}
}



