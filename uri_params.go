package GinContextMock

import (
	"errors"
	"github.com/gin-gonic/gin"
	"reflect"
)

var (
	ErrUnsupportedURIParamType = errors.New("unsupported uri param type")
)

type uriParams struct {
	URIParams interface{}
}

func (p *uriParams) SetURIParams(uriParams interface{}) {
	p.URIParams = uriParams
}

func (p *uriParams) writeURIParamsToContext(c *gin.Context) error {
	if p.URIParams == nil {
		return nil
	}

	switch reflect.ValueOf(p.URIParams).Kind() {
	case reflect.Map:
		return p.writeURIParamsWithMap(c)
	case reflect.Ptr, reflect.Struct:
		return p.writeURIParamsWithObject(c)
	default:
		return ErrUnsupportedURIParamType
	}
}

func (p *uriParams) writeURIParamsWithMap(c *gin.Context) error {
	uriParams, ok := p.URIParams.(map[string]string)
	if !ok {
		return ErrUnsupportedURIParamType
	}
	for key, value := range uriParams {
		c.Params = append(c.Params, gin.Param{Key: key, Value: value})
	}
	return nil
}

func (p *uriParams) writeURIParamsWithObject(c *gin.Context) error {
	var value reflect.Value
	switch reflect.ValueOf(p.URIParams).Kind() {
	case reflect.Ptr:
		value = reflect.ValueOf(p.URIParams).Elem()
		if value.Kind() != reflect.Struct {
			return ErrUnsupportedURIParamType
		}
	case reflect.Struct:
		value = reflect.ValueOf(p.URIParams)
	default:
		return ErrUnsupportedURIParamType
	}

	for i := 0; i < value.NumField(); i++ {
		field := value.Type().Field(i)
		tag := field.Tag.Get("uri")
		if tag == "" {
			continue
		}
		c.Params = append(c.Params, gin.Param{Key: tag, Value: value.Field(i).String()})
	}

	return nil
}
