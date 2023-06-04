package client

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/base64"
	"github.com/btcsuite/btcd/btcec/v2"
	"math/big"

	lib_did "fivi/lib/did"
	"github.com/pkg/errors"
)

func VerifySignature(didDoc *lib_did.DID, base64EncodedSig, message string, publicKeyID string) error {
	ecdsaPublicKey, err := getPublicKeyByKeyID(didDoc, publicKeyID)
	if err != nil {
		return err
	}

	return verifySignature(base64EncodedSig, message, ecdsaPublicKey)
}

func getPublicKeyByKeyID(did *lib_did.DID, publicKeyID string) (ecdsa.PublicKey, error) {
	var (
		ecdsaPublicKey ecdsa.PublicKey
		isFound        bool
	)

	for _, verificationMethod := range did.DidDocument.VerificationMethods {
		if publicKeyID == verificationMethod.ID {
			xBytes, err := base64.RawURLEncoding.DecodeString(verificationMethod.PublicKeyJwk.X)
			if err != nil {
				return ecdsaPublicKey, err
			}

			yBytes, err := base64.RawURLEncoding.DecodeString(verificationMethod.PublicKeyJwk.Y)
			if err != nil {
				return ecdsaPublicKey, err
			}

			ecdsaPublicKey = ecdsa.PublicKey{
				Curve: btcec.S256(),
				X:     big.NewInt(0).SetBytes(xBytes),
				Y:     big.NewInt(0).SetBytes(yBytes),
			}

			isFound = true
			break
		}
	}

	if !isFound {
		return ecdsaPublicKey, errors.Errorf("public key id %s isn't in did doc", publicKeyID)
	}

	return ecdsaPublicKey, nil
}

func VerifyAuthorizationSignature(base64EncodedSig, message string, XKey string, YKey string) error {
	xBytes, err := base64.RawURLEncoding.DecodeString(XKey)
	if err != nil {
		return err
	}

	yBytes, err := base64.RawURLEncoding.DecodeString(YKey)
	if err != nil {
		return err
	}

	ecdsaPublicKey := ecdsa.PublicKey{
		Curve: btcec.S256(),
		X:     big.NewInt(0).SetBytes(xBytes),
		Y:     big.NewInt(0).SetBytes(yBytes),
	}

	return verifySignature(base64EncodedSig, message, ecdsaPublicKey)
}

func verifySignature(base64EncodedSig string, message string, ecdsaPublicKey ecdsa.PublicKey) error {
	sig, err := base64.RawURLEncoding.DecodeString(base64EncodedSig)
	if err != nil {
		return err
	}
	sigR := big.Int{}
	sigS := big.Int{}
	sigR.SetBytes(sig[:32])
	sigS.SetBytes(sig[32:])

	messageHash := sha256.Sum256([]byte(message))
	digest := messageHash[:]
	ok := ecdsa.Verify(&ecdsaPublicKey, digest, &sigR, &sigS)
	if !ok {
		return errors.New("invalid signature")
	}

	return nil
}
