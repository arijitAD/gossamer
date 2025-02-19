package modules

import (
	"math/big"

	"github.com/ChainSafe/gossamer/dot/state"
	"github.com/ChainSafe/gossamer/dot/types"
	"github.com/ChainSafe/gossamer/lib/common"
	"github.com/ChainSafe/gossamer/lib/crypto"
	"github.com/ChainSafe/gossamer/lib/runtime"
	"github.com/ChainSafe/gossamer/lib/transaction"
)

// StorageAPI is the interface for the storage state
type StorageAPI interface {
	GetStorage(root *common.Hash, key []byte) ([]byte, error)
	GetStorageByBlockHash(bhash common.Hash, key []byte) ([]byte, error)
	Entries(root *common.Hash) (map[string][]byte, error)
	RegisterStorageChangeChannel(sub state.StorageSubscription) (byte, error)
	UnregisterStorageChangeChannel(id byte)
	GetStateRootFromBlock(bhash *common.Hash) (*common.Hash, error)
	GetKeysWithPrefix(root *common.Hash, prefix []byte) ([][]byte, error)
}

// BlockAPI is the interface for the block state
type BlockAPI interface {
	GetHeader(hash common.Hash) (*types.Header, error)
	BestBlockHash() common.Hash
	GetBlockByHash(hash common.Hash) (*types.Block, error)
	GetBlockHash(blockNumber *big.Int) (*common.Hash, error)
	GetFinalizedHash(uint64, uint64) (common.Hash, error)
	HasJustification(hash common.Hash) (bool, error)
	GetJustification(hash common.Hash) ([]byte, error)
	RegisterImportedChannel(ch chan<- *types.Block) (byte, error)
	UnregisterImportedChannel(id byte)
	RegisterFinalizedChannel(ch chan<- *types.Header) (byte, error)
	UnregisterFinalizedChannel(id byte)
	SubChain(start, end common.Hash) ([]common.Hash, error)
}

// NetworkAPI interface for network state methods
type NetworkAPI interface {
	Health() common.Health
	NetworkState() common.NetworkState
	Peers() []common.PeerInfo
	NodeRoles() byte
	Stop() error
	Start() error
	IsStopped() bool
}

// BlockProducerAPI is the interface for BlockProducer methods
type BlockProducerAPI interface {
	Pause() error
	Resume() error
	EpochLength() uint64
	SlotDuration() uint64
}

// TransactionStateAPI ...
type TransactionStateAPI interface {
	AddToPool(*transaction.ValidTransaction) common.Hash
	Pop() *transaction.ValidTransaction
	Peek() *transaction.ValidTransaction
	Pending() []*transaction.ValidTransaction
}

// CoreAPI is the interface for the core methods
type CoreAPI interface {
	InsertKey(kp crypto.Keypair)
	HasKey(pubKeyStr string, keyType string) (bool, error)
	GetRuntimeVersion(bhash *common.Hash) (runtime.Version, error)
	IsBlockProducer() bool
	HandleSubmittedExtrinsic(types.Extrinsic) error
	GetMetadata(bhash *common.Hash) ([]byte, error)
}

// RPCAPI is the interface for methods related to RPC service
type RPCAPI interface {
	Methods() []string
	BuildMethodNames(rcvr interface{}, name string)
}

// RuntimeAPI is the interface for runtime methods
type RuntimeAPI interface {
	ValidateTransaction(e types.Extrinsic) (*transaction.Validity, error)
}

// SystemAPI is the interface for handling system methods
type SystemAPI interface {
	SystemName() string
	SystemVersion() string
	NodeName() string
	Properties() map[string]interface{}
	ChainType() string
}
