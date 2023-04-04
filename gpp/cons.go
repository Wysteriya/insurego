package gpp

import (
	"baby-chain/blockchain"
	"baby-chain/blockchain/block"
	cs "baby-chain/blockchain/consensus_state"
)

var CSRegisterIns = cs.ConsensusState{
	Check: func(_ *blockchain.Blockchain, _ *cs.StateData, b block.Block) bool {
		if b.Header[block.Head] != "RegisterIns" {
			return false
		}
		return true
	},
	Run: func(bc *blockchain.Blockchain, sd *cs.StateData, b block.Block) error {
		return nil
	},
}
