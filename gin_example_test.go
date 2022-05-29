package GinContextMock

import (
	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGin(t *testing.T) {
	type TestBody struct {
		A string
		B int
		C float64
	}

	data := struct {
		Headers   map[string]string
		URIParams map[string]string
		Queries   map[string]string
		Body      TestBody
	}{}

	err := faker.FakeData(&data)
	require.NoError(t, err)

	mock := NewMock()
	mock.SetHeaders(data.Headers)
	mock.SetURIParams(data.URIParams)
	mock.SetQueries(data.Queries)
	mock.SetBody(data.Body)

	c, err := mock.GetContext()
	require.NoError(t, err)

	handler := func(c *gin.Context) {
		for key, value := range data.Headers {
			require.Equal(t, value, c.GetHeader(key))
		}

	}

	handler(c)
}
