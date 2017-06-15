// Copyright Piero de Salvia.
// All Rights Reserved

package gisqus

import (
	"context"
	"io/ioutil"
	"net/url"
	"os"
	"testing"

	"github.com/pierods/gisqus/mock"
)

var mockServer *mock.Server
var testErr error
var testGisqus Gisqus
var testCtx context.Context
var testValues url.Values
var testDataDir string

func init() {
	testGisqus = NewGisqus("secret")
	testCtx, _ = context.WithCancel(context.TODO())
	mockServer = mock.NewMockServer()

	goPath := os.Getenv("GOPATH")
	testDataDir = goPath + "/src/github.com/pierods/gisqus/testdata/"

}

func TestMain(m *testing.M) {

	defer mockServer.Close()
	retCode := m.Run()
	os.Exit(retCode)
}

func readTestFile(fileName string) (string, error) {

	f, err := os.Open(testDataDir + fileName)
	defer f.Close()

	if err != nil {
		return "", err
	}
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}

	return string(bytes), nil

}
