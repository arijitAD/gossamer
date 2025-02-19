// Copyright 2019 ChainSafe Systems (ON) Corp.
// This file is part of gossamer.
//
// The gossamer library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The gossamer library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the gossamer library. If not, see <http://www.gnu.org/licenses/>.

package genesis

import (
	"github.com/ChainSafe/gossamer/lib/common"
)

// Genesis stores the data parsed from the genesis configuration file
type Genesis struct {
	Name            string                 `json:"name"`
	ID              string                 `json:"id"`
	ChainType       string                 `json:"chainType"`
	Bootnodes       []string               `json:"bootNodes"`
	ProtocolID      string                 `json:"protocolId"`
	Genesis         Fields                 `json:"genesis"`
	Properties      map[string]interface{} `json:"properties"`
	ForkBlocks      []string               `json:"forkBlocks"`
	BadBlocks       []string               `json:"badBlocks"`
	ConsensusEngine string                 `json:"consensusEngine"`
}

// Data defines the genesis file data formatted for trie storage
type Data struct {
	Name            string
	ID              string
	ChainType       string
	Bootnodes       [][]byte
	ProtocolID      string
	Properties      map[string]interface{}
	ForkBlocks      []string
	BadBlocks       []string
	ConsensusEngine string
}

// Fields stores genesis raw data, and human readable runtime data
type Fields struct {
	Raw     map[string]map[string]string      `json:"raw,omitempty"`
	Runtime map[string]map[string]interface{} `json:"runtime,omitempty"`
}

// GenesisData formats genesis for trie storage
func (g *Genesis) GenesisData() *Data {
	return &Data{
		Name:            g.Name,
		ID:              g.ID,
		ChainType:       g.ChainType,
		Bootnodes:       common.StringArrayToBytes(g.Bootnodes),
		ProtocolID:      g.ProtocolID,
		Properties:      g.Properties,
		ForkBlocks:      g.ForkBlocks,
		BadBlocks:       g.BadBlocks,
		ConsensusEngine: g.ConsensusEngine,
	}
}

// GenesisFields returns the genesis fields including genesis raw data
func (g *Genesis) GenesisFields() Fields {
	return g.Genesis
}
