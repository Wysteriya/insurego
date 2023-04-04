package tools

import (
	"baby-chain/errors"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

type Hash [32]byte

func (h *Hash) MarshalJSON() ([]byte, error) {
	return json.Marshal(h.Hex())
}

func (h *Hash) UnmarshalJSON(save []byte) error {
	var _h string
	err := json.Unmarshal(save, &_h)
	if err != nil {
		return err
	}
	h_, err := hex.DecodeString(_h)
	if err != nil {
		return err
	}
	if len(h_) != 32 {
		return errors.InvalidSave("hash is 32 bytes")
	}
	copy(h[:], h_)
	return nil
}

func (h *Hash) Hex() string {
	return hex.EncodeToString(h[:])
}

func HashB(chunks ...[]byte) Hash {
	return sha256.Sum256(bytes.Join(chunks, nil))
}
