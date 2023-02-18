package test

import (
	"os"
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/kshamiev/urok/database/bbolt/stbb"
)

func TestNewInstance(t *testing.T) {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		t.Fatal(err)
	}
	cfg := &stbb.Config{}
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		t.Fatal(err)
	}
	_, err = stbb.NewInstance(cfg)
	if err != nil {
		t.Fatal(err)
	}
}
