package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"net/http"
)

type Config struct {
	Company              string   `json:"company"`
	Section              string   `json:"section"`
	HomebrewTaps         []string `json:"homebrew_taps"`
	HomebrewPackages     []string `json:"homebrew_packages"`
	HomebrewCaskPackages []string `json:"homebrew_cask_packages"`
}

func GetConfig(url string) (Config, error) {
	resp, err := http.Get(url)
	if err != nil {
		return Config{}, fmt.Errorf("failed to fetch YAML file: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Config{}, fmt.Errorf("failed to fetch YAML file: HTTP status %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read YAML file: %w", err)
	}

	var config Config
	err = yaml.Unmarshal(body, &config)
	if err != nil {
		return Config{}, fmt.Errorf("failed to parse YAML file: %w", err)
	}

	return config, nil
}

func ExportYaml(config Config) string {
	data, err := yaml.Marshal(&config)
	if err != nil {
		return fmt.Sprintf("failed to export config: %s", err)
	}
	return string(data)
}
