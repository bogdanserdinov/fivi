package client

import (
	"encoding/json"

	did_lib "fivi/lib/did"
)

func Decode(didBytes []byte) (*did_lib.DID, error) {
	did := new(did_lib.DID)
	err := json.Unmarshal(didBytes, did)
	return did, err
}
