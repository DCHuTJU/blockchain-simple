package api

import (
	"blockchain-simple/networks"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func HandleBlocks(w http.ResponseWriter, r *http.Request) {
	bs, _ := json.Marshal(networks.Blockchain)
	w.Write(bs)
}

func HandleMineBlock(w http.ResponseWriter, r *http.Request) {
	var v struct {
		Data  string `json:"data"`
		Label string `json:"label"`
	}
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err := decoder.Decode(&v)
	if err != nil {
		w.WriteHeader(http.StatusGone)
		log.Println("[API] invalid block data : ", err.Error())
		w.Write([]byte("invalid block data. " + err.Error() + "\n"))
		return
	}
	block := networks.GenerateNextBlock(v.Data, v.Label)
	networks.AddBlock(block)
	networks.Broadcast(networks.ResponseLatestMsg())
}

func HandlePeers(w http.ResponseWriter, r *http.Request) {
	var slice []string
	for _, socket := range networks.Sockets {
		// 确认连接
		if socket.IsClientConn() {
			slice = append(slice, strings.Replace(socket.LocalAddr().String(), "ws://", "", 1))
		} else {
			slice = append(slice, socket.Request().RemoteAddr)
		}
	}
	bs, _ := json.Marshal(slice)
	w.Write(bs)
}

func HandleAddPeer(w http.ResponseWriter, r *http.Request) {
	var v struct {
		Peer string `json:"peer"`
	}
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err := decoder.Decode(&v)
	if err != nil {
		w.WriteHeader(http.StatusGone)
		log.Println("[API] invalid peer data : ", err.Error())
		w.Write([]byte("invalid peer data. " + err.Error()))
		return
	}
	networks.ConnectToPeers([]string{v.Peer})
}
