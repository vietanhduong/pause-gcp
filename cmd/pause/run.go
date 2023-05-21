package pause

import (
	jsoniter "github.com/json-iterator/go"
	apis "github.com/vietanhduong/pause-gcp/apis/v1"
	"log"
	"os"
	"sigs.k8s.io/yaml"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func run(configFile string, force bool) error {
	// parse config
	cfg, err := parseConfigFile(configFile)
	if err != nil {
		return err
	}
	// validate the config
	if err = cfg.ValidateAll(); err != nil {
		return err
	}
	log.Printf("config: %v", cfg)
	return nil
}

func parseConfigFile(path string) (*apis.Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if b, err = yaml.YAMLToJSON(b); err != nil {
		return nil, err
	}

	var cfg apis.Config
	err = json.Unmarshal(b, &cfg)
	return &cfg, nil
}
