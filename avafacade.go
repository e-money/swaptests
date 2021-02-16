package main

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ava-labs/coreth"
	"github.com/btcsuite/btcutil/bech32"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/e-money/swaptests/avaabi"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

var HTLTEventHash = common.HexToHash("0xb3e26d98380491276a8dce9d38fd1049e89070230ff5f36ebb55ead64500ade1")

type ethFacade struct {
	Abi               abi.ABI
	Provider          string
	Client            ethClient
	SwapContractAddr  common.Address
	TokenContractAddr common.Address
}

func newEthFacade(provider string, contractAddr, tokenContractAddr common.Address) *ethFacade {
	contractAbi, err := abi.JSON(strings.NewReader(avaabi.ERC20AtomicSwapperABI))
	if err != nil {
		panic("marshal abi error")
	}

	var client ethClient

	client, err = Dial(provider)
	if err != nil {
		panic("new eth client error")
	}

	return &ethFacade{
		Provider: provider,

		Abi:               contractAbi,
		Client:            client,
		SwapContractAddr:  contractAddr,
		TokenContractAddr: tokenContractAddr,
	}
}

func (e *ethFacade) crEthTrx(timestamp int64, randomHash common.Hash) (string, error) {
	trxHash, cmnErr := e.HTLT(
		randomHash,
		timestamp,
		timeSpan,

		ethRecipientAddr, // runtime key provided
		emSenderAddr,     // constant address no key is needed
		emUserAddr,       //   ~~~

		big.NewInt(1))
	if cmnErr != nil {
		fmt.Printf("init eth swap error, err=%s", cmnErr)
		return "", cmnErr
	}

	fmt.Printf("init eth tx sent: %s", trxHash)

	return trxHash, nil
}

func (e *ethFacade) hasSwap(swapId common.Hash) (bool, error) {
	instance, err := avaabi.NewERC20AtomicSwapper(e.SwapContractAddr, e.Client)
	if err != nil {
		return false, err
	}

	return instance.IsSwapExist(nil, swapId)
}

func (e *ethFacade) getSwap(swapId common.Hash) (*SwapRequest, error) {
	htltEvent, err := e.getHTLTEvent(swapId)
	if err != nil {
		return nil, err
	}

	return &SwapRequest{
		Id:                  swapId,
		RandomNumberHash:    htltEvent.RandomNumberHash,
		ExpireTimestamp:     htltEvent.ExpireSeconds.Int64(),
		SenderAddress:       htltEvent.MsgSender.String(),
		RecipientAddress:    htltEvent.RecipientAddr.String(),
		RecipientOtherChain: htltEvent.Bep2Addr.String(),
		OutAmount:           htltEvent.OutAmount,
	}, nil
}

func (e *ethFacade) getHTLTEvent(swapId common.Hash) (*HTLTEvent, error) {
	topics := [][]common.Hash{{HTLTEventHash}, {}, {}, {swapId}}
	logs, err := e.Client.FilterLogs(context.Background(), coreth.FilterQuery{
		Topics: topics,
	})
	if err != nil {
		return nil, err
	}

	if len(logs) == 0 {
		return nil, fmt.Errorf("swap id does not exist, swap_id=%s", swapId.String())
	}

	htltEvent, err := parseHTLTEvent(&e.Abi, &logs[0])
	if err != nil {
		return nil, fmt.Errorf("parse event log error, er=%s", err.Error())
	}

	return htltEvent, nil
}

// DecodeAndConvert decodes a bech32 encoded string and converts to base64 encoded bytes
func bech32Bytes(bech string) ([]byte, error) {
	_, data, err := bech32.Decode(bech)
	if err != nil {
		return nil, errors.Wrap(err, "decoding bech32 failed")
	}
	converted, err := bech32.ConvertBits(data, 5, 8, false)
	if err != nil {
		return nil, errors.Wrap(err, "decoding bech32 failed")
	}
	return converted, nil
}

func (e *ethFacade) HTLT(randomNumberHash common.Hash, timestamp, timeSpan int64, recipientAddr,
	otherChainSenderAddr, otherChainRecipientAddr string, outAmount *big.Int) (string, error) {
	fmt.Println("Sending HTLT with timestamp: ", time.Unix(timestamp, 0).String(),
		"expiration @", time.Unix(timestamp+timeSpan, 0).String())

	auth, err := e.getTransactor()
	if err != nil {
		return "", err
	}

	instance, err := avaabi.NewERC20AtomicSwapper(e.SwapContractAddr, e.Client)
	if err != nil {
		return "", err
	}

	recvAddr := common.HexToAddress(recipientAddr)
	bep2RecipientAddr, err := bech32Bytes(otherChainRecipientAddr)
	assertNoError(err, fmt.Sprintf("bech32Bytes(otherChainRecipientAddr) %s", otherChainRecipientAddr))

	bep2SenderAddr, err := bech32Bytes(otherChainSenderAddr)
	assertNoError(err, fmt.Sprintf("bech32Bytes(otherChainSenderAddr) %s", otherChainSenderAddr))

	//instanceERC20, err := avaabi.NewIERC20(ethfacade.TokenContractAddr, ethfacade.Client)
	//if err != nil {
	//	return "", dc.NewError(err, false)
	//}
	//sendTrx, err := instanceERC20.Transfer(auth, recvAddr, big.NewInt(0))
	//if err != nil {
	//	return "", dc.NewError(err, false)
	//}
	//util.Logger.Debugf("sendTrx: %v", sendTrx)
	//
	//time.Sleep(2 * time.Second)

	//auth, err = ethfacade.GetTransactor()
	//if err != nil {
	//	return "", dc.NewError(err, false)
	//}

	tx, err := instance.Htlt(auth, randomNumberHash, uint64(timestamp), big.NewInt(timeSpan), recvAddr,
		common.BytesToAddress(bep2SenderAddr), common.BytesToAddress(bep2RecipientAddr), outAmount, big.NewInt(0))
	if err != nil {
		return "", err
	}
	fmt.Printf("init tx sent: %s\n", tx.Hash().Hex())
	return tx.Hash().String(), nil
}

func (e *ethFacade) getTransactor() (*avaabi.TransactOpts, error) {
	pendingNonce, err := e.Client.PendingNonceAt(context.Background(), common.HexToAddress(ethSenderAddr))
	if err != nil {
		return nil, err
	}

	fmt.Println("Pending nonce:", pendingNonce)

	auth := avaabi.NewKeyedTransactor(ethPrvKey)
	auth.Nonce = big.NewInt(int64(pendingNonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = gasLimit
	auth.GasPrice = big.NewInt(gasPrice)

	return auth, nil
}

func calculateSwapID(randomNumberHash []byte, sender []byte, senderOtherChain []byte) []byte {
	data := randomNumberHash
	data = append(data, []byte(sender)...)
	data = append(data, []byte(senderOtherChain)...)
	return shaHash(data)
}

func (e *ethFacade) calcSwapId(randomNumberHash common.Hash, sender string, senderOtherChain string) ([]byte, error) {
	bep2Addr := sdk.AccAddress{}
	if senderOtherChain != "" {
		parsedAddr, err := sdk.AccAddressFromBech32(senderOtherChain)
		if err != nil {
			return nil, err
		}

		bep2Addr = parsedAddr
	}

	return calculateSwapID(randomNumberHash[:], common.FromHex(sender), bep2Addr[:]), nil
}
