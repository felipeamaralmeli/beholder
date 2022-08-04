package configurationreader

import (
	"gopkg.in/yaml.v3"
	"os"
)

func ReadConfigsFromFile(path string) (map[interface{}]interface{}, error) {
	configs := make(map[interface{}]interface{})

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, configs)

	if err != nil {
		return nil, err
	}

	return configs, nil
}
