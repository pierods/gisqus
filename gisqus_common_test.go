// Copyright Piero de Salvia.
// All Rights Reserved

package gisqus

import (
	"context"
	"io/ioutil"
	"net/url"
	"os"

	"github.com/pierods/gisqus/mock"
)

var mockServer *mock.Server
var err error
var testGisqus Gisqus
var testCtx context.Context
var testValues url.Values
var testDataDir string

var mockUsersUrls UsersURLS

func init() {
	testGisqus = NewGisqus("secret")
	testCtx, _ = context.WithCancel(context.TODO())

	goPath := os.Getenv("GOPATH")
	testDataDir = goPath + "/src/github.com/pierods/gisqus/testdata/"

}

func readFile(fileName string) (string, error) {

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
