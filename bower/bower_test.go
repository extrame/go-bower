package bower

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// registry is configured to use the test HTTP server
	registry Registry

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
)

// setup sets up a test HTTP server along with a github.Client that is
// configured to talk to that test server.  Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup() {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	// registry configured to use test server
	registry.BaseURL, _ = url.Parse(server.URL)
}

// teardown closes the test HTTP server.
func teardown() {
	server.Close()
}

func TestLookup(t *testing.T) {
	setup()
	defer teardown()

	wantLR := LookupResponse{Name: "foo", URL: "git://example.com/foo.git"}
	var handled bool
	mux.HandleFunc("/packages/foo", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(wantLR)
		handled = true
	})

	lr, err := registry.Lookup("foo")
	if err != nil {
		t.Error(err)
	}
	if !handled {
		t.Error("!handled")
	}

	if wantLR != *lr {
		t.Errorf("want response %+v, got %+v", wantLR, lr)
	}
}
