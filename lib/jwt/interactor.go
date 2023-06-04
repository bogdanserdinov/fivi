package jwt

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

var defaultSigningMethod jwt.SigningMethod = jwt.SigningMethodHS256

type (
	JWT struct {
		signingKey []byte
		expIn      time.Duration
	}
)

// NewInteractor is a factory function, returns a new instance of the JWT struct
func NewInteractor(signingKey []byte, expiresIn time.Duration) *JWT {
	return &JWT{
		signingKey: signingKey,
		expIn:      expiresIn,
	}
}

// NewWithUserData returns signed JWT string with did in claims
func (i *JWT) NewWithUserData(did string) (uuid.UUID, string, error) {
	tokenID := uuid.New()
	claims := &Claims{
		DID: did,
		StandardClaims: jwt.StandardClaims{
			Id:        tokenID.String(),
			ExpiresAt: time.Now().Add(i.expIn).Unix(),
			NotBefore: time.Now().Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(defaultSigningMethod, claims)
	ss, err := token.SignedString(i.signingKey)
	if err != nil {
		return uuid.Nil, "", errors.Wrap(err, "could not sign token: %v")
	}
	return tokenID, ss, nil
}

func (i *JWT) ParseJWT(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return i.signingKey, nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "can't parse jwt token")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.Errorf("can't get claims from token")
	}
	if !token.Valid {
		return nil, errors.Errorf("token is invalid")
	}

	return claims, nil
}

func (i *JWT) AuthMiddleware() runtime.ServeMuxOption {
	return runtime.WithMetadata(func(ctx context.Context, request *http.Request) metadata.MD {
		reqToken := request.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		if len(splitToken) != 2 {
			return nil
		}

		token := splitToken[1]
		claims, err := i.ParseJWT(token)
		if err != nil {
			log.Error("couldn't parse jwt token", err)
			return nil
		}

		return metadata.Pairs(DIDKey.String(), claims.DID)
	})
}

// serveError replies to the request with specific code and error message.
func serveError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	var response struct {
		Error string `json:"error"`
	}

	response.Error = err.Error()
	if err = json.NewEncoder(w).Encode(response); err != nil {
		log.Error("failed to write json error response", err)
	}
}
