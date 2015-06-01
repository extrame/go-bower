package bower

import (
	"encoding/json"
)

// Component represents a Bower component (defined in a bower.json file). See
// http://bower.io/#defining-a-package for a quick summary and
// https://github.com/bower/bower.json-spec for the full bower.json
// specification.
type Component struct {
	Name            string            `json:"name"`
	Main            interface{}       `json:"main"`
	Version         string            `json:"version"`
	Ignore          []string          `json:"ignore"`
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
	Private         bool              `json:"private,omitempty"`

	// TODO(sqs): add Main
}

// ParseBowerJSON parses a bower.json file from data.
//
// TODO(sqs): apply defaults and normalizations like
// https://github.com/bower/json
func ParseBowerJSON(data []byte) (*Component, error) {
	var c *Component
	err := json.Unmarshal(data, &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
