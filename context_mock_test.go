package GinContextMock

import (
	"encoding/json"
	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func TestNewMock(t *testing.T) {
	require.NotNil(t, NewMock())
}

func TestContextMock_GetContext(t *testing.T) {
	t.Run("empty_mock", func(t *testing.T) {
		mock := NewMock()
		c, err := mock.GetContext()
		require.NoError(t, err)
		require.NotNil(t, c)
	})
	t.Run("headers", func(t *testing.T) {
		headers := map[string]string{}
		err := faker.FakeData(&headers)
		require.NoError(t, err)

		mock := NewMock()
		mock.SetHeaders(headers)

		c, err := mock.GetContext()
		require.NoError(t, err)
		require.NotNil(t, c)

		for key, value := range headers {
			require.Equal(t, value, c.Request.Header.Get(key))
		}
	})
	t.Run("fail_headers", func(t *testing.T) {
		headers := map[string]int{}
		err := faker.FakeData(&headers)
		require.NoError(t, err)

		mock := NewMock()
		mock.SetHeaders(headers)

		c, err := mock.GetContext()
		require.Equal(t, ErrUnsupportedHeaderType, err)
		require.Nil(t, c)
	})
	t.Run("uri_params", func(t *testing.T) {
		params := map[string]string{}
		err := faker.FakeData(&params)
		require.NoError(t, err)

		mock := NewMock()
		mock.SetURIParams(params)
		require.Equal(t, params, mock.URIParams)

		c, err := mock.GetContext()
		require.NoError(t, err)
		require.NotNil(t, c)

		require.Equal(t, len(c.Params), len(params))
		for _, param := range c.Params {
			require.Equal(t, params[param.Key], param.Value)
		}
	})
	t.Run("fail_uri_params", func(t *testing.T) {
		params := map[string]int{}
		err := faker.FakeData(&params)
		require.NoError(t, err)

		mock := NewMock()
		mock.SetURIParams(params)

		c, err := mock.GetContext()
		require.Equal(t, ErrUnsupportedURIParamType, err)
		require.Nil(t, c)
	})
	t.Run("queries", func(t *testing.T) {
		queries := map[string]string{}
		err := faker.FakeData(&queries)
		require.NoError(t, err)

		mock := NewMock()
		mock.SetQueries(queries)

		c, err := mock.GetContext()
		require.NoError(t, err)
		require.NotNil(t, c)

		for key, value := range queries {
			require.Equal(t, value, c.Request.URL.Query().Get(key))
		}
	})
	t.Run("fail_queries", func(t *testing.T) {
		queries := map[string]int{}
		err := faker.FakeData(&queries)
		require.NoError(t, err)

		mock := NewMock()
		mock.SetQueries(queries)

		c, err := mock.GetContext()
		require.Equal(t, ErrUnsupportedQueryType, err)
		require.Nil(t, c)
	})
	t.Run("body", func(t *testing.T) {
		type Body struct {
			A string
			B int
			C float64
		}
		body1 := Body{}
		err := faker.FakeData(&body1)
		require.NoError(t, err)

		mock := NewMock()
		mock.SetBody(body1)

		c, err := mock.GetContext()
		require.NoError(t, err)
		require.NotNil(t, c)

		body2 := Body{}
		decoder := json.NewDecoder(c.Request.Body)
		err = decoder.Decode(&body2)
		require.NoError(t, err)
		require.Equal(t, body1, body2)
	})
	t.Run("fail_body", func(t *testing.T) {
		body := make(chan int)

		mock := NewMock()
		mock.SetBody(body)

		c, err := mock.GetContext()
		require.Equal(t, ErrUnsupportedBodyType, err)
		require.Nil(t, c)
	})
	t.Run("all", func(t *testing.T) {
		data := struct {
			Headers   map[string]string
			URIParams map[string]string
			Queries   map[string]string
			Body      struct {
				A string
				B int
				C float64
			}
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
		require.NotNil(t, c)
	})
}
