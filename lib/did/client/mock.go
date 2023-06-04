//go:build mock_did

package client

import (
	"fmt"

	"github.com/golang/mock/gomock"

	lib_did "fivi/lib/did"
	"fivi/test/mock"
)

type stub struct{}

func (l *stub) Errorf(format string, args ...interface{}) {
	panic(fmt.Sprintf("unexpected mock Errorf stub is called!"+format, args...))
}

func (l *stub) Fatalf(format string, args ...interface{}) {
	panic(fmt.Sprintf("unexpected mock Fatalf stub is called!"+format, args...))
}

func NewClient(apiEndpoint, challengeEndpoint, discoverEndpoint string) lib_did.Interface {
	m := mock.GetMockObject(mock.DIDProvider)
	if m == nil {
		m = lib_did.NewMockInterface(gomock.NewController(&stub{}))
		m.(*lib_did.MockInterface).ExpectRegisterDIDAny()
		m.(*lib_did.MockInterface).ExpectResolveDIDAny()
	}
	return m.(lib_did.Interface)
}
