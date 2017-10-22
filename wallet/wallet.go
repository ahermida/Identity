/*
    Package wallet manages simple wallet functionality
    That is -- creation of private & public keypairs, as well as addresses
*/
package wallet

import (
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "crypto/sha256"

    "github.com/ahermida/Identity/common"
    "golang.org/x/crypto/ripemd160"
)

const version = byte(0x00)   //alpha
const addressChecksumLen = 4 //like btc

type Wallet struct {
    PrivateKey ecdsa.PrivateKey
    PublicKey  []byte
}

func NewWallet() *Wallet {
    private, public := newKeyPair()
    wallet := &Wallet{private, public}

    return wallet
}

func newKeyPair() (ecdsa.PrivateKey, []byte){
    curve := elliptic.P256()
    private, _ := ecdsa.GenerateKey(curve, rand.Reader)
    pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

    return *private, pubKey
}

func (w Wallet) GetAddress() common.Address {
    pubKeyHash := HashPubKey(w.PublicKey)

    versionedPayload := append([]byte{version}, pubKeyHash...)
    checksum := checksum(versionedPayload)

    fullPayload := append(versionedPayload, checksum...)
    address := common.Base58Encode(fullPayload)

    return address
}

func HashPubKey(pubKey []byte) []byte {
    publicSHA256 := sha256.Sum256(pubKey)

    RIPEMD160Hasher := ripemd160.New()
    RIPEMD160Hasher.Write(publicSHA256[:])
    publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)

    return publicRIPEMD160
}

func checksum(payload []byte) []byte {
    firstSHA := sha256.Sum256(payload)
    secondSHA := sha256.Sum256(firstSHA[:])

    return secondSHA[:addressChecksumLen]
}
