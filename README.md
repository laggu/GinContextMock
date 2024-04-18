# GinContextMock
Test your Gin handler easily!

## Install
```
go get -u github.com/laggu/GinContextMock
```

## Features
GinContextMock makes gin.Context which is set with various variables you need

### What you can do
* setting headers
* setting uri params
* setting queries
* setting body

## Examples

### Headers

#### by struct
```
header := struct{
    Foo string `header:"foo"`
    Bar string `header:"bar"`
}{
    Foo: "abc",
    Bar: "xyz",
}

mock := GinContextMock.NewMock()
mock.SetHeaders(header)

context, err := mock.GetContext()
require.NoError(t, err)

yourHandler(context)
```

#### by map
```
header := map[string]string
header["foo"] = "abc"
header["bar"] = "xyz"

mock := GinContextMock.NewMock()
mock.SetHeaders(header)

yourHandler(mock.GetContext())
```

### URI Params
#### by struct
```
param := struct{
    Foo string `uri:"foo"`
    Bar string `uri:"bar"`
}{
    Foo: "abc",
    Bar: "xyz",
}

mock := GinContextMock.NewMock()
mock.SetURIParams(param)

yourHandler(mock.GetContext())
```