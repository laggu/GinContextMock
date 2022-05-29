package GinContextMock

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
)

type ContextMock struct {
	headers
	uriParams
	queries
	body
}

func NewMock() *ContextMock {
	return &ContextMock{}
}

func (m *ContextMock) GetContext() (*gin.Context, error) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request = &http.Request{}

	if err := m.writeHeadersToContext(c); err != nil {
		return nil, err
	}
	if err := m.writeURIParamsToContext(c); err != nil {
		return nil, err
	}
	if err := m.writeQueriesToContext(c); err != nil {
		return nil, err
	}
	if err := m.writeBodyToContext(c); err != nil {
		return nil, err
	}

	return c, nil
}
