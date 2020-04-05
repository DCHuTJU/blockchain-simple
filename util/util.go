package util

import (
	"blockchain-simple/block"
	"crypto/sha256"
	"fmt"
	"log"
)

func ErrFatal(msg string, err error) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func CalculateHashForBlock(b *block.Block) string {
	sha := sha256.New()
	hash := sha.Sum([]byte(fmt.Sprintf("%d%s%d%s%s", b.Index, b.PreviousHash, b.Timestamp, b.Data, b.Label)))
	return string(hash)
}
