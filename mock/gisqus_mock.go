package mock

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
)

// JSONData is aliased to string, to signal intent of usage
type JSONData string

// MockedURL is aliased to string, to signal intent of usage
type MockedURL string

// Server makes it possible to redirect remote http calls to local calls
type Server struct {
	baseDir          string
	httpServer       *httptest.Server
	mockedURLAndData map[MockedURL]JSONData
}

// NewMockServer returns a MockServer reading test data from json mock data
func NewMockServer() *Server {

	ms := Server{}
	f := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Ratelimit-Remaining", "999")
		w.Header().Set("X-Ratelimit-Limit", "1000")
		w.Header().Set("X-Ratelimit-Reset", "1495785600")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)

		path := r.URL.Path
		fmt.Fprint(w, ms.mockedURLAndData[MockedURL(path)])

	}

	ms.httpServer = httptest.NewServer(http.HandlerFunc(f))
	ms.mockedURLAndData = make(map[MockedURL]JSONData)
	return &ms
}

// Close closes the underlying http test server
func (ms *Server) Close() {
	ms.httpServer.Close()
}

// SwitchHostAndScheme informs the Server of an url to mock, provides it with the mocking data and returns the mocked url
func (ms *Server) SwitchHostAndScheme(sourceURL, data string) (string, error) {

	newValues := ms.httpServer.URL
	newValuesU, err := url.Parse(newValues)
	if err != nil {
		return "", err
	}

	sourceU, err := url.Parse(sourceURL)
	if err != nil {
		return "", err
	}
	sourceU.Host = newValuesU.Host
	sourceU.Scheme = newValuesU.Scheme

	ms.mockedURLAndData[MockedURL(sourceU.Path)] = JSONData(data)

	return sourceU.String(), nil
}
