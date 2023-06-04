package jwt

import (
	"context"

	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc/metadata"
)

// Predefined context key for DID
var DIDKey DIDContextKey = "did"

type (
	Claims struct {
		DID string `json:"did,omitempty"`
		jwt.StandardClaims
	}

	// DID context key
	DIDContextKey string
)

// DIDContextKey to string
func (d DIDContextKey) String() string {
	return string(d)
}

func DIDFromCtx(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", NewErrDIDIsNotSet("can't get values from context")
	}

	if len(md[DIDKey.String()]) != 1 {
		return "", NewErrDIDIsNotSet("can't get did from context")
	}

	did := md[DIDKey.String()][0]
	if did == "" {
		return "", NewErrDIDIsNotSet("did is empty")
	}

	return did, nil
}

func DIDFromOutgoingCtx(ctx context.Context) (string, error) {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		return "", NewErrDIDIsNotSet("can't get values from context")
	}

	if len(md[DIDKey.String()]) != 1 {
		return "", NewErrDIDIsNotSet("can't get did from context")
	}

	did := md[DIDKey.String()][0]
	if did == "" {
		return "", NewErrDIDIsNotSet("did is empty")
	}

	return did, nil
}

// Set did to context
func SetDIDToCtx(ctx context.Context, did string) context.Context {
	return context.WithValue(ctx, DIDKey, did)
}
