package bls_signer

import (
	"encoding/hex"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"

	"github.com/herumi/bls-eth-go-binary/bls"
	e2types "github.com/wealdtech/go-eth2-types/v2"
)

type EthBLSAccount struct {
	*e2types.BLSPrivateKey
}

func InitEthBLS() error {
	if err := bls.Init(bls.BLS12_381); err != nil {
		return err
	}
	return bls.SetETHmode(bls.EthModeDraft07)
}

func NewEthBLSAccount() EthBLSAccount {
	err := InitEthBLS()
	if err != nil {
		panic(err)
	}
	k := NewEthBLSKey()
	return EthBLSAccount{k}
}

func NewEthBLSKey() *e2types.BLSPrivateKey {
	err := InitEthBLS()
	if err != nil {
		panic(err)
	}
	key, err := e2types.GenerateBLSPrivateKey()
	if err != nil {
		panic(err)
	}
	return key
}

func NewEthSignerBLSFromExistingKeyBytes(sk []byte) EthBLSAccount {
	err := InitEthBLS()
	if err != nil {
		panic(err)
	}
	k, err := e2types.BLSPrivateKeyFromBytes(sk)
	if err != nil {
		panic(err)
	}
	return EthBLSAccount{k}
}

func NewEthSignerBLSFromExistingKey(sk string) EthBLSAccount {
	err := InitEthBLS()
	if err != nil {
		panic(err)
	}
	data, err := hex.DecodeString(sk)
	if err != nil {
		panic(err)
	}
	k, err := e2types.BLSPrivateKeyFromBytes(data)
	if err != nil {
		panic(err)
	}
	return EthBLSAccount{k}
}

func (e *EthBLSAccount) PrivateKeyString() string {
	return ConvertBytesToString(e.BLSPrivateKey.Marshal())
}

func (e *EthBLSAccount) PublicKeyString() string {
	return ConvertBytesToString(e.PublicKey().Marshal())
}

func (e *EthBLSAccount) ZeroXPrefixedPublicKeyString() string {
	return strings_filter.AddHexPrefix(ConvertBytesToString(e.PublicKey().Marshal()))
}
