package main

import (
	"encoding/hex"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	rgbpp "github.com/shaojunda/rgbpp-sdk-demo/molecule"
)

func main() {

}

func GenBtcTimeLockArgs(toLock *types.Script, btcTxId string, after int64) (string, error) {
	b := rgbpp.NewBTCTimeLockBuilder()
	b.LockScript(*rgbpp.PackCkbScript(toLock))

	btcTxId, err := ReverseHash(btcTxId)
	if err != nil {
		return "", err
	}

	toLock.Serialize()

	btcTxIdBytes, err := hex.DecodeString(btcTxId)
	if err != nil {
		return "", err
	}

	b.BtcTxid(*rgbpp.Byte32FromSliceUnchecked(btcTxIdBytes))

	a, err := rgbpp.PackAfter(after)
	if err != nil {
		return "", err
	}

	b.After(a)
	btcTimeLock := b.Build()

	return hex.EncodeToString(btcTimeLock.AsSlice()), nil
}

func ReverseHash(txId string) (string, error) {
	if len(txId) == 0 {
		return "", fmt.Errorf("invalid hash")
	}

	hashBytes, err := hex.DecodeString(txId)
	if err != nil {
		return "", err
	}

	for i, j := 0, len(hashBytes)-1; i < j; i, j = i+1, j-1 {
		hashBytes[i], hashBytes[j] = hashBytes[j], hashBytes[i]
	}

	return hex.EncodeToString(hashBytes), nil
}
