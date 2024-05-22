package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"testing"
)

func TestGenBtcTimeLockArgs(t *testing.T) {
	type args struct {
		toLock  *types.Script
		btcTxId string
		after   int64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "valid input",
			args: args{
				toLock: &types.Script{
					CodeHash: types.HexToHash("0x0101010101010101010101010101010101010101010101010101010101010101"),
					HashType: types.HashTypeType,
					Args:     common.FromHex("0x0202020202020202020202020202020202020202"),
				},
				btcTxId: "0303030303030303030303030303030303030303030303030303030303030303",
				after:   42,
			},
			want:    "7d00000010000000590000005d000000490000001000000030000000310000000101010101010101010101010101010101010101010101010101010101010101011400000002020202020202020202020202020202020202022a0000000303030303030303030303030303030303030303030303030303030303030303",
			wantErr: false,
		},
		{
			name: "testnet real time lock args", // https://pudge.explorer.nervos.org/address/ckt1qqqvm786kru2ccu8tr4lt6j7gpftr4c73fmmnap389ccvg0ksjfjvqtlqqqqqyqqqqq9kqqqqp0sqqqqfvqqqqqsqqqqqvqqqqqrzqqqqrfrwcdnvssswdwpn3s9v8fp87emat306ctjwsm3nmlkjg8qyza2cqgkqqqqqqqph5u5w6fa54waqf2ymhlkl626lyk6njylqcqqqqq8w2xwvj3m2p0wswcv92hgvn0ym4ex7g46ukx474zjl8svt2qkjyrcahk7
			args: args{
				toLock: &types.Script{
					CodeHash: types.HexToHash("0xd23761b364210735c19c60561d213fb3beae2fd6172743719eff6920e020baac"),
					HashType: types.HashTypeType,
					Args:     common.FromHex("0x0001bd3947693da55dd02544ddff6fe95af92da9c89f"),
				},
				btcTxId: "9116a8c5e0f952545f8de5ba226f72dde44d86ae2a0c3be85e503b4ae68c7207",
				after:   6,
			},
			want:    "7f000000100000005b0000005f0000004b000000100000003000000031000000d23761b364210735c19c60561d213fb3beae2fd6172743719eff6920e020baac01160000000001bd3947693da55dd02544ddff6fe95af92da9c89f0600000007728ce64a3b505ee83b0c2aae864de4dd726f22bae58d5f5452f9e0c5a81691",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenBtcTimeLockArgs(tt.args.toLock, tt.args.btcTxId, tt.args.after)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenBtcTimeLockArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenBtcTimeLockArgs() got = %v, want %v", got, tt.want)
			}
		})
	}
}
