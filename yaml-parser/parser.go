package yamlparser

import (
	"os"

	"github.com/goccy/go-yaml"
)

func ParseYaml(yamlPath string) (Config, error) {
	data, err := os.ReadFile(yamlPath)
	if err != nil {
		return Config{}, err
	}
	var pingguConfig Config
	if err := yaml.Unmarshal(data, &pingguConfig); err != nil {
		return Config{}, err
	}
	return pingguConfig, nil
}
