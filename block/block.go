package block

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Index        int64  `json:"index"`
	PreviousHash string `json:"previousHash"`
	Timestamp    int64  `json:"timestamp"`
	Data         string `json:"data"`
	Label        string `json:"label"`
	Hash         string `json:"hash"`
}

func (b *Block) String() string {
	return fmt.Sprintf("index: %d, prevousHash: %s, timestamp: %d, data: %d, label: %d, hash: %s", b.Index, b.PreviousHash, b.Timestamp,
		b.Data, b.Label, b.Hash)
}

type ByIndex []*Block

func (b ByIndex) Len() int           { return len(b) }
func (b ByIndex) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByIndex) Less(i, j int) bool { return b[i].Index < b[j].Index }

type ResponseBlockchain struct {
	Type int     `json:"type"`
	Data string  `json:"data"`
}

// 生成 genesisblock
func GenerateGenesisBlock(data, label string) *Block {
	var block = &Block {
		Index: 0,
		PreviousHash: "0",
		Timestamp: time.Now().UnixNano(),
		Data: data,
		Label: label,
		Hash: "",
	}
	// int64 转换到 string
	time := strconv.FormatInt(block.Timestamp, 10)
	s := fmt.Sprintf("%d", block.Index) + block.PreviousHash + time + block.Data + block.Label
	sByte := []byte(s)
	hash := sha256.New()
	hash.Write(sByte)
	block.Hash = hex.EncodeToString(hash.Sum(nil))
	return block
}