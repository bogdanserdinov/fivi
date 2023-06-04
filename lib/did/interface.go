package did

type (
	DID struct {
		Context             string            `json:"@context"`
		DidDocument         *Document         `json:"didDocument"`
		DidDocumentMetadata *DocumentMetadata `json:"didDocumentMetadata"`
	}

	Document struct {
		ID      string        `json:"id"`
		Context []interface{} `json:"@context"`
		Service []struct {
			ID              string `json:"id"`
			Type            string `json:"type"`
			ServiceEndpoint struct {
				Nodes []string `json:"nodes"`
			} `json:"serviceEndpoint"`
		} `json:"service"`
		VerificationMethods []VerificationMethod `json:"verificationMethod"`
		Authentication      []string             `json:"authentication"`
	}

	VerificationMethod struct {
		ID           string `json:"id"`
		Controller   string `json:"controller"`
		Type         string `json:"type"`
		PublicKeyJwk struct {
			Crv string `json:"crv"`
			Kty string `json:"kty"`
			X   string `json:"x"`
			Y   string `json:"y"`
		} `json:"publicKeyJwk"`
	}

	DocumentMetadata struct {
		Method struct {
			Published          bool   `json:"published"`
			RecoveryCommitment string `json:"recoveryCommitment"`
			UpdateCommitment   string `json:"updateCommitment"`
		} `json:"method"`
		EquivalentId []string `json:"equivalentId"`
	}
)

//go:generate mockgen -destination=mock_client.go -package=did fivi/lib/did Interface
type Interface interface {
	RegisterDID(longFormURI string) (string, error)
	ResolveDID(did string) (*DID, error)
}
