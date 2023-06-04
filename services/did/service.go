package comments

import (
	"context"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fivi/gen/go/did/v1"
	profilepb "fivi/gen/go/profile/v1"
	did_lib "fivi/lib/did"
	did_client "fivi/lib/did/client"
	"fivi/lib/jwt"
	"fivi/services/did/repository"
	"fmt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/rand"
	"strings"
)

const (
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// len of message to sigh.
	messageLen = 30
)

var _ did.DIDServer = (*Service)(nil)

type (
	Service struct {
		repo repository.Queries

		client did_lib.Interface
		jwt    *jwt.JWT

		profiles profilepb.ProfileServiceClient
	}
)

func New(repo repository.Queries, profiles profilepb.ProfileServiceClient) *Service {
	return &Service{
		repo:     repo,
		profiles: profiles,
	}
}

func (s *Service) RegisterURIMapping(ctx context.Context, request *did.DIDMappingRequest) (*did.DIDMappingResponse, error) {
	_, err := s.client.RegisterDID(request.GetLongFormUri())
	if err != nil {
		return nil, err
	}

	doc, err := s.client.ResolveDID(request.GetLongFormUri())
	if err != nil {
		return nil, err
	}

	document, err := json.Marshal(doc)
	if err != nil {
		return nil, err
	}

	_, err = s.repo.CreateDIDMapping(ctx, repository.CreateDIDMappingParams{
		Did: request.GetDid(),
		// TODO: change to long form
		LongFormUri: request.GetLongFormUri(),
		Document:    document,
	})

	didDocument, err := json.Marshal(doc.DidDocument)
	if err != nil {
		return nil, err
	}

	log.Info("registred - ", request.GetDid(), request.GetLongFormUri(), string(didDocument))
	return &did.DIDMappingResponse{}, nil
}

func (s *Service) GetMessageToSign(ctx context.Context, request *did.MessageRequest) (*did.MessageResponse, error) {
	message := make([]rune, messageLen)
	for i := range message {
		message[i] = rune(letters[rand.Intn(len(letters))])
	}

	err := s.repo.UpsertDIDMessage(ctx, repository.UpsertDIDMessageParams{
		Did:     request.Did,
		Message: string(message),
	})
	if err != nil {
		return nil, err
	}

	return &did.MessageResponse{
		Message: string(message),
	}, nil
}

func (s *Service) VerifySignature(ctx context.Context, request *did.VerifySignatureRequest) (*did.VerifySignatureResponse, error) {
	didMessage, err := s.repo.GetDIDMessage(ctx, request.Did)
	if err != nil {
		return nil, err
	}

	id, err := s.GetDID(ctx, request.Did)
	if err != nil {
		return nil, err
	}

	if id == nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("couldn't resolve did doc from did %v", request.Did))
	}

	jwsParts := strings.Split(request.Jws, ".")
	if len(jwsParts) != 3 {
		return nil, errors.New("invalid signature, JWS should have 3 parts")
	}

	headers := jwsParts[0]
	signature := jwsParts[2]
	// format of encoded payload should be in format "word".
	decodedMessage := base64.RawStdEncoding.EncodeToString([]byte(`"` + didMessage.Message + `"`))
	msg := headers + "." + decodedMessage

	if len(id.DidDocument.VerificationMethods) == 0 {
		return nil, errors.Errorf("did %s doesn't have verification methods", id.DidDocument.ID)
	}

	err = did_client.VerifySignature(id, signature, msg, id.DidDocument.VerificationMethods[0].ID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to verify signature within DID-service VerifySignature API")
	}

	_, token, err := s.jwt.NewWithUserData(request.Did)

	return &did.VerifySignatureResponse{
		Jwt: token,
	}, err
}

func (s *Service) GetDID(ctx context.Context, didShortForm string) (*did_lib.DID, error) {
	doc := new(did_lib.DID)

	// retrieve from db.
	mapping, err := s.repo.GetDIDMapping(ctx, didShortForm)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return &did_lib.DID{}, err
	}
	if err == nil {
		doc, err = did_client.Decode(mapping.Document)
		if err != nil {
			return &did_lib.DID{}, err
		}

		return doc, nil
	}

	// retrieve from API client.
	doc, err = s.client.ResolveDID(didShortForm)
	if err != nil {
		return nil, errors.Wrap(err, "could not resolve DID ION document")
	}

	if doc.DidDocument == nil {
		return nil, errors.Wrap(err, fmt.Sprintf("could get DID ION document by did %v", didShortForm))
	}

	return doc, err
}
