package gpp

import (
	"baby-chain/blockchain"
	"baby-chain/blockchain/block"
	cs "baby-chain/blockchain/consensus_state"
	"baby-chain/tools/data"
	"encoding/hex"
	"fmt"
)

var CSGenesis = cs.ConsensusState{
	Check: func(_ *blockchain.Blockchain, _ *cs.StateData, b block.Block) bool {
		if b.Header[block.Head] != "Genesis" {
			return false
		}
		if _, ok := b.Data["master"].(string); !ok {
			return false
		}
		return true
	},
	Run: func(bc *blockchain.Blockchain, sd *cs.StateData, b block.Block) error {
		sd.Data["master"] = b.Data["master"]
		return nil
	},
}

var CSRegisterIns = cs.ConsensusState{
	Check: func(_ *blockchain.Blockchain, _ *cs.StateData, b block.Block) bool {
		if b.Header[block.Head] != "RegisterIns" {
			return false
		}
		if _, ok := b.Header["signature"].(string); !ok {
			return false
		}
		if _, ok := b.Data["public_key"].(string); !ok {
			return false
		}
		return true
	},
	Run: func(bc *blockchain.Blockchain, sd *cs.StateData, b block.Block) error {
		if err := cs.SignCheckBlock(b, "signature"); err != nil {
			return err
		}
		master := sd.Data["master"].(data.Data)
		if master[b.Data["public_key"].(string)] == nil {
			return fmt.Errorf("public key not registered")
		}
		if sd.Data["ins"] == nil {
			sd.Data["ins"] = data.Data{}
		}
		ins := sd.Data["ins"].(data.Data)
		ins[hex.EncodeToString(b.Hash[:])] = b.Data
		return nil
	},
}

var CSBuyIns = cs.ConsensusState{
	Check: func(_ *blockchain.Blockchain, _ *cs.StateData, b block.Block) bool {
		if b.Header[block.Head] != "BuyIns" {
			return false
		}
		if _, ok := b.Header["signature"].(string); !ok {
			return false
		}
		if _, ok := b.Data["public_key"].(string); !ok {
			return false
		}
		if _, ok := b.Data["policy_ref_id"].(string); !ok {
			return false
		}
		return true
	},
	Run: func(bc *blockchain.Blockchain, sd *cs.StateData, b block.Block) error {
		if cs.SignCheckBlock(b, "signature") != nil {
			return fmt.Errorf("signature check failed")
		}
		nodes, ok := sd.Data[cs.NODES].(data.Data)
		if !ok {
			return fmt.Errorf("no nodes")
		}
		node, ok := nodes[b.Data["public_key"].(string)].(data.Data)
		if !ok {
			return fmt.Errorf("public key not registered")
		}
		ins, ok := sd.Data["ins"].(data.Data)
		if !ok || ins[b.Data["policy_ref_id"].(string)] == nil {
			return fmt.Errorf("policy not registered")
		}
		pols, ok := node["Policies"].(data.Array)
		if !ok {
			node["Policies"] = data.Array{}
			pols = node["Policies"].(data.Array)
		}
		for _, pol := range pols {
			if pol == b.Data["policy_ref_id"].(string) {
				return fmt.Errorf("policy already bought")
			}
		}
		node["Policies"] = append(pols, b.Data["policy_ref_id"].(string))
		return nil
	},
}
