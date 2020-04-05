package networks

import (
	"blockchain-simple/block"
	"blockchain-simple/util"
	"blockchain-simple/validation"
	"encoding/json"
	"log"
	"sort"
	"time"
)

func GetLatestBlock() *block.Block {
	return Blockchain[len(Blockchain)-1]
}

func HandleBlockchainResponse(msg []byte) {
	var receivedBlocks = []*block.Block{}

	err := json.Unmarshal(msg, &receivedBlocks)
	util.ErrFatal("invalid blockchain", err)

	sort.Sort(block.ByIndex(receivedBlocks))

	latestBlockReceived := receivedBlocks[len(receivedBlocks)-1]
	latestBlockHeld := GetLatestBlock()
	if latestBlockReceived.Index > latestBlockHeld.Index {
		log.Printf("blockchain possibly behind. We got: %d Peer got: %d", latestBlockHeld.Index, latestBlockReceived.Index)
		if latestBlockHeld.Hash == latestBlockReceived.PreviousHash {
			log.Println("We can append the received block to our chain.")
			Blockchain = append(Blockchain, latestBlockReceived)
		} else if len(receivedBlocks) == 1 {
			log.Println("We have to query the chain from our peer.")
			Broadcast(QueryAllMsg())
		} else {
			log.Println("Received blockchain is longer than current blockchain.")
			ReplaceChain(receivedBlocks)
		}
	} else {
		log.Println("received blockchain is not longer than current blockchain. Do nothing.")
	}
}

// 更新update blockchain
func ReplaceChain(bc []*block.Block) {
	if validation.IsValidChain(bc) && len(bc) > len(Blockchain) {
		log.Println("Received blockchain is valid. Replacing current blockchain with received blockchain.")
		Blockchain = bc
	} else {
		log.Println("Received blockchain invalid.")
	}
}

// 生成新的区块
func GenerateNextBlock(data string, label string) *block.Block {
	var previousBlock = GetLatestBlock()
	nb := &block.Block{
		Data:         data,
		PreviousHash: previousBlock.Hash,
		Index:        previousBlock.Index + 1,
		Timestamp:    time.Now().Unix(),
		Label:        label,
	}
	nb.Hash = util.CalculateHashForBlock(nb)
	return nb
}

func AddBlock(b *block.Block) {
	if validation.IsValidNewBlock(b, GetLatestBlock()) {
		Blockchain = append(Blockchain, b)
	}
}
