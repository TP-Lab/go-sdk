package iost

import (
	"github.com/iost-official/go-iost/account"
	"github.com/iost-official/go-iost/common"
	"github.com/iost-official/go-iost/crypto"
	"github.com/iost-official/go-sdk/pb"
)

type Keychain struct {
	ID      string
	KeyPair map[string]*account.KeyPair
}

func NewKeychain(id string) *Keychain {
	return &Keychain{
		ID:      id,
		KeyPair: make(map[string]*account.KeyPair),
	}
}

func (k *Keychain) AddKey(seckey []byte, perm ...string) error {
	var alg crypto.Algorithm
	if len(seckey) == 64 {
		alg = crypto.Ed25519
	} else {
		alg = crypto.Secp256k1
	}
	kp, err := account.NewKeyPair(seckey, alg)
	if err != nil {
		return err
	}
	for _, p := range perm {
		k.KeyPair[p] = kp
	}
	return nil
}

func (k *Keychain) Sign(tx *rpcpb.TransactionRequest) {

}

func (k *Keychain) SignTx(tx *rpcpb.TransactionRequest) {
	tx.Publisher = k.ID
	sig := k.KeyPair["active"].Sign(common.Sha3(txToBytes(tx)))

	var thisSig rpcpb.Signature
	thisSig.PublicKey = sig.Pubkey
	thisSig.Algorithm = rpcpb.Signature_Algorithm(int32(uint8(sig.Algorithm)))
	thisSig.Signature = sig.Sig

	tx.PublisherSigs = append(tx.PublisherSigs, &thisSig)
}
