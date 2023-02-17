package stbb

import (
	"os"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestNewInstance(t *testing.T) {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		t.Fatal(err)
	}
	cfg := &Config{}
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		t.Fatal(err)
	}
	_, err = NewInstance(cfg)
	if err != nil {
		t.Fatal(err)
	}
}
