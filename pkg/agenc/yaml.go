package agenc

import (
	"io/ioutil"
	"path/filepath"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// ReadYAMLFile reads given yaml file and decodes it into given destination
func ReadYAMLFile(filename string, dest interface{}) error {
	data, err := ioutil.ReadFile(filepath.Clean(filename))
	if err != nil {
		return errors.Wrap(err, "can't read file")
	}
	return errors.Wrap(yaml.Unmarshal(data, dest), "can't unmarshal yaml")
}
