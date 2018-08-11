package oidcstatic

import (
	"io/ioutil"

	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
)

// MappingFile represents the config file that maps users to groups
type MappingFile struct {
	// Email maps user by email address
	Email map[string]struct {
		// Groups is a list of groups this user becomes a member of
		Groups []string `json:"groups"`
	} `json:"email"`
}

func LoadMappings(path string) (MappingFile, error) {
	var m MappingFile

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return m, errors.Wrapf(err, "Error reading %s", path)
	}

	if err := yaml.Unmarshal(b, &m); err != nil {
		return m, err
	}

	return m, nil
}
