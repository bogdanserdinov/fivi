package client

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"math/big"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/go-errors/errors"
)

func Sign(message string, dKey string, xKey string, yKey string) (string, error) {
	d, err := base64.RawURLEncoding.DecodeString(dKey)
	if err != nil {
		return "", err
	}
	dBigInt := big.Int{}
	dBigInt.SetBytes(d)

	x, err := base64.RawURLEncoding.DecodeString(xKey)
	if err != nil {
		return "", err
	}
	xBigInt := big.Int{}
	xBigInt.SetBytes(x)

	y, err := base64.RawURLEncoding.DecodeString(yKey)
	if err != nil {
		return "", err
	}
	yBigInt := big.Int{}
	yBigInt.SetBytes(y)

	ecdsaPublicKey := ecdsa.PublicKey{
		Curve: btcec.S256(),
		X:     &xBigInt,
		Y:     &yBigInt,
	}

	ecdsaPrivateKey := ecdsa.PrivateKey{
		PublicKey: ecdsaPublicKey,
		D:         &dBigInt,
	}

	messageHash := sha256.Sum256([]byte(message))
	digest := messageHash[:]
	r, s, err := ecdsa.Sign(rand.Reader, &ecdsaPrivateKey, digest)
	if err != nil {
		return "", err
	}

	rBytes := r.Bytes()
	for len(rBytes) < 32 {
		rBytes = append([]byte{0}, rBytes...)
	}
	sBytes := s.Bytes()
	for len(sBytes) < 32 {
		sBytes = append([]byte{0}, sBytes...)
	}

	signature := append(rBytes, sBytes...)
	base64Signature := base64.RawURLEncoding.EncodeToString(signature)

	ok := ecdsa.Verify(&ecdsaPublicKey, digest, r, s)
	if !ok {
		return "", errors.Errorf("can't verify just signed message")
	}

	return base64Signature, nil
}
