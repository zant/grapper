package crypto

import (
  "crypto/ecdsa"
  "crypto/x509"
  "encoding/pem"

  "github.com/ethereum/go-ethereum/common/hexutil"
  "github.com/ethereum/go-ethereum/crypto"
)

// TODO make this work
const privateKeyType = "PRIVATE_KEY"
const publicKeyType = "PUBLIC_KEY"

func Encode(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) (string, string) {
  encoded := hexutil.Encode(crypto.FromECDSA(privateKey))
  pemEncoded := pem.EncodeToMemory(&pem.Block{Type: privateKeyType, Bytes: []byte(encoded)})

  encodedPub := crypto.FromECDSAPub(publicKey)
  pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: publicKeyType, Bytes: encodedPub})

  return string(pemEncoded), string(pemEncodedPub)
}

func Decode(pemEncoded string, pemEncodedPub string) (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
  block, _ := pem.Decode([]byte(pemEncoded))
  x509Encoded := block.Bytes
  privateKey, _ := x509.ParseECPrivateKey(x509Encoded)

  blockPub, _ := pem.Decode([]byte(pemEncodedPub))
  x509EncodedPub := blockPub.Bytes
  genericPublicKey, _ := x509.ParsePKIXPublicKey(x509EncodedPub)
  publicKey := genericPublicKey.(*ecdsa.PublicKey)

  return privateKey, publicKey
}
