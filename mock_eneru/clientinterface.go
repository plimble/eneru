package mock_eneru

import "github.com/plimble/eneru"
import "github.com/stretchr/testify/mock"

import "bytes"

import "net/http"

type MockClientInterface struct {
	mock.Mock
}

func NewMockClientInterface() *MockClientInterface {
	return &MockClientInterface{}
}

func (m *MockClientInterface) CreateIndex(req *eneru.CreateIndexReq) (*eneru.CreateIndexResp, error) {
	ret := m.Called(req)

	var r0 *eneru.CreateIndexResp
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*eneru.CreateIndexResp)
	}
	r1 := ret.Error(1)

	return r0, r1
}
func (m *MockClientInterface) Request(method string, path string, query *eneru.Query, body *bytes.Buffer) (*http.Response, error) {
	ret := m.Called(method, path, query, body)

	var r0 *http.Response
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*http.Response)
	}
	r1 := ret.Error(1)

	return r0, r1
}
