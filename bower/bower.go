package bower

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// Component represents a Bower component (defined in a bower.json file). See
// http://bower.io/#defining-a-package for a quick summary and
// https://github.com/bower/bower.json-spec for the full bower.json
// specification.
type Component struct {
	Name            string            `json:"name"`
	Version         string            `json:"version"`
	Ignore          []string          `json:"ignore"`
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
	Private         bool              `json:"private,omitempty"`

	// TODO(sqs): add Main
}

type Registry struct {
	BaseURL *url.URL
}

var DefaultRegistry = Registry{BaseURL: &url.URL{Scheme: "https", Host: "bower.herokuapp.com"}}

type LookupResponse struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (r Registry) Lookup(pkg string) (*LookupResponse, error) {
	url := r.BaseURL.ResolveReference(&url.URL{Path: "/packages/" + url.QueryEscape(pkg)})
	resp, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}

	var lr *LookupResponse
	err = json.NewDecoder(resp.Body).Decode(&lr)
	if err != nil {
		return nil, err
	}

	return lr, nil
}
