package networks

import (
	"blockchain-simple/block"
	"encoding/json"
	"fmt"
)

func QueryLatestMsg() []byte {
	return []byte(fmt.Sprintf("{\"type\": %d}", QueryLatest))
}

func ResponseLatestMsg() []byte {
	var v = &block.ResponseBlockchain{Type:ResponseBlockchain}
	d, _ := json.Marshal(Blockchain[len(Blockchain)-1:])
	v.Data = string(d)
	bs, _ := json.Marshal(v)
	return bs
}

func QueryAllMsg() []byte {
	return []byte(fmt.Sprintf("{\"type\": %d}", QueryAll))
}
