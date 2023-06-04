//go:build !mock_did

package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/13x-tech/ion-api-go/pkg/api"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	lib_did "fivi/lib/did"
)

type (
	Client struct {
		apiEndpoint       string
		challengeEndpoint string
		discoverEndpoint  string
	}
)

func NewClient(apiEndpoint, challengeEndpoint, discoverEndpoint string) lib_did.Interface {
	return &Client{
		apiEndpoint:       apiEndpoint,
		challengeEndpoint: challengeEndpoint,
		discoverEndpoint:  discoverEndpoint,
	}
}

func (c *Client) RegisterDID(longFormURI string) (string, error) {
	suffixData, delta, err := api.ParseLongForm(longFormURI)
	if err != nil {
		return "", errors.Wrap(err, "invalid long form uri")
	}

	opts := []api.Options{
		api.WithEndpoint(c.apiEndpoint + "/operations"),
	}
	if c.challengeEndpoint != "" {
		opts = append(opts, api.WithChallenge(c.challengeEndpoint))
	}
	ionAPI, err := api.New(opts...)
	if err != nil {
		return "", errors.Wrap(err, "could not create api")
	}

	response, err := ionAPI.Submit(
		api.CreateOperation(suffixData, delta),
	)
	if err != nil {
		return "", errors.Wrap(err, "could not submit request")
	}

	return string(response), nil
}

func (c *Client) ResolveDID(did string) (*lib_did.DID, error) {
	url := fmt.Sprintf("%v/identifiers/%v", c.discoverEndpoint, did)

	log.Infof("we're trying to resolved DID: %v\n", did)
	log.Infof("apiEndpoint: %v\n", c.apiEndpoint)
	log.Infof("url: %v\n", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	rawBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	log.Infof("ResolveDID Status code: %v\n", resp.StatusCode)
	log.Infof("ResolveDID Raw body: %s\n", string(rawBody))
	if !isStatusCodeSuccess(resp.StatusCode) {
		return nil, errors.Errorf("unexpected status code: %v, body: %s", resp.StatusCode, rawBody)
	}

	var didDoc lib_did.DID
	if err := json.Unmarshal(rawBody, &didDoc); err != nil {
		return nil, err
	}

	return &didDoc, nil
}

func isStatusCodeSuccess(code int) bool {
	return code >= http.StatusOK && code < 300
}
