package validation

import (
	"blockchain-simple/block"
	"blockchain-simple/util"
	"log"
)

func IsValidChain(bc []*block.Block) bool {
	genesisblock := block.GenerateGenesisBlock("first block", "1")
	if bc[0].String() != genesisblock.String() {
		log.Println("No same GenesisBlock.", bc[0].String())
		return false
	}
	var temp = []*block.Block{bc[0]}
	for i := 1; i < len(bc); i++ {
		if IsValidNewBlock(bc[i], temp[i-1]) {
			temp = append(temp, bc[i])
		} else {
			return false
		}
	}
	return true
}

func IsValidNewBlock(nb, pb *block.Block) (ok bool) {
	if nb.Hash == util.CalculateHashForBlock(nb) &&
		pb.Index+1 == nb.Index &&
		pb.Hash == nb.PreviousHash {
		ok = true
	}
	return
}