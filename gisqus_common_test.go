package gisqus

import (
	"context"
	"net/http/httptest"
	"net/url"
	"os"

	"github.com/pierods/gisqus/mock"
)

var mockServer *httptest.Server
var err error
var testGisqus Gisqus
var testCtx context.Context
var testValues url.Values
var testDataDir string
var ms mock.MockServer

func init() {
	testGisqus = NewGisqus("secret")
	testCtx, _ = context.WithCancel(context.TODO())

	goPath := os.Getenv("GOPATH")
	testDataDir = goPath + "/src/github.com/pierods/gisqus/testdata/"

	ms = mock.NewMockServer(testDataDir)

}
