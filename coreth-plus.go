package main

import (
	"context"
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/coreth"
	"github.com/ava-labs/coreth/core/types"
	"github.com/ava-labs/coreth/ethclient"
	"github.com/ava-labs/coreth/rpc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// FullEthClient meets the Ethereum AbiGen generated bindings for an Ethereum
// client to interact the token and swap contracts in the context of Deputy.
type ethClient interface {
	// ChainId retrieves the current chain ID for transaction replay protection.
	ChainID(ctx context.Context) (*big.Int, error)

	// BlockByHash returns the given full block.
	//
	// Note that loading full blocks requires two requests. Use HeaderByHash
	// if you don't need all transactions or uncle headers.
	BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error)

	// BlockByNumber returns a block from the current canonical chain. If number is nil, the
	// latest known block is returned.
	//
	// Note that loading full blocks requires two requests. Use HeaderByNumber
	// if you don't need all transactions or uncle headers.
	BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error)

	// BlockNumber returns the most recent block number
	BlockNumber(ctx context.Context) (uint64, error)

	// getBlock(ctx context.Context, method string, args ...interface{}) (*types.Block, error)

	// HeaderByHash returns the block header with the given hash.
	HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error)

	// HeaderByNumber returns a block header from the current canonical chain. If number is
	// nil, the latest known header is returned.
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)

	// TransactionByHash returns the transaction with the given hash.
	TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error)

	// TransactionSender returns the sender address of the given transaction. The transaction
	// must be known to the remote node and included in the blockchain at the given block and
	// index. The sender is the one derived by the protocol at the time of inclusion.
	//
	// There is a fast-path for transactions retrieved by TransactionByHash and
	// TransactionInBlock. Getting their sender address can be done without an RPC interaction.
	TransactionSender(ctx context.Context, tx *types.Transaction, block common.Hash, index uint) (common.Address, error)

	// TransactionCount returns the total number of transactions in the given block.
	TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error)

	// TransactionInBlock returns a single transaction at index in the given block.
	TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (*types.Transaction, error)

	// TransactionReceipt returns the receipt of a transaction by transaction hash.
	// Note that the receipt is not available for pending transactions.
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)

	// SubscribeNewHead subscribes to notifications about the current blockchain head
	// on the given channel.
	SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (coreth.Subscription, error)

	// State Access

	// NetworkID returns the network ID (also known as the chain ID) for this chain.
	NetworkID(ctx context.Context) (*big.Int, error)

	// BalanceAt returns the wei balance of the given account.
	// The block number can be nil, in which case the balance is taken from the latest known block.
	BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error)

	// AssetBalanceAt returns the [assetID] balance of the given account
	// The block number can be nil, in which case the balance is taken from the latest known block.
	AssetBalanceAt(ctx context.Context, account common.Address, assetID ids.ID, blockNumber *big.Int) (*big.Int, error)

	// StorageAt returns the value of key in the contract storage of the given account.
	// The block number can be nil, in which case the value is taken from the latest known block.
	StorageAt(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int) ([]byte, error)

	// CodeAt returns the contract code of the given account.
	// The block number can be nil, in which case the code is taken from the latest known block.
	CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error)

	// NonceAt returns the account nonce of the given account.
	// The block number can be nil, in which case the nonce is taken from the latest known block.
	NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error)

	// Filters

	// FilterLogs executes a filter query.
	FilterLogs(ctx context.Context, q coreth.FilterQuery) ([]types.Log, error)

	// SubscribeFilterLogs subscribes to the results of a streaming filter query.
	SubscribeFilterLogs(ctx context.Context, q coreth.FilterQuery, ch chan<- types.Log) (coreth.Subscription, error)

	// Pending State is irrelevant in Coreth

	// PendingBalanceAt returns the wei balance of the given account in the pending state.
	PendingBalanceAt(ctx context.Context, account common.Address) (*big.Int, error)

	// PendingStorageAt returns the value of key in the contract storage of the given account in the pending state.
	PendingStorageAt(ctx context.Context, account common.Address, key common.Hash) ([]byte, error)

	// PendingCodeAt returns the contract code of the given account in the pending state.
	PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error)

	// PendingNonceAt returns the account nonce of the given account in the pending state.
	// This is the nonce that should be used for the next transaction.
	PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)

	// PendingTransactionCount returns the total number of transactions in the pending state.
	PendingTransactionCount(ctx context.Context) (uint, error)

	// Contract Calling

	// CallContract executes a message call transaction, which is directly executed in the VM
	// of the node, but never mined into the blockchain.
	//
	// blockNumber selects the block height at which the call runs. It can be nil, in which
	// case the code is taken from the latest known block. Note that state from very old
	// blocks might not be available.
	CallContract(ctx context.Context, msg coreth.CallMsg, blockNumber *big.Int) ([]byte, error)

	// PendingCallContract executes a message call transaction using the EVM.
	// The state seen by the contract call is the pending state.
	PendingCallContract(ctx context.Context, msg coreth.CallMsg) ([]byte, error)

	// SuggestGasPrice retrieves the currently suggested gas price to allow a timely
	// execution of a transaction.
	SuggestGasPrice(ctx context.Context) (*big.Int, error)

	// EstimateGas tries to estimate the gas needed to execute a specific transaction based on
	// the current pending state of the backend blockchain. There is no guarantee that this is
	// the true gas limit requirement as other transactions may be added or removed by miners,
	// but it should provide a basis for setting a reasonable default.
	EstimateGas(ctx context.Context, msg coreth.CallMsg) (uint64, error)

	// SendTransaction injects a signed transaction into the pending pool for execution.
	//
	// If the transaction was a contract creation use the TransactionReceipt method to get the
	// contract address after the transaction has been mined.
	SendTransaction(ctx context.Context, tx *types.Transaction) error

	Close()
}

// Compile time enforcement of minimum viable binding ethere
var _ ethClient = &avaClient{}

// Dial connects a client to the given URL.
func Dial(url string) (*avaClient, error) {
	c, err := rpc.DialContext(context.Background(), url)
	if err != nil {
		return nil, err
	}

	fc := &avaClient{
		Client: ethclient.NewClient(c),
		c:      c,
	}

	return fc, err
}

func (pc *avaClient) Close() {
	pc.c.Close()
}

type avaClient struct {
	*ethclient.Client
	c *rpc.Client
}

// PendingBalanceAt returns the wei balance of the given account in the pending state.
func (pc *avaClient) PendingBalanceAt(ctx context.Context, account common.Address) (*big.Int, error) {
	var result hexutil.Big
	err := pc.c.CallContext(ctx, &result, "eth_getBalance", account, "pending")
	return (*big.Int)(&result), err
}

// PendingStorageAt returns the value of key in the contract storage of the given account in the pending state.
func (pc *avaClient) PendingStorageAt(ctx context.Context, account common.Address, key common.Hash) ([]byte, error) {
	var result hexutil.Bytes
	err := pc.c.CallContext(ctx, &result, "eth_getStorageAt", account, key, "pending")
	return result, err
}

// PendingCodeAt returns the contract code of the given account in the pending state.
func (pc *avaClient) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	var result hexutil.Bytes
	err := pc.c.CallContext(ctx, &result, "eth_getCode", account, "pending")
	return result, err
}

// PendingNonceAt returns the account nonce of the given account in the pending state.
// This is the nonce that should be used for the next transaction.
func (pc *avaClient) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	var result hexutil.Uint64
	err := pc.c.CallContext(ctx, &result, "eth_getTransactionCount", account, "pending")
	return uint64(result), err
}

// PendingTransactionCount returns the total number of transactions in the pending state.
func (pc *avaClient) PendingTransactionCount(ctx context.Context) (uint, error) {
	var num hexutil.Uint
	err := pc.c.CallContext(ctx, &num, "eth_getBlockTransactionCountByNumber", "pending")
	return uint(num), err
}

// PendingCallContract executes a message call transaction using the EVM.
// The state seen by the contract call is the pending state.
func (pc *avaClient) PendingCallContract(ctx context.Context, msg coreth.CallMsg) ([]byte, error) {
	var hex hexutil.Bytes
	err := pc.c.CallContext(ctx, &hex, "eth_call", toCallArg(msg), "pending")
	if err != nil {
		return nil, err
	}
	return hex, nil
}

func toCallArg(msg coreth.CallMsg) interface{} {
	arg := map[string]interface{}{
		"from": msg.From,
		"to":   msg.To,
	}
	if len(msg.Data) > 0 {
		arg["data"] = hexutil.Bytes(msg.Data)
	}
	if msg.Value != nil {
		arg["value"] = (*hexutil.Big)(msg.Value)
	}
	if msg.Gas != 0 {
		arg["gas"] = hexutil.Uint64(msg.Gas)
	}
	if msg.GasPrice != nil {
		arg["gasPrice"] = (*hexutil.Big)(msg.GasPrice)
	}
	return arg
}
